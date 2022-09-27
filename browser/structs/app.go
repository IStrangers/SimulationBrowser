package structs

import ui "ui/structs"

type App struct {
	name    string
	windows []*ui.Window
}

func CreateApp(name string) *App {
	return &App{
		name: name,
	}
}

func (app *App) AddWindow(window *ui.Window) {
	app.windows = append(app.windows, window)
}

func (app *App) DestroyWindow(window *ui.Window) {
	var nWindows []*ui.Window

	for _, appWindow := range app.windows {
		if appWindow != window {
			nWindows = append(nWindows, appWindow)
		}
	}

	app.windows = nWindows
	window.Destroy()
}
