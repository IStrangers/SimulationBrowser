package structs

import "ui/structs"

type App struct {
	name    string
	windows []*structs.Window
}

func CreateApp(name string) *App {
	return &App{
		name: name,
	}
}

func (app *App) AddWindow(window *structs.Window) {
	app.windows = append(app.windows, window)
}

func (app *App) DestroyWindow(window *structs.Window) {
	var nWindows []*structs.Window

	for _, appWindow := range app.windows {
		if appWindow != window {
			nWindows = append(nWindows, appWindow)
		}
	}

	app.windows = nWindows
	window.destroy()
}
