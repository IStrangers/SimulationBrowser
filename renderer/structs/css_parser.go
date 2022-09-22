package structs

import (
	"errors"
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

func (parser *CSSParser) ParseCSS() []*CSSRule {
	cssRules := parser.parseCSSRules()
	return cssRules
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

func (parser *CSSParser) parseCSSRules() []*CSSRule {
	var cssRules []*CSSRule
	for !parser.isEnd() {
		cssRule := parser.parseCSSRule()
		cssRules = append(cssRules, cssRule)
	}
	return cssRules
}

func (parser *CSSParser) parseCSSRule() *CSSRule {
	startCursor := parser.getCursor()

	selector := parser.parseSelector()
	parser.advanceBy(1)
	declaration := parser.parseDeclaration()
	parser.advanceBy(1)

	cssRule := &CSSRule{
		Selector:    selector,
		Declaration: declaration,
		Location:    parser.getSelection(startCursor, parser.getCursor()),
	}
	return cssRule
}

func (parser *CSSParser) parseSelector() *CSSSelector {
	parser.advanceBySpaces()
	startCursor := parser.getCursor()
	selectorKey := parser.parseSelectorKey()
	selectors := strings.Split(selectorKey, ",")
	selector := &CSSSelector{
		SelectorKey: selectorKey,
		Selectors:   selectors,
		Location:    parser.getSelection(startCursor, parser.getCursor()),
	}
	return selector
}

func (parser *CSSParser) parseSelectorKey() string {
	endIndex := strings.Index(parser.CSS, "{")
	if endIndex != -1 {
		selector := parser.CSS[0:endIndex]
		parser.advanceBy(len(selector))
		return strings.TrimSpace(selector)
	} else {
		return ""
	}
}

func (parser *CSSParser) parseDeclaration() *CSSDeclaration {
	startCursor := parser.getCursor()

	endIndex := strings.Index(parser.CSS, "}")
	var content string
	var items []*CSSDeclarationItem
	if endIndex != -1 {
		content = strings.TrimSpace(parser.CSS[0:endIndex])
		parser.advanceBySpaces()
		items = ParseDeclarationItems(content)
	}

	parser.advanceBy(len(content))
	parser.advanceBySpaces()

	declaration := &CSSDeclaration{
		Content:  content,
		Items:    items,
		Location: parser.getSelection(startCursor, parser.getCursor()),
	}
	return declaration
}

func ParseDeclarationItems(cssItems string) []*CSSDeclarationItem {
	var items []*CSSDeclarationItem
	if cssItems != "" {
		for _, cssItem := range strings.Split(cssItems, ";") {
			item, err := ParseDeclarationItem(cssItem)
			if err != nil {
				continue
			}
			items = append(items, item)
		}
	}
	return items
}

func ParseDeclarationItem(css string) (*CSSDeclarationItem, error) {
	kv := strings.Split(strings.TrimSpace(css), ":")
	if len(kv) < 2 {
		return nil, errors.New("The length after splitting is less than 2")
	}
	name := kv[0]
	value := kv[1]
	item := &CSSDeclarationItem{
		Name:  name,
		Value: value,
	}
	return item, nil
}

func (parser *CSSParser) createStyleSheet() *CSSStyleSheet {
	styleSheet := &CSSStyleSheet{}
	return styleSheet
}
