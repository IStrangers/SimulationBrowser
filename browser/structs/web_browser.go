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
	Window   *ui_structs.Window

	App *App
	UI       *WebBrowserUI
	History  *History
	DebuggerMap map[*renderer_structs.Document]*Debugger
	Profiler *profiler_structs.Profiler
	Settings *Settings
}

func CreateWebBrowser() *WebBrowser {
	runtime.LockOSThread()

	defaultSettingsPath := "./settings.json"
	settings := LoadSettings(defaultSettingsPath)

	webBrowser := &WebBrowser{
		App: CreateApp(WebBrowserName),
		History:  CreateHistory(),
		DebuggerMap: make(map[*renderer_structs.Document]*Debugger),
		Settings: settings,
		Profiler: profiler_structs.CreateProfiler(),
	}

	window := ui_structs.CreateWindow(WebBrowserName, settings.WindowWidth, settings.WindowHeight, settings.HiDPI)
	webBrowser.Window = window
	webBrowser.UI = CreateWebBrowserUI(webBrowser)

	app := webBrowser.App
	app.AddWindow(window)
	app.Run(func() {})

	return webBrowser
}

func (webBrowser *WebBrowser) Start() {
	webBrowser.Window.Show()
}
