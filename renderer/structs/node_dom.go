package structs

const (
	NodeType_Root = iota
	NodeType_Element
	NodeType_Text
	NodeType_Common
)

type NodeDOM struct {
	Document *Document
	//父节点
	Parent *NodeDOM
	//节点类型
	NodeType int
	//标签名
	NodeName string
	//是否是自闭合标签
	IsSelfClosing bool
	//文本节点和注释节点内容
	TextContent string
	//属性
	Attributes []*Attribute
	//是否需要回流
	NeedsReflow bool
	//是否需要重绘
	NeedsRepaint bool
	//样式表
	Style *CSSStyleSheet
	//渲染盒子
	RenderBox *RenderBox
	//子节点
	Children []*NodeDOM
	//节点解析信息
	Location *Selection
}
