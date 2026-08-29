//go:generate goversioninfo
package main

import (
	"github.com/lxn/walk"
)

func main() {
	// メインウィンドウの作成
	mw, err := walk.NewMainWindow()
	if err != nil {
		panic(err)
	}
	
	// メインウィンドウの設定
	mw.SetTitle("001: walk.MainWindow")
	mw.SetSize(walk.Size{600, 400})
	mw.SetMinMaxSize(
		walk.Size{300, 200}, // MinSize 
		walk.Size{800, 600}, // MaxSize
	)

	// レイアウト設定
	layout := walk.NewVBoxLayout()
	mw.SetLayout(layout)
	
	// ウィンドウを表示してイベントループを開始
	mw.Show()
	mw.Run()
}
