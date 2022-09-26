package structs

import (
	browser "browser/structs"
	"image"
	"image/draw"
)

type Widget interface {
	Buffer() *image.RGBA
	SetBuffer(rgba *image.RGBA)
	SetNeedsRepaint(bool)
	NeedsRepaint() bool
	Widgets() []Widget
	ComputedBox() *Box
	SetWindow(window *browser.Window)
	BaseWidget() *BaseWidget

	draw()
}

func GetCoreWidgets(widgets []Widget) []*BaseWidget {
	var coreWidgets []*BaseWidget
	for _, widget := range widgets {
		coreWidgets = append(coreWidgets, widget.BaseWidget())
	}
	return coreWidgets
}

func CopyWidgetToBuffer(widget Widget, src image.Image) {
	computedBox := widget.ComputedBox()
	top, left, width, height := int(computedBox.top), int(computedBox.left), int(computedBox.width), int(computedBox.height)

	buffer := widget.Buffer()
	if buffer == nil || buffer.Bounds().Max.X != width && buffer.Bounds().Max.Y != height {
		widget.BaseWidget().SetBuffer(image.NewRGBA(image.Rectangle{
			Min: image.Point{},
			Max: image.Point{
				X: width,
				Y: height,
			},
		}))
	}

	draw.Draw(widget.BaseWidget().buffer, image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: width, Y: height},
	}, src, image.Point{X: left, Y: top}, draw.Over)
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
