package structs

import (
	"SimulationBrowser/common"
	"regexp"
	"strings"
)

var elementTagRegexp = regexp.MustCompile(`^<\/?([A-z][^ \t\r\n/>]*)`)
var elementAttrRegexp = regexp.MustCompile(`^[^\t\r\n\f />][^\t\r\n\f />=]*`)

type HTMLParserOptions struct {
	RemoveExtraSpaces bool
}

/*
Parser结构
*/
type HTMLParser struct {
	Options      *HTMLParserOptions
	Document     *Document
	OriginalHTML string
	HTML         string
	Line         int
	Column       int
	Offset       int
}

/*
创建Parser
*/
func CreateHTMLParser(html string, document *Document, options *HTMLParserOptions) *HTMLParser {
	html = strings.TrimSpace(html)
	parser := &HTMLParser{
		Options:      options,
		Document:     document,
		OriginalHTML: html,
		HTML:         html,
		Line:         1,
		Column:       1,
		Offset:       0,
	}
	return parser
}

/*
解析节点
*/
func (parser *HTMLParser) ParseHTML() *NodeDOM {
	rootNodeDOM := parser.createRootNodeDOM()
	rootNodeDOM.Children = parser.parseChildren(rootNodeDOM)
	return rootNodeDOM
}

/*
获取当前解析进度
*/
func (parser *HTMLParser) getCursor() *Cursor {
	cursor := &Cursor{
		parser.Line,
		parser.Column,
		parser.Offset,
	}
	return cursor
}

func (parser *HTMLParser) updateCursor(cursor *Cursor) {
	parser.Line = cursor.Line
	parser.Column = cursor.Column
	parser.Offset = cursor.Offset
}

/*
获取当前解析段信息
*/
func (parser *HTMLParser) getSelection(startCursor *Cursor, endCursor *Cursor) *Selection {
	selection := &Selection{
		StartCursor: startCursor,
		EndCursor:   endCursor,
		Content:     parser.OriginalHTML[startCursor.Offset:endCursor.Offset],
	}
	return selection
}

/*
 */
