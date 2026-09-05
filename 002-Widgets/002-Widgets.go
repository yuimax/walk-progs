//go:generate goversioninfo
package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

func main() {
	imageFiles := [...]string{
		"../files/daisy-vroid.jpg",
		"../files/cathy-vroid.jpg",
		"../files/daisy-renoir.jpg",
		"../files/cathy-renoir.jpg",
		"../files/daisy-ukiyoe.jpg",
		"../files/cathy-ukiyoe.jpg",
	}

	imageFileIndex := -1

	imageViewModes := [...]struct {
		Name string
		Mode walk.ImageViewMode
	}{
		{"Corner", walk.ImageViewModeCorner},
		{"Center", walk.ImageViewModeCenter},
		{"Zoom", walk.ImageViewModeZoom},
		{"Stretch", walk.ImageViewModeStretch},
	}
	//	{"Ideal", walk.ImageViewModeIdeal},
	//	{"Shrink", walk.ImageViewModeShrink},

	const defaultImageViewMode = walk.ImageViewModeCorner

	// ■ mainWindow
	mainWindow, err := walk.NewMainWindow()
	if err != nil {
		panic(err)
	}
	mainWindow.SetTitle("002-Widgets")
	mainWindow.SetSize(walk.Size{600, 400})
	mainWindow.SetMinMaxSize(walk.Size{300, 200}, walk.Size{})
	//mainWindow.SetBackground(rgbBrush(240, 128, 128))

	mainWindowLayout := walk.NewHBoxLayout()
	mainWindowLayout.SetMargins(walk.Margins{8, 2, 8, 8})
	mainWindow.SetLayout(mainWindowLayout)

	// ■ mainWindow/tabWidget
	tabWidget, err := walk.NewTabWidget(mainWindow)
	if err != nil {
		panic(err)
	}
	//tabWidget.SetBackground(rgbBrush(240, 64, 64))

	// ■ mainWindow/tabWidget/page1
	page1, err := walk.NewTabPage()
	if err != nil {
		panic(err)
	}
	page1.SetTitle("文字表示")
	page1Layout := walk.NewVBoxLayout()
	page1Layout.SetMargins(walk.Margins{})
	page1.SetLayout(page1Layout)

	if err := tabWidget.Pages().Add(page1); err != nil {
		panic(err)
	}

	// ■ mainWindow/tabWidget/page1/textEdit
	textEdit, err := walk.NewTextEdit(page1)
	if err != nil {
		panic(err)
	}
	textEdit.SetReadOnly(false)
	textEdit.SetFont(createFont("Noto Sans JP", 10, 0))
	//textEdit.SetBackground(rgbBrush(200, 220, 240))

	// ■ mainWindow/tabWidget/page2
	page2, err := walk.NewTabPage()
	if err != nil {
		panic(err)
	}
	page2.SetTitle("画像表示")
	page2Layout := walk.NewVBoxLayout()
	page2Layout.SetMargins(walk.Margins{})
	page2.SetLayout(page2Layout)

	if err := tabWidget.Pages().Add(page2); err != nil {
		panic(err)
	}

	// ■ mainWindow/tabWidget/page2/imageView
	imageView, err := walk.NewImageView(page2)
	if err != nil {
		panic(err)
	}
	imageView.SetMode(defaultImageViewMode)
	//imageView.SetClearsBackground(true)
	//imageView.SetPaintMode(walk.PaintBuffered)

	// 初期選択
	tabWidget.SetCurrentIndex(0)

	// ■ mainWindow/panel
	panel, err := walk.NewComposite(mainWindow)
	if err != nil {
		panic(err)
	}
	//panel.SetBackground(rgbBrush(220, 240, 200))
	panel.SetMinMaxSize(
		walk.Size{Width: 88, Height: 0},	// Heightは無視される
		walk.Size{Width: 88, Height: 0},
	)

	panelLayout := walk.NewVBoxLayout()
	panelLayout.SetMargins(walk.Margins{0, 0, 0, 0})
	panelLayout.SetSpacing(2)
	panel.SetLayout(panelLayout)

	// ■ mainWindow/panel/縦スペーサー(高さ16ピクセル)
	createVSpacer(panel, 16)

	// ■ mainWindow/panel/groupBox
	groupBox, err := walk.NewGroupBox(panel)
	if err != nil {
		panic(err)
	}
	groupBox.SetTitle("ImageView")

	groupBoxLayout := walk.NewVBoxLayout()
	groupBoxLayout.SetMargins(walk.Margins{4, 0, 0, 0}) // 左, 上, 右, 下
	groupBoxLayout.SetSpacing(0)
	groupBox.SetLayout(groupBoxLayout)


	// ■ mainWindow/panel/groupBox/縦スペーサー(高さ4)
	createVSpacer(groupBox, 4)

	// ■ mainWindow/panel/groupBox/ラジオボタン
	for _, item := range imageViewModes {
		checked := (item.Mode == defaultImageViewMode)
		rb := createRadioButton(groupBox, item.Name, item.Mode, checked)
		rb.CheckedChanged().Attach(func() {
			if rb.Checked() {
				imageView.SetMode(item.Mode)
			}
		})
	}
	
	// ■ mainWindow/panel/groupBox/横スペーサー(伸び縮み)
	createHSpacer(groupBox, 0)
	// ↑メモ
	// RadioButtonは幅が伸縮しないので、それだけを含むGroupBoxは幅が伸縮しない
	// HSpacerは幅が伸縮するので、これを入れるとGroupBoxの幅が伸縮するようになる

	// ■ mainWindow/panel/画像表示ボタン
	createPushButton(panel, "画像表示", func() {
		const n = len(imageFiles)
		if isCtrlPressed() {
			imageFileIndex = ((imageFileIndex-1)%n + n) % n
		} else {
			imageFileIndex = (imageFileIndex + 1) % n
		}
		newImage := loadImage(imageFiles[imageFileIndex], mainWindow.DPI())
		setImage(imageView, newImage)
		tabWidget.SetCurrentIndex(1)
	})

	// ■ mainWindow/panel/リストボックス
	listBoxItems := []string{"Item 1", "Item 2"}
	listBox, err := walk.NewListBox(panel)
	if (err != nil) {
		panic(err)
	}
	listBox.SetModel(listBoxItems)

	// ■ mainWindow/panel/テキスト表示ボタン
	createPushButton(panel, "テキスト表示", func() {
		text := fmt.Sprintf(""+
			"Window Size = %#v\r\n"+
			"Client Size = %#v\r\n"+
			"Panel Size = %#v\r\n"+
			"GroupBox Size = %#v\r\n"+
			"imageView.PaintMode = %s\r\n"+
			"imageView.Mode = %s\r\n"+
			"",
			mainWindow.Size(),
			mainWindow.ClientBounds().Size(),
			panel.Size(),
			groupBox.Size(),
			toString(imageView.PaintMode()),
			toString(imageView.Mode()),
		)
		textEdit.SetText(text)
		tabWidget.SetCurrentIndex(0)
	})

	// ■ mainWindow/panel/消去ボタン
	createPushButton(panel, "消去", func() {
		textEdit.SetText("")
		setImage(imageView, nil)
		imageFileIndex = -1
	})

	// ■ mainWindow/panel/縦スペーサー(自由に伸びる)
	createVSpacer(panel, 0)

	// ■ mainWindow/panel/終了ボタン
	createPushButton(panel, "終了", func() {
		mainWindow.Close()
	})

	//panel.SetFocus()
	mainWindow.Show()
	mainWindow.Run()
}

