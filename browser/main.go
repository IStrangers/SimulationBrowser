package browser

import (
	"browser/structs"
	"runtime"
)

const (
	WebBrowserName    = "Aix"
	WebBrowserVersion = "1.0.0"
)

func StartWebBrowser() *structs.WebBrowser {
	runtime.LockOSThread()
	webBrowser := &structs.WebBrowser{
		Name:    WebBrowserName,
		Version: WebBrowserVersion,
	}
	return webBrowser
}
