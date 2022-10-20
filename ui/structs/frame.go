package structs

import (
	"strings"
)

type Frame struct {
	BaseWidget

	orientation FrameOrientation
}

func CreateFrame(orientation FrameOrientation) *Frame {
	var widgets []Widget
	frame := &Frame{
		BaseWidget: BaseWidget{
			widgetType: frameWidget,

			needsRepaint: true,
			widgets:      widgets,

			backgroundColor: "#fff",
		},
		orientation: orientation,
	}
	return frame
}

func (frame *Frame) SetBackgroundColor(backgroundColor string) {
	if len(backgroundColor) > 0 && strings.HasPrefix(backgroundColor, "#") {
		frame.backgroundColor = backgroundColor
		frame.needsRepaint = true
	}
}

func (frame *Frame) SetWidth(width float64) {
	frame.box.width = width
	frame.fixedWidth = true
	frame.RequestReflow()
}

func (frame *Frame) GetWidth() float64 {
	return frame.box.width
}

func (frame *Frame) SetHeight(height float64) {
	frame.box.height = height
	frame.fixedHeight = true
	frame.RequestReflow()
}

func (frame *Frame) GetHeight() float64 {
	return frame.box.height
}

func (frame *Frame) draw() {
	orientation := frame.orientation
	top, left, width, height := frame.computedBox.GetCoords()
	window := frame.window
	context := window.GetContext()
	context.SetHexColor(frame.backgroundColor)
	context.DrawRectangle(left, top, width, height)
	context.Fill()

	CopyWidgetToBuffer(frame, context.GetImage())

	widgets := frame.widgets
	childrenLen := len(widgets)
	if childrenLen > 0 {
		childrenWidgets := GetCoreWidgets(widgets)
		childrenLayout := CalculateChildrenWidgetsLayout(childrenWidgets, top, left, width, height, orientation)

		for idx, child := range widgets {
			child.ComputedBox().SetCoords(childrenLayout[idx].box.GetCoords())
			child.draw()
		}
	}
}
