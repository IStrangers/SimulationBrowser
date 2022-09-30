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
