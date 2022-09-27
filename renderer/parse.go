package renderer

import (
	renderer_structs "renderer/structs"
)

func ParseHTMLDocument(html string) *renderer_structs.Document {
	document := &renderer_structs.Document{
		RawDocument: html,
	}
	document.DOM = ParseHTML(html, document)
	return document
}

func ParseHTML(html string, document *renderer_structs.Document) *renderer_structs.NodeDOM {
	options := &renderer_structs.HTMLParserOptions{
		RemoveExtraSpaces: true,
	}
	return ParseHTMLByOptions(html, document, options)
}
func ParseHTMLByOptions(html string, document *renderer_structs.Document, options *renderer_structs.HTMLParserOptions) *renderer_structs.NodeDOM {
	parser := renderer_structs.CreateHTMLParser(html, document, options)
	return parser.ParseHTML()
}

func ParseCSS(css string) []*renderer_structs.CSSRule {
	options := &renderer_structs.CSSParserOptions{}
	return ParseCSSByOptions(css, options)
}
func ParseCSSByOptions(css string, options *renderer_structs.CSSParserOptions) []*renderer_structs.CSSRule {
	parser := renderer_structs.CreateCSSParser(css, options)
	return parser.ParseCSS()
}
