package structs

import (
	ui_structs "ui/structs"
)

type WebBrowserUI struct {
	HeadBar  *HeadBar
	Viewport *Viewport
}

func CreateWebBrowserUI(webBrowser *WebBrowser) *WebBrowserUI {
	window := webBrowser.Window

	rootFrame := ui_structs.CreateFrame(ui_structs.HorizontalFrame)

	headBar := CreateHeadBar()
	window.RegisterInput(headBar.UrlInput)
	rootFrame.AddWidget(headBar.Frame)

	viewport := CreateViewport(webBrowser)
	rootFrame.AddWidget(viewport.Frame)

	window.SetRootFrame(rootFrame)

	return &WebBrowserUI{
		HeadBar:  headBar,
		Viewport: viewport,
	}
}
