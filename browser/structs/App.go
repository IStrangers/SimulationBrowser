package structs

type App struct {
	name    string
	windows []*Window
}

func CreateApp(name string) *App {
	return &App{
		name: name,
	}
}

func (app *App) AddWindow(window *Window) {
	app.windows = append(app.windows, window)
}

func (app *App) DestroyWindow(window *Window) {
	var nWindows []*Window

	for _, appWindow := range app.windows {
		if appWindow != window {
			nWindows = append(nWindows, appWindow)
		}
	}

	app.windows = nWindows
	window.destroy()
}
