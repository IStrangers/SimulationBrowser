package structs

import (
	"SimulationBrowser/assets"
	renderer_structs "SimulationBrowser/renderer/structs"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/freetype/truetype"
	"image"
)

type ButtonWidget struct {
	BaseWidget
	content string

	icon      image.Image
	fontSize  float64
	fontColor string
	selected  bool
	padding   float64
	onClick   func()
}

func CreateButtonWidget(label string, asset []byte) *ButtonWidget {
	var widgets []Widget
	font, _ := truetype.Parse(assets.OpenSans(400))
	icon, _ := renderer_structs.LoadAsset(asset)

	return &ButtonWidget{
		BaseWidget: BaseWidget{
			needsRepaint: true,
			widgets:      widgets,

			widgetType: buttonWidget,

			cursor: glfw.CreateStandardCursor(glfw.HandCursor),

			backgroundColor: "#fff",

			font: font,
		},

		icon:      icon,
		content:   label,
		fontSize:  20,
		padding:   0,
		fontColor: "#000",
		selected:  false,
	}

}

func (button *ButtonWidget) SetWidth(width float64) {
	button.box.width = width
	button.fixedWidth = true
	button.RequestReflow()
}

func (button *ButtonWidget) SetHeight(height float64) {
	button.box.height = height
	button.fixedHeight = true
	button.RequestReflow()
}

func (button *ButtonWidget) SetFontSize(fontSize float64) {
	button.fontSize = fontSize
	button.needsRepaint = true
}

func (button *ButtonWidget) SetPadding(padding float64) {
	button.padding = padding
	button.needsRepaint = true
}

func (button *ButtonWidget) SetContent(content string) {
	button.content = content
	button.needsRepaint = true
}

func (button *ButtonWidget) Click() {
	button.onClick()
}

func (button *ButtonWidget) GetContent() string {
	return button.content
}

func (button *ButtonWidget) SetFontColor(fontColor string) {
	if len(fontColor) > 0 && string(fontColor[0]) == "#" {
		button.fontColor = fontColor
		button.needsRepaint = true
	}
}

func (button *ButtonWidget) SetBackgroundColor(backgroundColor string) {
	if len(backgroundColor) > 0 && string(backgroundColor[0]) == "#" {
		button.backgroundColor = backgroundColor
		button.needsRepaint = true
	}
}

func (button *ButtonWidget) draw() {
	context := button.window.context
	top, left, width, height := button.computedBox.GetCoords()

	if button.selected {
		context.SetHexColor("#ccc")
	} else {
		context.SetHexColor("#ddd")
	}

	context.DrawRectangle(
		left+button.padding,
		top+button.padding,
		width-(button.padding*2),
		height-(button.padding*2),
	)
	context.Fill()

	if button.content != "" {
		context.SetHexColor(button.fontColor)
		context.SetFont(button.font, button.fontSize)
		context.DrawString(button.content, float64(left)+button.padding, float64(top)+button.padding+button.fontSize)
	}

	if button.icon != nil {
		context.DrawImage(button.icon, int(left+4), int(top+2))
	}

	CopyWidgetToBuffer(button, context.GetImage())
}
