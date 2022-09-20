package structs

const (
	NodeType_Root = iota
	NodeType_Element
	NodeType_Text
	NodeType_Common
)

type NodeDOM struct {
	Parent      *NodeDOM
	NodeType    int
	NodeName    string
	Attributes  *[]Attribute
	TextContent string
	Children    []*NodeDOM
	Location    *Selection
}
