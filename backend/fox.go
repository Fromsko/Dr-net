package backend

import (
	_ "embed"
	"fmt"
	"guiforclash/backend/api"
	"guiforclash/backend/cmd"
	"guiforclash/backend/core"
	"net"
	"strings"
)

type (
	Fox struct {
		// ctx context.Context
	}
	ApiIOResult struct {
		Flag bool   `json:"flag"`
		Data string `json:"data"`
	}
)

var (
	ipArea     = make(chan *core.NetworkPair, 3)
	scanStatus = make(chan bool)
)

func (f *Fox) StartScan() *core.NetworkPair {
	go scanHandler()

	if <-scanStatus {
		return <-ipArea
	} else {
		return &core.NetworkPair{}
	}
}

func (f *Fox) SetNetWork(ip string) string {
	// 设置 IP
	if err := cmd.WinNetSet(ip); err == nil {
		return "伪装成功"
	} else {
		return "伪装失败"
	}
}

func (f *Fox) CheckNetWork() string {
	if content, err := api.LoginCheck(); err == nil {
		if content.SchoolCard != "" || content.UserName != "" {
			var role string
			if len(content.SchoolCard) > 11 {
				role = "学生"
			} else {
				role = "教师"
			}
			return fmt.Sprintf("已连接(%s)", role)
		}
	}
	return "未连接"
}

func (f *Fox) GetClientIP() string {
	return NowLocal()
}

func (f *Fox) RestoreNetWork() string {
	return cmd.DefaultNet()
}

// NowLocal 当前IP地址
func NowLocal() (ipAddress string) {
	ipAddress = "127.0.0.1"

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
func scanHandler() {
	splitData := strings.Split(NowLocal(), ".")
	s := fmt.Sprintf("%s.%s.%s.1/22", splitData[0], splitData[1], splitData[2])
	go func() {
		if result := cmd.ScanApp(s); result != nil {
			ipArea <- result
			scanStatus <- true
			return
		}
		scanStatus <- false
	}()
}

func NewApp() *Fox {
	return &Fox{}
}
