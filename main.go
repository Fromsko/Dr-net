package main

import (
	"context"
	"fmt"
	"github.com/flopp/go-findfont"
	"net"
	"net-tools/src/api"
	"net-tools/src/cmd"
	"net-tools/src/core"
	"net/http"
	"os"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var ipArea = make(chan *core.NetworkPair, 3)

func init() {
	fontPaths := findfont.List()
	for _, fontPath := range fontPaths {
		fmt.Println(fontPath)
		if strings.Contains(fontPath, "simkai.ttf") {
			err := os.Setenv("FYNE_FONT", fontPath)
			if err != nil {
				return
			}
			break
		}
	}
}

// 构建指令 => fyne package -os windows -icon myapp.png -tags noconsole
func main() {
	// TODO: 启动扫描的时候，设置网络后，使用检测, 概率出现闪退状况
	a := app.New()

	win := a.NewWindow("网络小工具")
	win.Resize(fyne.Size{Width: 225, Height: 90})
	win.SetFixedSize(true)

	clock := widget.NewLabel("")
	updateTime(clock)

	connLabel := widget.NewLabel("连接网络: " + NowLocal())
	statusLabel := widget.NewLabel("网络状况: 未知")
	setStaLabel := widget.NewLabel("伪装状态: 暂未伪装")
	scanStatusLabel := widget.NewLabel("扫描状态: 未启动")

	statusButton := widget.NewButton("检测", func() {
		go statusCheck(statusLabel)
	})
	scanButton := widget.NewButton("扫描", func() {
		scanHandler(scanStatusLabel)
	})
	webStartButton := widget.NewButton("启动服务器", func() {
		scanWeb("run", nil)
	})
	webStopButton := widget.NewButton("关闭服务器", func() {
		scanWeb("stop", scanStatusLabel)
	})
	setButton := widget.NewButton("设置", func() {
		fmt.Println("设置网络")
		setStaLabel.SetText("伪装状态: " + cmd.NetSet(""))
		connLabel.SetText("连接网络: " + NowLocal())
	})
	resetButton := widget.NewButton("恢复", func() {
		setStaLabel.SetText("伪装状态: " + cmd.DefaultNet())
	})

	content := container.NewVBox(
		container.NewHBox(clock),                                            // 时间信息
		container.NewHBox(connLabel),                                        // 连接网络信息
		container.NewHBox(setStaLabel),                                      // 是否伪装
		container.NewHBox(statusLabel),                                      // 网络状况
		container.NewHBox(scanStatusLabel),                                  // 扫描状态
		container.NewHBox(scanButton, setButton, resetButton, statusButton), // 各种按钮
		container.NewHBox(webStartButton, webStopButton),
	)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	defer func() {
		_ = os.Unsetenv("FYNE_FONT")
	}()

	win.SetContent(content)
	win.ShowAndRun()
}

// updateTime 更新时间
func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("当前时间: 2006.01.06 03:04:05")
	clock.SetText(formatted)
}

// statusCheck 状态检测
func statusCheck(statusLabel *widget.Label) {
	if content, err := api.LoginCheck(); err == nil {
		if content.SchoolCard != "" || content.UserName != "" {
			statusLabel.SetText("网络状况: 已连接")
			return
		}
	}
	statusLabel.SetText("网络状况: 未连接")
}

// NowLocal 当前IP地址
func NowLocal() (ipAddress string) {
	ipAddress = "Failed"

	address, err := net.InterfaceAddrs()
	if err != nil {
		return ipAddress
	}

	for _, addr := range address {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() == nil {
			continue
		}
		if ipInfo := ipNet.IP.String()[:3]; ipInfo != "10." {
			continue
		}
		ipAddress = ipNet.IP.String()
	}
	return ipAddress
}

// scanHandler 启动任务扫描
func scanHandler(scanLabel *widget.Label) {
	splitData := strings.Split(NowLocal(), ".")
	s := fmt.Sprintf("%s.%s.%s.1/22", splitData[0], splitData[1], splitData[2])
	go func(scanLabel *widget.Label) {
		scanLabel.SetText("扫描状态: 正在扫描")
		if result := cmd.ScanApp(s); result != nil {
			scanLabel.SetText("扫描状态: 完成扫描")
			ipArea <- result
		}
	}(scanLabel)
}

var server *http.Server

func scanWeb(Option string, scanLabel *widget.Label) {
	if Option == "run" {
		// 启动一个Gin服务
		go func() {
			server = &http.Server{
				Addr:    ":8080",                 // 设置Gin服务监听的地址
				Handler: cmd.GetRouter(<-ipArea), // 这里假设cmd.GetRouter()返回Gin的路由器实例
			}
			err := server.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				// 处理启动失败的情况
				panic("启动失败!")
			}
		}()
	} else if Option == "stop" {
		// 发送信号关闭已经运行的Gin服务
		go func(scanLabel *widget.Label) {
			if server != nil {
				// 创建一个超时上下文，以便在一定时间后强制关闭服务
				ctx, cancel := context.WithTimeout(
					context.Background(),
					3*time.Second,
				)
				// 关闭服务
				if err := server.Shutdown(ctx); err != nil {
					panic("关闭失败")
				}
				defer cancel()
				fmt.Println("服务器已关闭")
				scanLabel.SetText("扫描状态: 未启动")
			}
		}(scanLabel)
	}
}
