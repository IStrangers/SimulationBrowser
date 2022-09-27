package structs

import (
	"assets"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/freetype/truetype"
)

type ScrollBarWidget struct {
	BaseWidget

	orientation ScrollBarOrientation
	selected bool
	thumbSize float64
	thumbColor string

	scrollerSize float64
	scrollerOffset float64
}

func CreateScrollBarWidget(orientation ScrollBarOrientation) *ScrollBarWidget {
	var widgets []Widget
	font,_ := truetype.Parse(assets.OpenSans(400))

	return &ScrollBarWidget {
		BaseWidget: BaseWidget{

			needsRepaint: true,
			widgets: widgets,

			widgetType: scrollbarWidget,

			cursor: glfw.CreateStandardCursor(glfw.ArrowCursor),

			backgroundColor: "#fff",

			font: font,
		},
		orientation: orientation,
	}
}

func (scrollBar *ScrollBarWidget) SetWidth(width float64)  {
	scrollBar.box.width = width
	scrollBar.fixedWidth = true
	scrollBar.RequestReflow()
}

func (scrollBar *ScrollBarWidget) SetHeight(height float64) {
	scrollBar.box.height = height
	scrollBar.fixedHeight = true
	scrollBar.RequestReflow()
}

func (scrollBar *ScrollBarWidget) SetTrackColor(backgroundColor string) {
	if len(backgroundColor) > 0 && string(backgroundColor[0]) == "#" {
		scrollBar.backgroundColor = backgroundColor
		scrollBar.needsRepaint = true
	}
}

func (scrollBar *ScrollBarWidget) SetScrollerSize(scrollerSize float64) {
	scrollBar.scrollerSize = scrollerSize
	scrollBar.needsRepaint = true
}

func (scrollBar *ScrollBarWidget) SetThumbSize(thumbSize float64) {
	scrollBar.thumbSize = thumbSize
	scrollBar.needsRepaint = true
}

func (scrollBar *ScrollBarWidget) SetThumbColor(thumbColor string) {
	scrollBar.thumbColor = thumbColor
	scrollBar.needsRepaint = true
}

func (scrollBar *ScrollBarWidget) SetScrollerOffset(scrollerOffset float64) {
	scrollBar.scrollerOffset = scrollerOffset
	scrollBar.needsRepaint = true
}

func (scrollBar *ScrollBarWidget) draw() {
	context := scrollBar.window.context
	computedBox := scrollBar.computedBox

	top,left,width,height := computedBox.top,computedBox.left,computedBox.width,computedBox.height

	context.SetHexColor(scrollBar.backgroundColor)
	context.DrawRectangle(left,top,width,height)
	context.Fill()

	if scrollBar.scrollerSize > height {
		thumbSize := height * (height / scrollBar.scrollerSize)
		thumbOffset := scrollBar.scrollerOffset

		scrollJump := (scrollBar.scrollerSize - height) / (height - thumbSize)

		context.SetHexColor(scrollBar.thumbColor)
		context.DrawRectangle(left + 1,top - (thumbOffset / scrollJump),width - 2,thumbSize)
		context.Fill()
	}

	CopyWidgetToBuffer(scrollBar,context.GetImage())
}