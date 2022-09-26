package structs

import (
	browser "browser/structs"
	"image"
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

			backgroundColor: "#ffffff",
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
	context := window.getContext()
	context.SetColor(frame.backgroundColor)
	context.DrawRectangle(float64(left), float64(top), float64(width), float64(height))
	context.Fill()

	CopyWidgetToBuffer(frame, context.Image())

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
