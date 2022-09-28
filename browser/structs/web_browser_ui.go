package structs

import (
	"layout"
	renderer_structs "renderer/structs"
	ui_structs "ui/structs"
)

type WebBrowserUI struct {
	Viewport *ui_structs.CanvasWidget
}

func CreateWebBrowserUI(webBrowser *WebBrowser) *WebBrowserUI {
	window := webBrowser.Window

	rootFrame := ui_structs.CreateFrame(ui_structs.HorizontalFrame)
	viewArea := ui_structs.CreateFrame(ui_structs.VerticalFrame)

	viewport := ui_structs.CreateCanvasWidget(GetViewportRenderer(webBrowser))
	scrollBar := ui_structs.CreateScrollBarWidget(ui_structs.VerticalScrollBar)
	viewArea.AddWidget(viewport)
	viewArea.AddWidget(scrollBar)

	rootFrame.AddWidget(viewArea)
	window.SetRootFrame(rootFrame)

	return &WebBrowserUI{

		Viewport: viewport,
	}
}

func GetViewportRenderer(webBrowser *WebBrowser) func(*ui_structs.CanvasWidget) {
	return func(canvas *ui_structs.CanvasWidget) {
		go func() {
			profiler := webBrowser.Profiler
			document := webBrowser.CurrentDocument
			profiler.Start("render")
			ctxBounds := canvas.GetContext().GetImage().Bounds()
			drawingContext := renderer_structs.CreateContext(ctxBounds.Max.X, ctxBounds.Max.Y)

			err := layout.LayoutDocument(drawingContext, document)
			if err != nil {
				println("render", "Can't render page: "+err.Error())
			}

			canvas.SetContext(drawingContext)
			canvas.RequestReflow()
			profiler.Stop("render")

		}()
	}
}
