package GUI

import (
	"ImageTextConverter/GUI/Icon"
	image2text "ImageTextConverter/Image2Text"
	text2image "ImageTextConverter/Text2Image"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/widget"
	"strings"
)

const (
	Image = iota
	Text
)

type myApp struct {
	App fyne.App
}

func NewApp() myApp {
	return myApp{App: app.New()}
}
func openfile(parent *fyne.Window, fileType int) func() {
	return func() {
		var uri string
		myDialog := dialog.NewFileOpen(func(url fyne.URIReadCloser, err error) {
			fmt.Println(url)
			if url != nil {
				uri = url.URI().String()
				wholepath := strings.Split(uri, "//")[1]
				path := wholepath[:strings.LastIndex(wholepath, `\`)]
				filename := strings.Split(wholepath[strings.LastIndex(wholepath, `\`)+1:], ".")[0]
				fmt.Println(path, filename)
				var err error
				switch fileType {
				case 0:
					err = image2text.PathConvert2Text(wholepath, fmt.Sprintf("%s\\%s.%s", path, filename, "txt"))
				case 1:
					err = text2image.PathConvert2Image(wholepath, fmt.Sprintf("%s\\%s.%s", path, filename, "png"))
				}
				if err != nil {
					dialog.ShowError(err, *parent)
				} else {
					dialog.ShowInformation("保存成功", fmt.Sprintf("成功将%s转格式并将同名文件保存在%s中", filename, path), *parent)
				}
			}
		}, *parent)
		myDialog.SetDismissText("Choose a txt file to convert")
		FilterArray := map[int][]string{0: []string{".png"}, 1: []string{".txt"}}
		myDialog.SetFilter(storage.NewExtensionFileFilter(FilterArray[fileType]))
		//myDialog.SetOnClosed(func(){
		//	fmt.Println("test")
		//	//*parent.
		//})
		myDialog.Show()
	}
}
func (a *myApp) Run() {
	myWindow := a.App.NewWindow("ImageTextConverter")
	myWindow.SetIcon(Icon.ConvertIcon())
	myWindow.Resize(fyne.Size{
		Width:  1000,
		Height: 600,
	})
	toImageBtn := widget.NewButtonWithIcon("选择文档", Icon.TextFileIcon(), openfile(&myWindow, Text))
	toTextBtn := widget.NewButtonWithIcon("选择图片", Icon.ImageFileIcon(), openfile(&myWindow, Image))
	mainContainer := fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(2), toImageBtn, toTextBtn)
	myWindow.SetContent(mainContainer)
	myWindow.ShowAndRun()
}
