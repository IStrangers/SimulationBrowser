package structs

import (
	"net/url"
	ui_structs "ui/structs"
)

type Document struct {
	Title       string
	ContentType string
	URL         *url.URL

	RawDocument string
	DOM         *NodeDOM

	DebugFlag       bool
	DebugWindow     *ui_structs.Window
	DebugTree       *ui_structs.TreeWidget
	SelectedElement *NodeDOM

	OffsetY int
}
