package structs

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	renderer_structs "renderer/structs"
)

type CanvasWidget struct {
	BaseWidget

	context        *renderer_structs.Context
	drawingContext *renderer_structs.Context

	renderer func(widget *CanvasWidget)

	scrollable bool
	offset     int

	drawingRepaint bool
}

func CreateCanvasWidget(renderer func(*CanvasWidget)) *CanvasWidget {
	var widgets []Widget

	return &CanvasWidget{
		BaseWidget: BaseWidget{
			needsRepaint: true,
			widgets:      widgets,

			widgetType: canvasWidget,

			cursor: glfw.CreateStandardCursor(glfw.ArrowCursor),

			backgroundColor: "#fff",
		},

		context:        renderer_structs.CreateContext(0, 0),
		drawingContext: renderer_structs.CreateContext(0, 0),
		renderer:       renderer,
		drawingRepaint: true,
	}
}

func (canvas *CanvasWidget) SetWidth(width float64) {
	canvas.box.width = width
	canvas.fixedWidth = true
	canvas.RequestReflow()
}

func (canvas *CanvasWidget) SetHeight(height float64) {
	canvas.box.height = height
	canvas.fixedHeight = true
	canvas.RequestReflow()
}

func (canvas *CanvasWidget) EnableScrolling() {
	canvas.scrollable = true
}

func (canvas *CanvasWidget) DisableScrolling() {
	canvas.scrollable = false
	canvas.offset = 0
}

func (canvas *CanvasWidget) SetOffset(offset int) {
	canvas.offset = offset
}

func (canvas *CanvasWidget) GetOffset() int {
	return canvas.offset
}

func (canvas *CanvasWidget) GetContext() *renderer_structs.Context {
	return canvas.drawingContext
}

func (canvas *CanvasWidget) SetContext(ctx *renderer_structs.Context) {
	canvas.drawingContext = ctx
}

func (canvas *CanvasWidget) SetDrawingRepaint(repaint bool) {
	canvas.drawingRepaint = repaint
}

func (canvas *CanvasWidget) draw() {
	context := canvas.window.context
	top, left, width, height := canvas.computedBox.GetCoords()
	currentContextSize := canvas.context.GetImage().Bounds().Size()

	if currentContextSize.X != int(width) || currentContextSize.Y != int(height) {
		canvas.context = renderer_structs.CreateContext(int(width), int(height))
		canvas.drawingContext = renderer_structs.CreateContext(int(width), 12000)
		canvas.drawingRepaint = true
	}

	if canvas.drawingRepaint {
		canvas.renderer(canvas)
		canvas.drawingRepaint = false
	}

	canvas.context.DrawImage(canvas.drawingContext.GetImage(), int(left), canvas.offset)
	context.DrawImage(canvas.context.GetImage(), int(left), int(top))
	CopyWidgetToBuffer(canvas, context.GetImage())
}
