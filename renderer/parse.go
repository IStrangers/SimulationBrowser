package renderer

import (
	"renderer/structs"
)

func ParseHTMLDocument(html string) *structs.Document {
	document := &structs.Document{
		RawDocument: html,
	}
	document.DOM = ParseHTML(html, document)
	return document
}

func ParseHTML(html string, document *structs.Document) *structs.NodeDOM {
	options := &structs.HTMLParserOptions{
		RemoveExtraSpaces: true,
	}
	return ParseHTMLByOptions(html, document, options)
}
func ParseHTMLByOptions(html string, document *structs.Document, options *structs.HTMLParserOptions) *structs.NodeDOM {
	parser := structs.CreateHTMLParser(html, document, options)
	return parser.ParseHTML()
}

func ParseCSS(css string) []*structs.CSSRule {
	options := &structs.CSSParserOptions{}
	return ParseCSSByOptions(css, options)
}
func ParseCSSByOptions(css string, options *structs.CSSParserOptions) []*structs.CSSRule {
	parser := structs.CreateCSSParser(css, options)
	return parser.ParseCSS()
}
