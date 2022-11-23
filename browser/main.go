package browser

import browser_structs "SimulationBrowser/browser/structs"

func StartWebBrowser() {
	webBrowser := browser_structs.CreateWebBrowser()
	webBrowser.Start()
}
