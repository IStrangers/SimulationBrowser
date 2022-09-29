package structs

import ui_structs "ui/structs"

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
