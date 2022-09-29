package structs

import (
	profiler_structs "profiler/structs"
	renderer_structs "renderer/structs"
	"runtime"
	ui_structs "ui/structs"
)

const (
	WebBrowserName    = "Aix"
	WebBrowserVersion = "1.0.0"
)

type WebBrowser struct {
	CurrentDocument *renderer_structs.Document
	Documents       []*renderer_structs.Document

	UI       *WebBrowserUI
	Window   *ui_structs.Window
	Profiler *profiler_structs.Profiler
	Settings *Settings
}

func CreateWebBrowser() *WebBrowser {
	runtime.LockOSThread()

	defaultSettingsPath := "./settings.json"
	settings := LoadSettings(defaultSettingsPath)

	webBrowser := &WebBrowser{
		Settings: settings,
	}

	webBrowser.Profiler = profiler_structs.CreateProfiler()

	app := CreateApp(WebBrowserName)
	window := ui_structs.CreateWindow(WebBrowserName, settings.WindowWidth, settings.WindowHeight, settings.HiDPI)
	app.AddWindow(window)
	webBrowser.Window = window
	webBrowser.UI = CreateWebBrowserUI(webBrowser)

	return webBrowser
}

func (webBrowser *WebBrowser) Start() {
	webBrowser.Window.Show()
}
