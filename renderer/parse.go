package renderer

import (
	"gitee.com/QQXQQ/Aix/renderer/structs"
)

func ParseHTML(html string) *structs.NodeDOM {
	options := &structs.HTMLParserOptions{
		RemoveExtraSpaces: true,
	}
	return ParseHTMLByOptions(html, options)
}
func ParseHTMLByOptions(html string, options *structs.HTMLParserOptions) *structs.NodeDOM {
	parser := structs.CreateHTMLParser(html, options)
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
