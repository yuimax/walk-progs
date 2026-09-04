//go:generate goversioninfo
package main

import (
	"fmt"
	"github.com/lxn/walk"
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
	mainWindow.SetMinMaxSize(walk.Size{320, 240}, walk.Size{})
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

	panelLayout := walk.NewVBoxLayout()
	panelLayout.SetMargins(walk.Margins{0, 0, 0, 0})
	panelLayout.SetSpacing(2)
	panel.SetLayout(panelLayout)

	// ■ mainWindow/panel/縦スペーサー(高さ16ピクセル)
	createVSpacer(panel, 16)

	// ■ mainWindow/panel/group1
	group1, err := walk.NewGroupBox(panel)
	if err != nil {
		panic(err)
	}
	group1.SetTitle("ImageView")
	group1Layout := walk.NewVBoxLayout()
	group1Layout.SetMargins(walk.Margins{8, 0, 0, 0})
	group1Layout.SetSpacing(0)
	group1.SetLayout(group1Layout)

	// ■ mainWindow/panel/group1/縦スペーサー(高さ4)
	createVSpacer(group1, 4)

	// ■ mainWindow/panel/group1/ラジオボタン
	for _, item := range imageViewModes {
		checked := (item.Mode == defaultImageViewMode)
		rb := createRadioButton(group1, item.Name, item.Mode, checked)
		rb.CheckedChanged().Attach(func() {
			if rb.Checked() {
				imageView.SetMode(item.Mode)
			}
		})
	}

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

	// ■ mainWindow/panel/テキスト表示ボタン
	createPushButton(panel, "テキスト表示", func() {
		text := fmt.Sprintf(""+
			"Window Size = %#v\r\n"+
			"Client Size = %#v\r\n"+
			"imageView.PaintMode = %s\r\n"+
			"imageView.Mode = %s\r\n"+
			"",
			mainWindow.Size(),
			mainWindow.ClientBounds().Size(),
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
