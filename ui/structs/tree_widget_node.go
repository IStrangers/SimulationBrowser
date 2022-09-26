package structs

type TreeWidgetNode struct {
	Key      string
	Value    string
	Parent   *TreeWidgetNode
	Children []*TreeWidgetNode

	isOpen     bool
	isSelected bool
	index      int
	box        Box
}

func (node *TreeWidgetNode) Toggle() {
	if node.isOpen {
		node.isOpen = false
	} else {
		node.isOpen = true
	}
}

func (node *TreeWidgetNode) Close() {
	node.isOpen = false

}
func (node *TreeWidgetNode) Open() {
	node.isOpen = true
}

func (node *TreeWidgetNode) AddNode(childNode *TreeWidgetNode) {
	childNode.Parent = node
	childNode.index = len(node.Children)
	node.Children = append(node.Children, childNode)
}

func (node *TreeWidgetNode) NextSibling() *TreeWidgetNode {
	selfIdx := node.index
	if selfIdx+1 < len(node.Parent.Children) {
		return node.Parent.Children[selfIdx+1]
	}

	return nil
}

func (node *TreeWidgetNode) PreviousSibling() *TreeWidgetNode {
	selfIdx := node.index
	if selfIdx-1 >= 0 {
		return node.Parent.Children[selfIdx-1]
	}

	return nil
}
