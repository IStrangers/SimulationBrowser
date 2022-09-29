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
	rootFrame.AddWidget(headBar.Frame)

	viewport := CreateViewport(webBrowser)
	rootFrame.AddWidget(viewport.Frame)

	window.SetRootFrame(rootFrame)

	registerUIWidget(webBrowser,headBar)

	return &WebBrowserUI{
		HeadBar:  headBar,
		Viewport: viewport,
	}
}

func registerUIWidget(webBrowser *WebBrowser,headBar *HeadBar) {
	window := webBrowser.Window

	urlInput := headBar.UrlInput

	window.RegisterInput(urlInput)
	urlInput.SetReturnCallback(func() {
		loadDocumentByUrl(webBrowser)
	})

	window.RegisterButton(headBar.ToolsButton, func() {

		window.AddContextMenuEntry("首页", func() {
			urlInput.SetValue(WebBrowserName + "://HomePage")
			loadDocumentByUrl(webBrowser)
		})

		window.AddContextMenuEntry("历史记录", func() {
			urlInput.SetValue(WebBrowserName + "://History")
			loadDocumentByUrl(webBrowser)
		})

		window.AddContextMenuEntry("关于", func() {
			urlInput.SetValue(WebBrowserName + "://About")
			loadDocumentByUrl(webBrowser)
		})

		app := webBrowser.App
		currentDocument := webBrowser.CurrentDocument
		if currentDocument.DebugFlag {

			window.AddContextMenuEntry("关闭调试", func() {
				window.RemoveStaticOverlay("debugOverlay")
				currentDocument.DebugFlag = false

				if currentDocument.DebugWindow != nil {
					app.DestroyWindow(currentDocument.DebugWindow)
					currentDocument.DebugWindow = nil
					currentDocument.DebugTree = nil
				}
			})

			if currentDocument.DebugWindow != nil {

				window.AddContextMenuEntry("隐藏DOM树", func() {
					app.DestroyWindow(currentDocument.DebugWindow)
					currentDocument.DebugWindow = nil
					currentDocument.DebugTree = nil
				})

			} else {

				window.AddContextMenuEntry("显示DOM树", func() {
					tree := ui_structs.CreateTreeWidget()
					tree.SetFontSize(14)

					currentDocument.DebugWindow = ui_structs.CreateWindow("HTML结构树", 600, 800, true)
					currentDocument.DebugTree = tree

					frame := ui_structs.CreateFrame(ui_structs.HorizontalFrame)
					frame.AddWidget(tree)

					currentDocument.DebugWindow.RegisterTree(tree)
					currentDocument.DebugWindow.SetRootFrame(frame)
					currentDocument.DebugWindow.Show()

					app.AddWindow(currentDocument.DebugWindow)

					treeNodeDOM := treeNodeFromDOM(currentDocument.DOM)
					tree.SetSelectCallback(func(selectedNode *ui_structs.TreeWidgetNode) {
						if currentDocument.DebugFlag {
							child, _ := currentDocument.DOM.FindByXPath(selectedNode.Value)
							currentDocument.SelectedElement = child
							showDebugOverlay(webBrowser)
						}
					})

					tree.RemoveNodes()
					tree.AddNode(treeNodeDOM)
					tree.RequestRepaint()
				})

			}

		} else {

			window.AddContextMenuEntry("开启调试", func() {
				currentDocument.DebugFlag = true
			})

		}

		window.DrawContextMenu()
	})

	window.RegisterButton(headBar.ReloadButton, func() {
		loadDocumentByUrl(webBrowser)
	})

	window.RegisterButton(headBar.NextButton, func() {
		history := webBrowser.History
		if len(history.NextPages()) <= 0 {
			return
		}
		history.PopNext()
		urlInput.SetValue(history.Last().String())
		loadDocumentByUrl(webBrowser)
	})

	window.RegisterButton(headBar.PreviousButton, func() {
		history := webBrowser.History
		if history.PageCount() <= 1 {
			return
		}
		history.Pop()
		urlInput.SetValue(history.Last().String())
		loadDocumentByUrl(webBrowser)
	})
}
