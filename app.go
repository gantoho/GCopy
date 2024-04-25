package main

import (
	"context"
	"fmt"
	"os"
	"time"

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

func (a *App) NowTime() time.Time {
	return time.Now()
}

func (a *App) ShowWin() {
	runtime.EventsOn(a.ctx, "showWindow", func(args ...interface{}) {
		runtime.Show(a.ctx)
	})
}

type Note struct {
	ID       int64  `json:"id"`
	FileName string `json:"filename"`
	Content  string `json:"content"`
}

func (a *App) GetContent() ([]Note, error) {
	dst := "./notes/"
	files, err := os.Open(dst)
	if err != nil {
		fmt.Printf("error opening directory: %v", err) //print error if directory is not opened
		err := os.Mkdir("./notes", os.ModePerm)
		if err != nil {
			fmt.Printf("notes directory not created: %v", err)
			fmt.Print("请先在程序根目录下创建notes文件夹")
			return nil, err
		}
		return nil, err
	}
	defer func() {
		err := files.Close()
		if err != nil {
			return
		}
	}()
	fileInfos, err := files.Readdir(-1)
	if err != nil {
		fmt.Println("error reading directory:", err) //if directory is not read properly print error message
		return nil, err
	}
	//fmt.Println("files in notes:", fileInfos)

	var returnNotes []Note
	for index, fileInfos := range fileInfos {
		fileContent, err := os.ReadFile(dst + fileInfos.Name())
		if err != nil {
			fmt.Println("error reading file:", err)
			return nil, err
		}
		//fmt.Println(string(fileContent))
		returnNotes = append(returnNotes, Note{
			ID:       int64(index),
			FileName: fileInfos.Name(),
			Content:  string(fileContent),
		})
	}
	return returnNotes, nil
}

// 向notes文件夹中添加文件
func (a *App) AddNote(note Note) string {
	dst := "./notes/"
	// 创建文件
	file, err := os.Create(dst + note.FileName)
	if err != nil {
		fmt.Println("error creating file:", err)
		return err.Error()
	}
	defer file.Close()

	// 写入内容
	_, err = file.WriteString(note.Content)
	if err != nil {
		fmt.Println("errorwriting to file:", err)
		return err.Error()
	}
	return "ok"
}
