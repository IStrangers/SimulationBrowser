package renderer

import (
	"gitee.com/QQXQQ/Aix/renderer/structs"
)

func ParserHTML(html string) *structs.NodeDOM {
	parser := structs.CreateHTMLParser(html)
	return parser.ParserHTML()
}
