package browser

import "browser/structs"

func StartWebBrowser() {
	webBrowser := structs.CreateWebBrowser()
	webBrowser.Start()
}
