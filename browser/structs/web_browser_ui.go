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
		debugger := webBrowser.DebuggerMap[currentDocument]
		if debugger != nil {

			window.AddContextMenuEntry("关闭调试", func() {
				window.RemoveStaticOverlay("debugOverlay")
				debugger.DebugFlag = false

				if debugger.DebugWindow != nil {
					app.DestroyWindow(debugger.DebugWindow)
					debugger.DebugWindow = nil
					debugger.DebugTree = nil
				}
				webBrowser.DebuggerMap[currentDocument] = nil
			})

			if debugger.DebugWindow != nil {

				window.AddContextMenuEntry("隐藏DOM树", func() {
					app.DestroyWindow(debugger.DebugWindow)
					debugger.DebugWindow = nil
					debugger.DebugTree = nil
				})

			} else {

				window.AddContextMenuEntry("显示DOM树", func() {
					tree := ui_structs.CreateTreeWidget()
					tree.SetFontSize(14)

					debugger.DebugWindow = ui_structs.CreateWindow("HTML结构树", 600, 800, true)
					debugger.DebugTree = tree

					frame := ui_structs.CreateFrame(ui_structs.HorizontalFrame)
					frame.AddWidget(tree)

					debugger.DebugWindow.RegisterTree(tree)
					debugger.DebugWindow.SetRootFrame(frame)
					debugger.DebugWindow.Show()

					app.AddWindow(debugger.DebugWindow)

					treeNodeDOM := treeNodeFromDOM(currentDocument.DOM)
					tree.SetSelectCallback(func(selectedNode *ui_structs.TreeWidgetNode) {
						if debugger.DebugFlag {
							child, _ := currentDocument.DOM.FindByXPath(selectedNode.Value)
							debugger.SelectedElement = child
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
				webBrowser.DebuggerMap[currentDocument] = CreateDebugger(true)
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
