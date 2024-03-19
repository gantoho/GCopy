package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	// runtime.EventsOn(ctx, "ctrlo", func(args ...interface{}) {
	// 	arg := args[0].(string)
	// 	windows.MessageBox(0, windows.StringToUTF16Ptr(arg), windows.StringToUTF16Ptr(arg), 0)
	// })
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ShowWin() {
	runtime.EventsOn(a.ctx, "showWindow", func(args ...interface{}) {
		runtime.Show(a.ctx)
	})
}