func (parser *HTMLParser) advancePositionWithMutation(cursor *Cursor, html string, length int) {
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
func (parser *HTMLParser) advanceBy(length int) {
	cursor := parser.getCursor()
	parser.advancePositionWithMutation(cursor, parser.HTML, length)
	parser.updateCursor(cursor)
	parser.HTML = parser.HTML[length:]
}

/*
 */
func (parser *HTMLParser) advanceBySpaces() {
	result := common.Regexp_Constant_Spaces.FindString(parser.HTML)
	parser.advanceBy(len(result))
}

/*
是否是结束符
*/
func (parser *HTMLParser) isEnd() bool {
	return parser.HTML == "" || strings.HasPrefix(parser.HTML, "</")
}

/*
是否是注释
*/
func (parser *HTMLParser) isComment() bool {
	return strings.HasPrefix(parser.HTML, "<!--")
}

/*
是否是Element
*/
func (parser *HTMLParser) isElement() bool {
	return strings.HasPrefix(parser.HTML, "<")
}

/*
解析注释
*/
func (parser *HTMLParser) parseComment(parent *NodeDOM) *NodeDOM {
	startCursor := parser.getCursor()
	commentStart := "<!--"
	commentEnd := "-->"

	parser.advanceBy(len(commentStart))
	parser.advanceBySpaces()

	innerStartCursor := parser.getCursor()
	innerEndCursor := parser.getCursor()

	endIndex := strings.Index(parser.HTML, commentEnd)
	preContent := parser.parseTextData(endIndex)
	content := strings.TrimSpace(preContent)
	startOffset := strings.Index(preContent, content)

	if startOffset > 0 {
		parser.advancePositionWithMutation(innerStartCursor, preContent, startOffset)
	}
	endOffset := startOffset + len(content)
	parser.advancePositionWithMutation(innerEndCursor, preContent, endOffset)
	parser.advanceBy(len(commentEnd))

	return &NodeDOM{
		Document:    parser.Document,
		Parent:      parent,
		NodeType:    NodeType_Common,
		NodeName:    "html:comment",
		TextContent: strings.TrimSpace(content),
		Style:       CreateCSSStyleSheetByInitialStyle("html:comment", ""),
		RenderBox:   &RenderBox{},
		Location:    parser.getSelection(startCursor, parser.getCursor()),
	}
}

/*
解析Element
*/
func (parser *HTMLParser) parseElement(parent *NodeDOM) *NodeDOM {
	nodeDOM := parser.parseElementTag(parent)
	nodeDOM.Children = parser.parseChildren(nodeDOM)
	if strings.HasPrefix(parser.HTML, "</") {
		parser.parseElementTag(parent)
	}
	nodeDOM.Location = parser.getSelection(nodeDOM.Location.StartCursor, parser.getCursor())
	return nodeDOM
}

/*
解析标签
*/
func (parser *HTMLParser) parseElementTag(parent *NodeDOM) *NodeDOM {
	startCursor := parser.getCursor()
	tagName := elementTagRegexp.FindString(parser.HTML)
	if tagName == "" {
	}
	parser.advanceBy(len(tagName))
	parser.advanceBySpaces()
	tagName = strings.ToLower(tagName[1:])
	attributes := parser.parseElementAttributes()
	isSelfClosing := strings.HasPrefix(parser.HTML, "/>")
	if isSelfClosing {
		parser.advanceBy(2)
	} else {
		parser.advanceBy(1)
	}
	initialStyle := GetInitialStyleByAttributes(attributes...)
	return &NodeDOM{
		Document:      parser.Document,
		Parent:        parent,
		NodeType:      NodeType_Element,
		NodeName:      tagName,
		IsSelfClosing: isSelfClosing,
		Attributes:    attributes,
		NeedsReflow:   true,
		NeedsRepaint:  true,
		Style:         CreateCSSStyleSheetByInitialStyle(tagName, initialStyle),
		RenderBox:     &RenderBox{},
		Location:      parser.getSelection(startCursor, parser.getCursor()),
	}
}

func GetInitialStyleByAttributes(attributes ...*Attribute) string {
	initialStyle := ""
	var newAttributes []any
	for _, attr := range attributes {
		newAttributes = append(newAttributes, attr)
	}
	attribute := common.Just(newAttributes...).Filter(func(attr any) bool {
		return attr.(*Attribute).Name == "style"
	}).First()
	if attribute != nil {
		initialStyle = attribute.(*Attribute).Value
	}
	return initialStyle
}

/*
解析属性
*/
func (parser *HTMLParser) parseElementAttributes() []*Attribute {
	var attributes []*Attribute
	for !strings.HasPrefix(parser.HTML, ">") && !parser.isEnd() {
		attr := parser.parseElementAttribute()
		attributes = append(attributes, attr)
		parser.advanceBySpaces()
	}
	return attributes
}

/*
解析属性
*/
func (parser *HTMLParser) parseElementAttribute() *Attribute {
	startCursor := parser.getCursor()
	attrName := elementAttrRegexp.FindString(parser.HTML)
	if attrName == "" {
	}
	var value string
	parser.advanceBy(len(attrName))
	parser.advanceBySpaces()
	if strings.HasPrefix(parser.HTML, "=") {
		parser.advanceBy(1)
		parser.advanceBySpaces()
		value = parser.parseElementAttributeValue()
	}
	return &Attribute{
		Name:     attrName,
		Value:    value,
		Location: parser.getSelection(startCursor, parser.getCursor()),
	}
}

/*
解析属性值
*/
func (parser *HTMLParser) parseElementAttributeValue() string {
	startCursor := parser.getCursor()
	quote := string(parser.HTML[0])

	var value string
	if quote == `"` || quote == `'` {
		parser.advanceBy(1)
		endQuoteIndex := strings.Index(parser.HTML, quote)
		value = parser.parseTextData(endQuoteIndex)
		parser.advanceBy(1)
	} else {
		endSpacesIndex := strings.Index(parser.HTML, " ")
		endCloseIndex := strings.Index(parser.HTML, ">")
		if endCloseIndex == -1 {
			endCloseIndex = strings.Index(parser.HTML, "/>")
		}
		endIndex := endSpacesIndex
		if endSpacesIndex == -1 || endSpacesIndex > endCloseIndex {
			endIndex = endCloseIndex
		}
		value = parser.parseTextData(endIndex)
	}

	parser.getSelection(startCursor, parser.getCursor())
	return value
}

/*
解析文本
*/
func (parser *HTMLParser) parseText(parent *NodeDOM) *NodeDOM {
	tokens := []string{"<"}
	endTextIndex := len(parser.HTML)
	for i := 0; i < len(tokens); i++ {
		index := strings.Index(parser.HTML, tokens[i])
		if index != -1 && index < endTextIndex {
			endTextIndex = index
		}
	}
	startCursor := parser.getCursor()
	content := parser.parseTextData(endTextIndex)
	return &NodeDOM{
		Document:    parser.Document,
		Parent:      parent,
		NodeType:    NodeType_Text,
		NodeName:    "html:text",
		TextContent: strings.TrimSpace(content),
		Style:       CreateCSSStyleSheetByInitialStyle("html:text", ""),
		RenderBox:   &RenderBox{},
		Location:    parser.getSelection(startCursor, parser.getCursor()),
	}
}

/*
解析文本
*/
func (parser *HTMLParser) parseTextData(endTextIndex int) string {
	data := parser.HTML[0:endTextIndex]
	parser.advanceBy(endTextIndex)
	return data
}

/*
解析子节点
*/
func (parser *HTMLParser) parseChildren(parent *NodeDOM) []*NodeDOM {
	var children []*NodeDOM
	for !parser.isEnd() {
		var node *NodeDOM
		if parser.isComment() {
			node = parser.parseComment(parent)
		} else if parser.isElement() {
			node = parser.parseElement(parent)
		} else {
			node = parser.parseText(parent)
			if parser.Options.RemoveExtraSpaces && strings.TrimSpace(node.TextContent) == "" {
				continue
			}
		}
		children = append(children, node)
	}
	return children
}

/*
创建Root节点
*/
func (parser *HTMLParser) createRootNodeDOM() *NodeDOM {
	rootNodeDOM := &NodeDOM{
		Document:   parser.Document,
		NodeType:   NodeType_Root,
		NodeName:   "ROOT",
		Attributes: nil,
		Style:      CreateCSSStyleSheetByInitialStyle("ROOT", ""),
		RenderBox:  &RenderBox{},
		Children:   nil,
	}
	return rootNodeDOM
}
