package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
	"golang.org/x/sys/windows"
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
			return fmt.Sprintf("%v (%T)", v, v)
		}

	case walk.ImageViewMode:
		switch v {
		case walk.ImageViewModeIdeal:
			return "ImageViewModeIdeal"
		case walk.ImageViewModeCorner:
			return "ImageViewModeCorner"
		case walk.ImageViewModeCenter:
			return "ImageViewModeCenter"
		case walk.ImageViewModeZoom:
			return "ImageViewModeZoom"
		case walk.ImageViewModeStretch:
			return "ImageViewModeStretch"
		case walk.ImageViewModeShrink:
			return "ImageViewModeShrink"
		default:
			return fmt.Sprintf("%v (%T)", v, v)
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

// CTRLキーが押されているか調べる
var (
	user32           = windows.NewLazySystemDLL("user32.dll")
	getAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

const VK_CONTROL = 0x11

func isCtrlPressed() bool {
	ret, _, _ := getAsyncKeyState.Call(uintptr(VK_CONTROL))
	return (ret & 0x8000) != 0
}