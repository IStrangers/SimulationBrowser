package structs

import (
	profiler_structs "profiler/structs"
	renderer_structs "renderer/structs"
	"runtime"
	"ui"
	ui_structs "ui/structs"
)

const (
	WebBrowserName    = "Aix"
	WebBrowserVersion = "1.0.0"
)

type WebBrowser struct {
	CurrentDocument *renderer_structs.Document
	Documents       []*renderer_structs.Document
	Window          *ui_structs.Window
	Profiler        *profiler_structs.Profiler
	Settings        *Settings
}

func CreateWebBrowser() *WebBrowser {
	runtime.LockOSThread()

	defaultSettingsPath := "./settings.json"
	settings := LoadSettings(defaultSettingsPath)

	app := CreateApp(WebBrowserName)
	window := ui_structs.CreateWindow(WebBrowserName, settings.WindowWidth, settings.WindowHeight, settings.HiDPI)
	rootFrame := ui_structs.CreateFrame(ui_structs.HorizontalFrame)

	viewPort := ui_structs.CreateCanvasWidget(ui.ViewPortRenderer)
	scrollBar := ui_structs.CreateScrollBarWidget(ui_structs.VerticalScrollBar)

	viewArea := ui_structs.CreateFrame(ui_structs.VerticalFrame)
	viewArea.AddWidget(viewPort)
	viewArea.AddWidget(scrollBar)

	rootFrame.AddWidget(viewArea)
	window.SetRootFrame(rootFrame)
	app.AddWindow(window)

	webBrowser := &WebBrowser{
		Window:   window,
		Profiler: profiler_structs.CreateProfiler(),
		Settings: settings,
	}
	return webBrowser
}

func (webBrowser *WebBrowser) Start() {
	webBrowser.Window.Show()
}
