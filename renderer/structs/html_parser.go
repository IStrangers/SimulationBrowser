package structs

import (
	"common"
	"strings"
)

/*
创建Parser
*/
func CreateHTMLParser(html string) *Parser {
	parser := &Parser{
		OriginalHTML: html,
		HTML:         html,
		Line:         1,
		Column:       1,
		Offset:       0,
	}
	return parser
}

/*
Parser结构
*/
type Parser struct {
	OriginalHTML string
	HTML         string
	Line         int
	Column       int
	Offset       int
}

/*
解析节点
*/
func (parser *Parser) ParserHTML() *NodeDOM {
	rootNodeDOM := parser.createRootNodeDOM()
	rootNodeDOM.Children = parser.ParserChildren(rootNodeDOM)
	return rootNodeDOM
}

/*
	获取当前解析进度
*/
func (parser *Parser) getCursor() *Cursor {
	cursor := &Cursor{
		parser.Line,
		parser.Column,
		parser.Offset,
	}
	return cursor
}

func (parser *Parser) updateCursor(cursor *Cursor) {
	parser.Line = cursor.Line
	parser.Column = cursor.Column
	parser.Offset = cursor.Offset
}

/*
	获取当前解析段信息
*/
func (parser *Parser) getSelection(startCursor *Cursor, endCursor *Cursor) *Selection {
	selection := &Selection{
		StartCursor: startCursor,
		EndCursor:   endCursor,
		HTML:        parser.OriginalHTML[startCursor.Offset:endCursor.Offset],
	}
	return selection
}

/*
 */
func (parser *Parser) advancePositionWithMutation(cursor *Cursor, html string, length int) {
	linesCount := 0
	linePos := -1
	for i := 0; i < length; i++ {
		if html[0] == 10 {
			linesCount++
			linePos = i
		}
	}
	cursor.Line += linesCount
	if linePos == -1 {
		cursor.Column += length
	} else {
		cursor.Column = length - linePos
	}
	cursor.Offset += length
}

/*
 */
func (parser *Parser) advanceBy(length int) {
	cursor := parser.getCursor()
	parser.advancePositionWithMutation(cursor, parser.HTML, length)
	parser.updateCursor(cursor)
	parser.HTML = parser.HTML[length:]
}

/*
 */
func (parser *Parser) advanceBySpaces() {
	result := common.REGEXP_CONSTANT_SPACES.FindString(parser.HTML)
	parser.advanceBy(len(result))
}

/*
	是否是结束符
*/
func (parser *Parser) isEnd() bool {
	return parser.HTML == "" || strings.HasPrefix(parser.HTML, "</")
}

/*
	是否是注释
*/
func (parser *Parser) isComment() bool {
	return strings.HasPrefix(parser.HTML, "<!--")
}

/*
	是否是Element
*/
func (parser *Parser) isElement() bool {
	return strings.HasPrefix(parser.HTML, "<")
}

/*
	解析注释
*/
func (parser *Parser) parseComment(parent *NodeDOM) *NodeDOM {

}

/*
	解析Element
*/
func (parser *Parser) parseElement(parent *NodeDOM) *NodeDOM {

}

/*
	解析文本
*/
func (parser *Parser) parseText(parent *NodeDOM) *NodeDOM {

}

/*
解析子节点
*/
func (parser *Parser) ParserChildren(parent *NodeDOM) []*NodeDOM {
	var children []*NodeDOM
	for parser.isEnd() {
		var node *NodeDOM
		if parser.isComment() {
			node = parser.parseComment(parent)
		} else if parser.isElement() {
			node = parser.parseElement(parent)
		} else {
			node = parser.parseText(parent)
		}
		children = append(children, node)
	}
	return children
}

/*
创建Root节点
*/
func (parser *Parser) createRootNodeDOM() *NodeDOM {
	rootNodeDOM := &NodeDOM{
		NodeName:   "ROOT",
		Attributes: nil,
		Children:   nil,
	}
	return rootNodeDOM
}
