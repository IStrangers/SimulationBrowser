package structs

import (
	"SimulationBrowser/assets"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/freetype/truetype"
)

type LabelWidget struct {
	BaseWidget
	content string

	fontSize  float64
	fontColor string
}

func CreateLabelWidget(content string) *LabelWidget {
	var widgets []Widget
	font, _ := truetype.Parse(assets.OpenSans(400))

	return &LabelWidget{
		BaseWidget: BaseWidget{

			needsRepaint: true,
			widgets:      widgets,

			widgetType: labelWidget,

			cursor: glfw.CreateStandardCursor(glfw.ArrowCursor),

			backgroundColor: "#fff",

			font: font,
		},
		content: content,

		fontSize:  20,
		fontColor: "#000",
	}
}

func (label *LabelWidget) SetWidth(width float64) {
	label.box.width = width
	label.fixedWidth = true
	label.RequestReflow()
}

func (label *LabelWidget) SetHeight(height float64) {
	label.box.height = height
	label.fixedHeight = true
	label.RequestReflow()
}

func (label *LabelWidget) SetFontSize(fontSize float64) {
	label.fontSize = fontSize
	label.needsRepaint = true
}

func (label *LabelWidget) SetContent(content string) {
	label.content = content
	label.needsRepaint = true
}

func (label *LabelWidget) GetContent() string {
	return label.content
}

func (label *LabelWidget) SetFontColor(fontColor string) {
	if len(fontColor) > 0 && string(fontColor[0]) == "#" {
		label.fontColor = fontColor
		label.needsRepaint = true
	}
}

func (label *LabelWidget) SetBackgroundColor(backgroundColor string) {
	if len(backgroundColor) > 0 && string(backgroundColor[0]) == "#" {
		label.backgroundColor = backgroundColor
		label.needsRepaint = true
	}
}

func (label *LabelWidget) draw() {
	context := label.window.context
	top, left, width, height := label.computedBox.GetCoords()

	context.SetHexColor(label.backgroundColor)
	context.DrawRectangle(left, top, width, height)
	context.Fill()

	context.SetHexColor(label.fontColor)
	context.SetFont(label.font, label.fontSize)
	context.DrawString(label.content, left+label.fontSize/4, top+label.fontSize*2/2)
	context.Fill()

	CopyWidgetToBuffer(label, context.GetImage())
}
