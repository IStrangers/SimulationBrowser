package structs

import (
	profiler "profiler/structs"
	renderer "renderer/structs"
	"runtime"
	ui "ui/structs"
)

const (
	WebBrowserName    = "Aix"
	WebBrowserVersion = "1.0.0"
)

type WebBrowser struct {
	CurrentDocument *renderer.Document
	Documents       []*renderer.Document
	Window          *ui.Window
	Profiler        *profiler.Profiler
	Settings        *Settings
}

func CreateWebBrowser() *WebBrowser {
	runtime.LockOSThread()

	defaultSettingsPath := "./settings.json"
	settings := LoadSettings(defaultSettingsPath)

	app := CreateApp(WebBrowserName)
	window := ui.CreateWindow(WebBrowserName, settings.WindowWidth, settings.WindowHeight, settings.HiDPI)
	rootFrame := ui.CreateFrame(ui.HorizontalFrame)
	window.SetRootFrame(rootFrame)
	app.AddWindow(window)

	webBrowser := &WebBrowser{
		Window:   window,
		Profiler: profiler.CreateProfiler(),
		Settings: settings,
	}
	return webBrowser
}

func (webBrowser *WebBrowser) Start() {
	webBrowser.Window.Show()
}
