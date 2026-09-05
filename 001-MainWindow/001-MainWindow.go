//go:generate goversioninfo
package main

import (
	"github.com/lxn/walk"
)

func main() {
	// メインウィンドウの作成
	mainWindow, err := walk.NewMainWindow()
	if err != nil {
		panic(err)
	}

	// タイトル、サイズ
	mainWindow.SetTitle("walk 001-MainWindow")
	mainWindow.SetSize(walk.Size{600, 400})

	// 背景色
	bgBrush, err := walk.NewSystemColorBrush(walk.SysColorWindow)
	if err != nil {
		panic(err)
	}
	mainWindow.SetBackground(bgBrush)

	// レイアウト
	layout := walk.NewVBoxLayout()
	mainWindow.SetLayout(layout)

	// ウィンドウを表示してイベントループを開始
	mainWindow.Show()
	mainWindow.Run()
}