// ImageViewに新しいImageを表示し、古いImageは破棄する
func setImage(iv *walk.ImageView, img walk.Image) {
	oldImage := iv.Image()
	if err := iv.SetImage(img); err != nil {
		panic(err)
	}
	if oldImage != nil {
		oldImage.Dispose()
	}
}

// プッシュボタンを作成
func createPushButton(parent walk.Container, text string, onClick func()) *walk.PushButton {
	btn, err := walk.NewPushButton(parent)
	if err != nil {
		panic(err)
	}
	btn.SetText(text)
	btn.Clicked().Attach(onClick)
	return btn
}

// ラジオボタンを作成
func createRadioButton(parent walk.Container, text string, value any, checked bool) *walk.RadioButton {
	btn, err := walk.NewRadioButton(parent)
	if err != nil {
		panic(err)
	}
	btn.SetText(text)
	btn.SetValue(value)
	btn.SetChecked(checked)
	return btn
}

// 縦スペーサーを作成
// height > 0  なら指定した高さに固定
// height <= 0 なら自由に伸縮する
func createVSpacer(parent walk.Container, height int) *walk.Spacer {
	var spacer *walk.Spacer
	var err error

	if height > 0 {
		spacer, err = walk.NewVSpacerFixed(parent, height)
	} else {
		spacer, err = walk.NewVSpacer(parent)
	}
	if err != nil {
		panic(err)
	}
	return spacer
}

// 横スペーサーを作成
// width > 0  なら指定した幅に固定
// width <= 0 なら自由に伸縮する
func createHSpacer(parent walk.Container, width int) *walk.Spacer {
	var spacer *walk.Spacer
	var err error

	if width > 0 {
		spacer, err = walk.NewHSpacerFixed(parent, width)
	} else {
		spacer, err = walk.NewHSpacer(parent)
	}
	if err != nil {
		panic(err)
	}
	return spacer
}

// 各種定数の名前
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
			return fmt.Sprintf("(%T) %v", v, v)
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
			return fmt.Sprintf("(%T) %v", v, v)
		}

	case nil:
		return "nil"

	default:
		return fmt.Sprintf("(%T) %v", v, v)
	}
}

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
