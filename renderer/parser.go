package renderer

import (
	"gitee.com/QQXQQ/Aix/renderer/structs"
)

func ParserHTML(html string) *structs.NodeDOM {
	options := &structs.HTMLParserOptions{
		RemoveExtraSpaces: true,
	}
	return ParserHTMLByOptions(html, options)
}
func ParserHTMLByOptions(html string, options *structs.HTMLParserOptions) *structs.NodeDOM {
	parser := structs.CreateHTMLParser(html, options)
	return parser.ParserHTML()
}

func ParserCSS(css string) *structs.CSSStyleSheet {
	options := &structs.CSSParserOptions{}
	return ParserCSSByOptions(css, options)
}
func ParserCSSByOptions(css string, options *structs.CSSParserOptions) *structs.CSSStyleSheet {
	parser := structs.CreateCSSParser(css, options)
	return parser.ParserCSS()
}
