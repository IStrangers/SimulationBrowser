package structs

import (
	renderer "renderer/structs"
	"runtime"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	WebBrowserName    = "Aix"
	WebBrowserVersion = "1.0.0"
)

type WebBrowser struct {
	CurrentDocument *renderer.Document
	Documents []*renderer.Document
	Window *Window
	Settings *Settings
}

func CreateWebBrowser() *WebBrowser {
	beforeInit()

	defaultSettingsPath := "./settings.json"
	settings := LoadSettings(defaultSettingsPath)

	webBrowser := &WebBrowser{
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