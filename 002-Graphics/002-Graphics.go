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
	mainWindow.SetTitle("walk 002-Graphics")
	mainWindow.SetSize(walk.Size{600, 400})

	// レイアウト
	layout := walk.NewVBoxLayout()
	mainWindow.SetLayout(layout)

	// 描画用の CustomWidget
	customWidget, err := walk.NewCustomWidgetPixels(mainWindow, 0, myPaintFunc)
	if err != nil {
		panic(err)
	}
	customWidget.SetInvalidatesOnResize(true)

	// ウィンドウを表示してイベントループを開始
	mainWindow.Show()
	mainWindow.Run()
}

// 描画関数
func myPaintFunc(canvas *walk.Canvas, updateBounds walk.Rectangle) error {

	// 背景ブラシを作成
	bgBrush, err := walk.NewSystemColorBrush(walk.SysColorWindow)
	if err != nil {
		return err
	}
	defer bgBrush.Dispose()

	// 背景クリア
	if err := canvas.FillRectanglePixels(bgBrush, updateBounds); err != nil {
		return err
	}

	// 描画ペンを作成
	pen, err := walk.NewCosmeticPen(walk.PenSolid, walk.RGB(0, 0, 0))
	if err != nil {
		return err
	}
	defer pen.Dispose()
	
	// 四角形の描画
	canvas.DrawRectanglePixels(pen, walk.Rectangle{X:50, Y:50, Width:150, Height:100})

	// 楕円の描画
	canvas.DrawEllipse(pen, walk.Rectangle{X:100, Y:100, Width:150, Height:100})

	// フォントを作成
	font, err := walk.NewFont("Noto Sans JP", 12, 0)
	if err != nil {
		return err
	}
	defer font.Dispose()

	// テキストを描画
	if err := canvas.DrawTextPixels(
		"Hello World\nこんにちは世界\n",
		font,
		walk.RGB(0, 0, 0),
		walk.Rectangle{X: 70, Y: 100}, // TextNoClip の場合、Width と Height は不用
		walk.TextLeft|walk.TextTop|walk.TextNoClip,
	); err != nil {
		return err
	}

	return nil
}
