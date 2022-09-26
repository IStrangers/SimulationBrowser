package structs

import "image"

type TreeWidget struct {
	BaseWidget

	fontSize  float64
	fontColor string
	nodes     []*TreeWidgetNode

	openIcon  image.Image
	closeIcon image.Image

	selectCallback func(*TreeWidgetNode)
}

func (widget *TreeWidget) RemoveNodes() {
	widget.nodes = nil
}

func (widget *TreeWidget) AddNode(childNode *TreeWidgetNode) {
	widget.nodes = append(widget.nodes, childNode)
}

func (button *TreeWidget) draw() {

}
