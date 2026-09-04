package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
)

func toString(value any) string {
	switch v := value.(type) {
	case walk.PaintMode:
		switch v {
		case walk.PaintNormal:
			return "PaintNormal"
		case walk.PaintNoErase:
			return "PaintNoErase"
		case walk.PaintBuffered:
			return "PaintBuffered"
		default:
			return fmt.Sprintf("%v (%T)", v)
		}
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("%v (%T)", v, v)
	}
}

//////////////////////////////////////////////////// 各種オブジェクト

// フォントを作成
func createFont(family string, pointSize int, style walk.FontStyle) *walk.Font {
	font, err := walk.NewFont(family, pointSize, style)
	if err != nil {
		panic(err)
	}
	return font
}

// RGBで指定した色のブラシを作成
func rgbBrush(r, g, b byte) *walk.SolidColorBrush {
	brush, err := walk.NewSolidColorBrush(walk.RGB(r, g, b))
	if err != nil {
		panic(err)
	}
	return brush
}

// 画像ファイルを読み込む
func loadImage(path string, dpi int) *walk.Bitmap {
	bmp, err := walk.NewBitmapFromFileForDPI(path, dpi)
	if err != nil {
		panic(err)
	}
	return bmp
}

//////////////////////////////////////////////////// windows API

// ウィンドウスタイルをビットごとに設定または解除する
func setWindowStyle(w walk.Window, styleBit uint32, set bool) {
	hwnd := w.Handle()
	style := uint32(win.GetWindowLong(hwnd, win.GWL_STYLE))

	if set {
		style |= styleBit
	} else {
		style &^= styleBit
	}

	win.SetWindowLong(hwnd, win.GWL_STYLE, int32(style))
}

// ウィンドウの枠まで含めて再描画する
// ウィンドウスタイルを動的に変更した場合にこれを呼ぶとよい
func redrawFrame(w walk.Window) {
	win.SetWindowPos(
		w.Handle(),
		0, 0, 0, 0, 0,
		win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_NOZORDER|win.SWP_FRAMECHANGED,
	)
}
