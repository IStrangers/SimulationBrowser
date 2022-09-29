package browser

import browser_structs "browser/structs"

func StartWebBrowser() {
	webBrowser := browser_structs.CreateWebBrowser()
	webBrowser.Start()
}
