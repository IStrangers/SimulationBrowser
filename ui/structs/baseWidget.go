package structs

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/freetype/truetype"
	"image"
)

type BaseWidget struct {
	box            Box
	computedBox    Box
	widgetPosition WidgetPosition

	font *truetype.Font

	needsRepaint bool
	fixedWidth   bool
	fixedHeight  bool

	widgets []Widget

	backgroundColor string

	widgetType WidgetType
	cursor     *glfw.Cursor

	focusable  bool
	selectable bool

	focused  bool
	selected bool

	buffer *image.RGBA
	window *Window
}

func GetCoreWidgets(widgets []Widget) []*BaseWidget {
	var coreWidgets []*BaseWidget
	for _, widget := range widgets {
		coreWidgets = append(coreWidgets, widget.GetBaseWidget())
	}
	return coreWidgets
}

func CalculateChildrenWidgetsLayout(children []*BaseWidget, top, left, width, height float64, orientation FrameOrientation) []*BaseWidget {
	var childrenLayout []*BaseWidget

	childrenLen := len(children)
	for i := 0; i < childrenLen; i++ {
		currentLayout := &BaseWidget{}

		if orientation == VerticalFrame {
			if i > 0 {
				prev := childrenLayout[i-1]
				currentLayout.box.left = prev.box.left + prev.box.width
			} else {
				currentLayout.box.left = left
			}

			if children[i].fixedWidth {
				currentLayout.box.width = children[i].box.width
			} else {
				fixedWidthElements, volatileWidthElements := GetFixedWidthElements(children)
				remainingWidth := CalculateFlexibleWidth(width, fixedWidthElements)
				currentLayout.box.width = remainingWidth / float64(len(volatileWidthElements))
			}

			currentLayout.box.top = top
			currentLayout.box.height = height
		} else if orientation == HorizontalFrame {
			if i > 0 {
				prev := childrenLayout[i-1]
				currentLayout.box.top = prev.box.top + prev.box.height
			} else {
				currentLayout.box.height = top
			}

			if children[i].fixedHeight {
				currentLayout.box.height = children[i].box.height
			} else {
				fixedHeightElements, volatileHeightElements := GetFixedHeightElements(children)
				remainingHeight := CalculateFlexibleHeight(height, fixedHeightElements)
				currentLayout.box.height = remainingHeight / float64(len(volatileHeightElements))
			}

			currentLayout.box.left = left
			currentLayout.box.width = width
		} else {
			continue
		}
		childrenLayout = append(childrenLayout, currentLayout)
	}

	return childrenLayout
}

func GetFixedWidthElements(elements []*BaseWidget) ([]*BaseWidget, []*BaseWidget) {
	var fixedWidthElements []*BaseWidget
	var volatileWidthElements []*BaseWidget

	for _, element := range elements {
		if element.fixedWidth {
			fixedWidthElements = append(fixedWidthElements, element)
		} else {
			volatileWidthElements = append(volatileWidthElements, element)
		}
	}

	return fixedWidthElements, volatileWidthElements
}

func GetFixedHeightElements(elements []*BaseWidget) ([]*BaseWidget, []*BaseWidget) {
	var fixedHeightElements []*BaseWidget
	var volatileHeightElements []*BaseWidget

	for _, element := range elements {
		if element.fixedHeight {
			fixedHeightElements = append(fixedHeightElements, element)
		} else {
			volatileHeightElements = append(volatileHeightElements, element)
		}
	}
	return fixedHeightElements, volatileHeightElements
}

func CalculateFlexibleWidth(availableWidth float64, elements []*BaseWidget) float64 {
	for _, el := range elements {
		availableWidth = availableWidth - el.box.width
	}

	if availableWidth < 0 {
		return 0
	}

	return availableWidth
}

func CalculateFlexibleHeight(availableHeight float64, elements []*BaseWidget) float64 {
	for _, el := range elements {
		availableHeight = availableHeight - el.box.height
	}

	if availableHeight < 0 {
		return 0
	}

	return availableHeight
}

func (widget *BaseWidget) Buffer() *image.RGBA {
	return widget.buffer
}

func (widget *BaseWidget) NeedsRepaint() bool {
	return widget.needsRepaint
}

func (widget *BaseWidget) Widgets() []Widget {
	return widget.widgets
}

func (widget *BaseWidget) ComputedBox() *Box {
	return &widget.computedBox
}

func (widget *BaseWidget) GetBaseWidget() *BaseWidget {
	return widget
}

func (widget *BaseWidget) RequestReflow() {
	if widget.window != nil {
		widget.window.SetNeedsReflow(true)
	}
}

func (widget *BaseWidget) SetNeedsRepaint(needsRepaint bool) {
	widget.needsRepaint = needsRepaint
}

func (widget *BaseWidget) SetWindow(window *Window) {
	widget.window = window

	for _, childWidget := range widget.widgets {
		childWidget.SetWindow(window)
	}
}

func (widget *BaseWidget) SetBuffer(buffer *image.RGBA) {
	widget.buffer = buffer
}


func (widget *BaseWidget) AddWidget(wd Widget) {
	wd.SetWindow(widget.window)
	widget.widgets = append(widget.widgets,wd)

	if widget.widgets != nil && widget.window.rootFrame != nil {
		widget.window.rootFrame.RequestReflow()
	}
}