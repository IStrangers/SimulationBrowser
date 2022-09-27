package structs

import (
	"layout"
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

	Viewport *ui_structs.CanvasWidget
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
	rootFrame := ui_structs.CreateFrame(ui_structs.HorizontalFrame)

	viewport := ui_structs.CreateCanvasWidget(GetViewportRenderer(webBrowser))
	scrollBar := ui_structs.CreateScrollBarWidget(ui_structs.VerticalScrollBar)

	viewArea := ui_structs.CreateFrame(ui_structs.VerticalFrame)
	viewArea.AddWidget(viewport)
	viewArea.AddWidget(scrollBar)

	rootFrame.AddWidget(viewArea)
	window.SetRootFrame(rootFrame)
	app.AddWindow(window)
	webBrowser.Window = window

	return webBrowser
}

func GetViewportRenderer(webBrowser *WebBrowser) func(*ui_structs.CanvasWidget) {
	return func(canvas *ui_structs.CanvasWidget) {
		go func() {
			profiler := webBrowser.Profiler
			document := webBrowser.CurrentDocument
			profiler.Start("render")
			ctxBounds := canvas.GetContext().GetImage().Bounds()
			drawingContext := renderer_structs.CreateContext(ctxBounds.Max.X, ctxBounds.Max.Y)

			err := layout.LayoutDocument(drawingContext, document)
			if err != nil {
				println("render", "Can't render page: "+err.Error())
			}

			canvas.SetContext(drawingContext)
			canvas.RequestReflow()
			profiler.Stop("render")

		}()
	}
}

func (webBrowser *WebBrowser) Start() {
	webBrowser.Window.Show()
}
