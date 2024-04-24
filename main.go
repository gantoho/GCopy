package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	// Create an instance of the app structure
	app := NewApp()

	//AppMenu := menu.NewMenu()
	//FileMenu := AppMenu.AddSubmenu("File")
	//FileMenu.AddSeparator()
	//FileMenu.AddText("Copy", keys.CmdOrCtrl("c"), func(_ *menu.CallbackData) {
	//	str, _ := runtime.ClipboardGetText(app.ctx)
	//	runtime.EventsEmit(app.ctx, "Copy", str)
	//})

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "GCopy",
		Width:         350,
		Height:        500,
		DisableResize: true,
		// Frameless:         true,
		StartHidden:       false,
		HideWindowOnClose: true,
		AlwaysOnTop:       true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		//Menu: AppMenu,
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			BackdropType:         windows.Mica,
			DisableWindowIcon:    true,
			OnSuspend: func() {
				fmt.Println("进入休眠")
			},
			OnResume: func() {
				fmt.Println("进入工作")
			},
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
		BackgroundColour: &options.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
