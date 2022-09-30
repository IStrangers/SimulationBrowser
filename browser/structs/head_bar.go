package structs

import (
	"assets"
	"common"
	ui_structs "ui/structs"
)

type HeadBar struct {
	Frame          *ui_structs.Frame
	StatusLabel    *ui_structs.LabelWidget
	ToolsButton    *ui_structs.ButtonWidget
	NextButton     *ui_structs.ButtonWidget
	PreviousButton *ui_structs.ButtonWidget
	ReloadButton   *ui_structs.ButtonWidget
	UrlInput       *ui_structs.InputWidget
}

func CreateHeadBar() *HeadBar {
	frame := ui_structs.CreateFrame(ui_structs.HorizontalFrame)
	frame.SetHeight(62)

	inputFrame := ui_structs.CreateFrame(ui_structs.VerticalFrame)
	urlInput := ui_structs.CreateInputWidget()
	icon := ui_structs.CreateFrame(ui_structs.VerticalFrame)
	img := ui_structs.CreateImageWidget(assets.Logo())

	previousButton := ui_structs.CreateButtonWidget("", []byte(""))
	previousButton.SetWidth(30)

	nextButton := ui_structs.CreateButtonWidget("", []byte(""))
	nextButton.SetWidth(30)

	reloadButton := ui_structs.CreateButtonWidget("", []byte(""))
	reloadButton.SetWidth(30)

	toolsButton := ui_structs.CreateButtonWidget("", []byte(""))
	toolsButton.SetWidth(34)

	rv := ui_structs.CreateFrame(ui_structs.HorizontalFrame)
	rv.SetBackgroundColor("#ddd")
	rv.SetWidth(5)

	img.SetWidth(50)
	icon.AddWidget(img)
	icon.SetBackgroundColor("#ddd")
	icon.SetWidth(50)

	inputFrame.AddWidget(icon)
	inputFrame.AddWidget(previousButton)
	inputFrame.AddWidget(rv)
	inputFrame.AddWidget(nextButton)
	inputFrame.AddWidget(rv)
	inputFrame.AddWidget(reloadButton)
	inputFrame.AddWidget(rv)
	inputFrame.AddWidget(rv)
	inputFrame.AddWidget(urlInput)
	inputFrame.AddWidget(rv)
	inputFrame.AddWidget(toolsButton)
	inputFrame.AddWidget(rv)

	urlInput.SetFontSize(15)

	dv := ui_structs.CreateFrame(ui_structs.HorizontalFrame)
	dv.SetBackgroundColor("#ddd")
	dv.SetHeight(6)

	pv := ui_structs.CreateFrame(ui_structs.HorizontalFrame)
	pv.SetBackgroundColor("#bfbfbf")
	pv.SetHeight(1)

	statusBar := ui_structs.CreateFrame(ui_structs.HorizontalFrame)
	statusLabel := ui_structs.CreateLabelWidget("The " + common.WebBrowserName + " Web Browser")
	statusLabel.SetBackgroundColor("#ddd")
	statusLabel.SetFontColor("#333")
	statusLabel.SetFontSize(15)
	statusBar.AddWidget(statusLabel)
	statusBar.SetHeight(20)

	frame.AddWidget(dv)
	frame.AddWidget(inputFrame)
	frame.AddWidget(dv)
	frame.AddWidget(pv)
	frame.AddWidget(statusBar)
	frame.AddWidget(pv)

	return &HeadBar{
		Frame:          frame,
		StatusLabel:    statusLabel,
		ToolsButton:    toolsButton,
		NextButton:     nextButton,
		PreviousButton: previousButton,
		ReloadButton:   reloadButton,
		UrlInput:       urlInput,
	}
}
