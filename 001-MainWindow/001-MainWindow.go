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
    defer mw.Dispose()
	
    // メインウィンドウの設定
    mw.SetTitle("001: walk.MainWindow")
    mw.SetSize(walk.Size{Width:600, Height:400})

    // レイアウト設定
    layout := walk.NewVBoxLayout()
    mw.SetLayout(layout)
	
    // ウィンドウを表示してイベントループを開始
    mw.Show()
    mw.Run()
}
