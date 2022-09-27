package structs

import ui_structs "ui/structs"

type App struct {
	name    string
	windows []*ui_structs.Window
}

func CreateApp(name string) *App {
	return &App{
		name: name,
	}
}

func (app *App) AddWindow(window *ui_structs.Window) {
	app.windows = append(app.windows, window)
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
