package structs

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
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
	beforeInit()

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

func beforeInit() {
	runtime.LockOSThread()
	glfw.Init()
	gl.Init()
	setGLFWHints()
}

func setGLFWHints() {
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

func (webBrowser *WebBrowser) Start() {
	webBrowser.Window.Show()
}
