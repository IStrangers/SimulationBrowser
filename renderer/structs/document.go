package structs

import (
	"net/url"
)

type Document struct {
	Title       string
	ContentType string
	URL         *url.URL

	RawDocument string
	DOM         *NodeDOM

	SelectedElement *NodeDOM

	OffsetY int
}

func (document *Document) GetDocumentTitle() string {
	pageTitle := ""
	if document.DOM != nil {
		titleNode, err := document.DOM.FindChildByName("title")
		if err == nil && len(titleNode.Children) > 0 {
			return titleNode.Children[0].TextContent
		}
	}
	return pageTitle
}
