package structs

import (
	"assets"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/freetype/truetype"
	renderer_structs "renderer/structs"
)

type InputWidget struct {
	BaseWidget

	value           string
	selected        bool
	active          bool
	padding         float64
	fontSize        float64
	context         *renderer_structs.Context
	fontColor       string
	cursorFloat     bool
	cursorPosition  int
	cursorDirection bool
	returnCallback  func()
}

func CreateInputWidget() *InputWidget {
	var widgets []Widget
	font, _ := truetype.Parse(assets.OpenSans(400))

	return &InputWidget{
		BaseWidget: BaseWidget{

			needsRepaint: true,
			widgets:      widgets,

			widgetType: inputWidget,

			cursor: glfw.CreateStandardCursor(glfw.IBeamCursor),

			backgroundColor: "#fff",

			font: font,
		},

		fontSize:  20,
		fontColor: "#000",
	}
}

func (input *InputWidget) SetWidth(width float64) {
	input.box.width = width
	input.fixedWidth = true
	input.RequestReflow()
}

func (input *InputWidget) SetHeight(height float64) {
	input.box.height = height
	input.fixedHeight = true
	input.RequestReflow()
}

func (input *InputWidget) SetFontSize(fontSize float64) {
	input.fontSize = fontSize
	input.needsRepaint = true
}

func (input *InputWidget) SetReturnCallback(returnCallback func()) {
	input.returnCallback = returnCallback
}

func (input *InputWidget) SetFontColor(fontColor string) {
	if len(fontColor) > 0 && string(fontColor[0]) == "#" {
		input.fontColor = fontColor
		input.needsRepaint = true
	}
}

func (input *InputWidget) SetValue(value string) {
	input.value = value
	input.needsRepaint = true
}

func (input *InputWidget) GetValue() string {
	return input.value
}

func (input *InputWidget) GetCursorPos() int {
	return input.cursorPosition
}

func (input *InputWidget) SetBackgroundColor(backgroundColor string) {
	if len(backgroundColor) > 0 && string(backgroundColor[0]) == "#" {
		input.backgroundColor = backgroundColor
		input.needsRepaint = true
	}
}

func (input *InputWidget) draw() {
	input.padding = 4
	top, left, width, height := input.computedBox.GetCoords()
	totalPadding := input.padding * 2
	if input.context == nil || input.context.Width() != int(width-totalPadding) || input.context.Height() != int(height-totalPadding) {
		input.context = renderer_structs.CreateContext(int(width-totalPadding), int(height-totalPadding))
	}

	window := input.window
	context := input.context

	if input.selected {
		window.context.SetHexColor("#e4e4e4")
		context.SetHexColor("#e4e4e4")
		context.Clear()
	} else {
		window.context.SetHexColor("#efefef")
		context.SetHexColor("#efefef")
		context.Clear()
	}

	if input.active {
		window.context.SetHexColor("#fff")
		context.SetHexColor("#fff")
		context.Clear()
	} else {
		input.cursorPosition = 0
	}

	window.context.DrawRectangle(left, top, width, height)
	window.context.Fill()

	context.SetHexColor("#2f2f2f")

	context.SetFont(input.font, input.fontSize)
	w, h := context.MeasureString(input.value)

	cursorP := width - totalPadding*2
	cP, _ := context.MeasureString(input.value[len(input.value)+input.cursorPosition:])
	cursorP = cursorP - cP

	if cursorP > 0 {
		input.cursorFloat = true
	} else {
		input.cursorFloat = false
	}

	valueBiggerThanInput := w > float64(width)-input.fontSize
	if valueBiggerThanInput && input.active {
		if cursorP > 0 {
			context.DrawStringAnchored(input.value, float64(width)-input.fontSize, float64(height+totalPadding/2)/2, 1, 0)
		} else {
			context.DrawStringAnchored(input.value, cP, float64(height+totalPadding/2)/2, 1, 0)
		}
	} else {
		context.DrawString(input.value, 0, float64(height+totalPadding/2)/2)
	}

	context.Fill()

	if input.active {
		context.SetHexColor("#000")

		if valueBiggerThanInput {
			if cursorP > 0 {
				context.DrawRectangle(cursorP, h/4, 1.3, float64(input.fontSize))
			} else {
				context.DrawRectangle(0, h/4, 1.3, float64(input.fontSize))
			}

		} else {
			cursorDefaultPosition, _ := context.MeasureString(input.value[:len(input.value)+input.cursorPosition])
			context.DrawRectangle(cursorDefaultPosition, h/4, 1.3, float64(input.fontSize))
		}

		context.Fill()
	}

	window.context.DrawImage(context.GetImage(), int(left+totalPadding/2), int(top+totalPadding/2))
	window.context.SetHexColor("#000")
	window.context.SetLineWidth(.4)

	window.context.DrawRectangle(
		left+1,
		top+1,
		width-2,
		height-2,
	)

	window.context.SetLineJoinRound()
	window.context.Stroke()

	CopyWidgetToBuffer(input, window.context.GetImage())
}
