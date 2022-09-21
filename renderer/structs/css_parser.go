package structs

import (
	"gitee.com/QQXQQ/Aix/common"
	"strings"
)

type CSSParserOptions struct {
}

/*
Parser结构
*/
type CSSParser struct {
	Options     *CSSParserOptions
	OriginalCSS string
	CSS         string
	Line        int
	Column      int
	Offset      int
}

func CreateCSSParser(css string, options *CSSParserOptions) *CSSParser {
	css = strings.TrimSpace(css)
	parser := &CSSParser{
		Options:     options,
		OriginalCSS: css,
		CSS:         css,
		Line:        1,
		Column:      1,
		Offset:      0,
	}
	return parser
}

func (parser *CSSParser) ParserCSS() *CSSStyleSheet {
	styleSheet := parser.createStyleSheet()
	styleSheet.CSSRules = parser.parserCSSRules()
	return styleSheet
}

/*
获取当前解析进度
*/
func (parser *CSSParser) getCursor() *Cursor {
	cursor := &Cursor{
		parser.Line,
		parser.Column,
		parser.Offset,
	}
	return cursor
}

func (parser *CSSParser) updateCursor(cursor *Cursor) {
	parser.Line = cursor.Line
	parser.Column = cursor.Column
	parser.Offset = cursor.Offset
}

/*
获取当前解析段信息
*/
func (parser *CSSParser) getSelection(startCursor *Cursor, endCursor *Cursor) *Selection {
	selection := &Selection{
		StartCursor: startCursor,
		EndCursor:   endCursor,
		Content:     parser.OriginalCSS[startCursor.Offset:endCursor.Offset],
	}
	return selection
}

/*
 */
func (parser *CSSParser) advancePositionWithMutation(cursor *Cursor, html string, length int) {
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
func (parser *CSSParser) advanceBy(length int) {
	cursor := parser.getCursor()
	parser.advancePositionWithMutation(cursor, parser.CSS, length)
	parser.updateCursor(cursor)
	parser.CSS = parser.CSS[length:]
}

/*
 */
func (parser *CSSParser) advanceBySpaces() {
	result := common.Regexp_Constant_Spaces.FindString(parser.CSS)
	parser.advanceBy(len(result))
}

func (parser *CSSParser) isEnd() bool {
	return parser.CSS == ""
}

func (parser *CSSParser) parserCSSRules() []*CSSRule {
	var cssRules []*CSSRule
	for !parser.isEnd() {
		cssRule := parser.parserCSSRule()
		cssRules = append(cssRules, cssRule)
	}
	return cssRules
}

func (parser *CSSParser) parserCSSRule() *CSSRule {
	startCursor := parser.getCursor()

	selector := parser.parserSelector()
	declaration := parser.parserDeclaration()

	return &CSSRule{
		Selector:    selector,
		Declaration: declaration,
		Location:    parser.getSelection(startCursor, parser.getCursor()),
	}
}

func (parser *CSSParser) parserSelector() *CSSSelector {
	parser.advanceBySpaces()
	startCursor := parser.getCursor()
	selectorKey := parser.parserSelectorKey()
	selectors := strings.Split(selectorKey, ",")
	selector := &CSSSelector{
		SelectorKey: selectorKey,
		Selectors:   selectors,
		Location:    parser.getSelection(startCursor, parser.getCursor()),
	}
	return selector
}

func (parser *CSSParser) parserSelectorKey() string {
	length := len(parser.CSS)
	endIndex := -1
	for index := 0; index < length; index++ {
		if string(parser.CSS[index]) == "{" {
			endIndex = index
			break
		}
	}
	if endIndex != -1 {
		selector := parser.CSS[0:endIndex]
		parser.advanceBy(len(selector))
		return strings.TrimSpace(selector)
	} else {
		return ""
	}
}

func (parser *CSSParser) parserDeclaration() *CSSDeclaration {
	parser.advanceBy(1)
	parser.advanceBySpaces()
	startCursor := parser.getCursor()

	var content string
	length := len(parser.CSS)
	for index := 0; index < length; index++ {
		if string(parser.CSS[index]) == "}" {
			content = parser.OriginalCSS[startCursor.Offset : startCursor.Offset+index]
			break
		}

	}

	declaration := &CSSDeclaration{
		Content:  content,
		Location: parser.getSelection(startCursor, parser.getCursor()),
	}
	parser.advanceBySpaces()
	parser.advanceBy(1)
	return declaration
}

func (parser *CSSParser) createStyleSheet() *CSSStyleSheet {
	styleSheet := &CSSStyleSheet{}
	return styleSheet
}
