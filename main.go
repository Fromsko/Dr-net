package main

import (
	"embed"
	"guiforclash/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()
	fox := backend.NewApp()

	err := wails.Run(&options.App{
		Title:     "clash for cute-fox",
		Width:     800,
		Height:    550,
		Frameless: true,
		MinWidth:  380,
		MinHeight: 300,
		Windows: &windows.Options{
			WebviewIsTransparent:              true, //网页透明
			WindowIsTranslucent:               true, //窗口透明
			DisableFramelessWindowDecorations: true,
			BackdropType:                      windows.Acrylic,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.startup,
		OnDomReady:       app.OnDomReady,
		Bind: []interface{}{
			app,
			fox,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
