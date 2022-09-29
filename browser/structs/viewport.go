package structs

import (
	"layout"
	"log"
	profiler_structs "profiler/structs"
	renderer_structs "renderer/structs"
	ui_structs "ui/structs"
)

type Viewport struct {
	Frame     *ui_structs.Frame
	View      *ui_structs.CanvasWidget
	ScrollBar *ui_structs.ScrollBarWidget
}

func CreateViewport(webBrowser *WebBrowser) *Viewport {
	frame := ui_structs.CreateFrame(ui_structs.VerticalFrame)
	view := ui_structs.CreateCanvasWidget(GetViewportRenderer(webBrowser))
	scrollBar := ui_structs.CreateScrollBarWidget(ui_structs.VerticalScrollBar)
	scrollBar.SetTrackColor("#ccc")
	scrollBar.SetThumbColor("#aaa")
	scrollBar.SetWidth(12)

	frame.AddWidget(view)
	frame.AddWidget(scrollBar)
	return &Viewport{
		Frame:     frame,
		View:      view,
		ScrollBar: scrollBar,
	}
}

func GetViewportRenderer(webBrowser *WebBrowser) func(*ui_structs.CanvasWidget) {
	return func(canvas *ui_structs.CanvasWidget) {
		go func() {
			profiler := webBrowser.Profiler
			document := webBrowser.CurrentDocument
			ui := webBrowser.UI
			headBar := ui.HeadBar
			viewport := ui.Viewport

			profiler.Start("render")
			ctxBounds := canvas.GetContext().GetImage().Bounds()
			drawingContext := renderer_structs.CreateContext(ctxBounds.Max.X, ctxBounds.Max.Y)

			err := layout.LayoutDocument(drawingContext, document)
			if err != nil {
				log.Fatal("render", "Can't render page: "+err.Error())
			}

			canvas.SetContext(drawingContext)
			canvas.RequestReflow()
			profiler.Stop("render")

			statusLabel := headBar.StatusLabel
			statusLabel.SetContent(createStatusLabel(profiler))
			statusLabel.RequestReflow()
			canvas.RequestReflow()

			body, err := webBrowser.CurrentDocument.DOM.FindChildByName("body")
			if err != nil {
				log.Fatal("render", "can't find body element: "+err.Error())
				return
			}

			scrollBar := viewport.ScrollBar
			scrollBar.SetScrollerSize(body.RenderBox.Height)
			scrollBar.RequestReflow()
		}()
	}
}

func createStatusLabel(perf *profiler_structs.Profiler) string {
	return "Loaded; " +
		"Render took: " + perf.GetProfile("render").GetElapsedTime().String() + "; "
}
