package structs

import ui_structs "SimulationBrowser/ui/structs"

type App struct {
	name    string
	windows []*ui_structs.Window
}

func CreateApp(name string) *App {
	return &App{
		name: name,
	}
}

func (app *App) Run(callback func()) {
	for {
		for _, window := range app.windows {
			if window.IsVisible() && !window.GetGLW().ShouldClose() {
				window.ProcessFrame()
			}
		}
		callback()
	}
}

func setWidgetWindow(widget ui_structs.Widget, window *ui_structs.Window) {
	widget.SetWindow(window)

	for _, childWidget := range widget.Widgets() {
		setWidgetWindow(childWidget, window)
	}
}

func (app *App) AddWindow(window *ui_structs.Window) {
	app.windows = append(app.windows, window)
	setWidgetWindow(window.GetRootFrame(), window)
}

func (app *App) DestroyWindow(window *ui_structs.Window) {
	var nWindows []*ui_structs.Window

	for _, appWindow := range app.windows {
		if appWindow != window {
			nWindows = append(nWindows, appWindow)
		}
	}

	app.windows = nWindows
	window.Destroy()
}
