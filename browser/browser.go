package browser

import browser "browser/structs"

func StartWebBrowser() {
	webBrowser := browser.CreateWebBrowser()
	webBrowser.Start()
}
