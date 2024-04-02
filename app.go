package main

import (
	"context"
	_ "embed"
	"fmt"
	"guiforclash/backend/core"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed build\windows\icon.ico
var iconData []byte
var mShow, mHide, mQuit *systray.MenuItem

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	onExit := func() {
		fmt.Println("Is exited!")
	}

	go systray.Run(onReady, onExit)
}

func (a *App) OnDomReady(ctx context.Context) {
	go func() {
		for {
			select {
			case <-mShow.ClickedCh:
				runtime.Show(ctx)
				runtime.WindowUnminimise(ctx)
				runtime.WindowSetAlwaysOnTop(ctx, true)
			case <-mHide.ClickedCh:
				runtime.Hide(ctx)
			case <-mQuit.ClickedCh:
				runtime.Quit(ctx)
			}
		}
	}()
}

func onReady() {
	systray.SetIcon(iconData)
	systray.SetTooltip("Foxer")
	mShow = systray.AddMenuItem("显示", "显示窗口")
	mHide = systray.AddMenuItem("隐藏", "隐藏窗口")
	systray.AddSeparator()
	mQuit = systray.AddMenuItem("退出", "退出程序")
}

type ApiIOResult struct {
	Flag bool   `json:"flag"`
	Data string `json:"data"`
}

// Writefile 写入文件
func (a *App) Writefile(path string, content string) ApiIOResult {
	var (
		cwd string
		err error
	)

	if result := a.GetUserCwd(); result.Flag {
		cwd = result.Data
	} else {
		cwd, err = os.Getwd()
		if err != nil {
			return ApiIOResult{false, err.Error()}
		}
	}

	path = filepath.Join(cwd, path)
	err = os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return ApiIOResult{false, err.Error()}
	}
	return ApiIOResult{true, "Success"}
}

// Readfile 读取文件
func (a *App) Readfile(path string) ApiIOResult {
	var (
		cwd string
		err error
	)

	if result := a.GetUserCwd(); result.Flag {
		cwd = result.Data
	} else {
		cwd, err = os.Getwd()
		if err != nil {
			return ApiIOResult{false, err.Error()}
		}
	}

	path = filepath.Join(cwd, path)
	b, err := os.ReadFile(path)
	if err != nil {
		return ApiIOResult{false, err.Error()}
	}
	return ApiIOResult{true, string(b)}
}

// Notification 通知
func (a *App) Notification(title string, message string, icon string) ApiIOResult {
	fmt.Println("Notification:", title, message, icon)
	err := beeep.Notify(title, message, icon)
	if err != nil {
		return ApiIOResult{false, err.Error()}
	}
	return ApiIOResult{true, "Success"}
}

// Getcwd 获取当前路径
func (a *App) Getcwd() ApiIOResult {
	cwd, err := os.Getwd()
	if err != nil {
		return ApiIOResult{false, err.Error()}
	}
	return ApiIOResult{true, cwd}
}

// Exec 执行 CMD 命令
func (a *App) Exec(path string, args string) ApiIOResult {
	fmt.Println("Exec:", path, args)
	cwd, err := os.Getwd()
	if err != nil {
		return ApiIOResult{false, err.Error()}
	}
	path = filepath.Join(cwd, path)

	cmd := exec.Command(path, args)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return ApiIOResult{false, err.Error()}
	}

	return ApiIOResult{true, string(out)}
}

// GetUserCwd 获取用户目录
func (a *App) GetUserCwd() ApiIOResult {
	flag, home := core.GetUserPath()
	return ApiIOResult{Flag: flag, Data: home}
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
