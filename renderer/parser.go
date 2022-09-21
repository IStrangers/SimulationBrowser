package renderer

import (
	"gitee.com/QQXQQ/Aix/renderer/structs"
)

func ParserHTML(html string) *structs.NodeDOM {
	options := &structs.ParserOptions{
		RemoveExtraSpaces: true,
	}
	return ParserHTMLByOptions(html, options)
}

func ParserHTMLByOptions(html string, options *structs.ParserOptions) *structs.NodeDOM {
	parser := structs.CreateHTMLParser(html, options)
	return parser.ParserHTML()
}
