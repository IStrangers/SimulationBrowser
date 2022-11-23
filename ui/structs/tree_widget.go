package structs

import (
	"SimulationBrowser/assets"
	renderer_structs "SimulationBrowser/renderer/structs"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/freetype/truetype"
	"image"
)

type TreeWidget struct {
	BaseWidget

	fontSize  float64
	fontColor string
	nodes     []*TreeWidgetNode

	openIcon  image.Image
	closeIcon image.Image

	selectCallback func(*TreeWidgetNode)
}

func CreateTreeWidget() *TreeWidget {
	var widgets []Widget
	font, _ := truetype.Parse(assets.OpenSans(400))

	openIcon, _ := renderer_structs.LoadAsset(nil)
	closeIcon, _ := renderer_structs.LoadAsset(nil)

	return &TreeWidget{
		BaseWidget: BaseWidget{

			needsRepaint: true,
			widgets:      widgets,

			widgetType: treeWidget,

			cursor: glfw.CreateStandardCursor(glfw.ArrowCursor),

			backgroundColor: "#fff",

			font: font,
		},

		openIcon:  openIcon,
		closeIcon: closeIcon,
		fontSize:  20,
		fontColor: "#000",
	}
}

func (tree *TreeWidget) Click() {
	x, y := tree.window.GetCursorPosition()
	node := getIntersectedNode(tree.nodes, x, y)

	if node != nil {
		fireNodeEvents(tree, node, x, y, tree.fontSize)
	}
}

func (tree *TreeWidget) RemoveNodes() {
	tree.nodes = nil
}

func (tree *TreeWidget) AddNode(childNode *TreeWidgetNode) {
	tree.nodes = append(tree.nodes, childNode)
}

func (tree *TreeWidget) SetSelectCallback(selectCallback func(*TreeWidgetNode)) {
	tree.selectCallback = selectCallback
}

func deselectNodes(nodes []*TreeWidgetNode) {
	for _, node := range nodes {
		node.isSelected = false
		deselectNodes(node.Children)
	}
}

func (tree *TreeWidget) SelectNode(node *TreeWidgetNode) {
	for _, treeNode := range tree.nodes {
		treeNode.isSelected = false
		deselectNodes(treeNode.Children)
	}

	node.isSelected = true
}

func (tree *TreeWidget) SelectNodeByValue(value string) {
	if tree != nil {
		selectNodeByValue(tree.nodes, value)
	}
}

func selectNodeByValue(nodes []*TreeWidgetNode, value string) {
	for _, childNode := range nodes {
		selectNodeByValue(childNode.Children, value)

		if childNode.Value == value {
			childNode.isSelected = true
		} else {
			childNode.isSelected = false
		}
	}
}

func selectNode(nodes []*TreeWidgetNode, node *TreeWidgetNode) {
	for _, childNode := range nodes {
		selectNode(childNode.Children, node)

		if childNode == node {
			node.isSelected = true
		} else {
			node.isSelected = false
		}
	}
}

func fireNodeEvents(tree *TreeWidget, node *TreeWidgetNode, x, y, nodeHeight float64) {
	t, l, _, _ := node.box.GetCoords()
	if y > t && y < t+nodeHeight {
		tree.SelectNode(node)
		tree.selectCallback(node)

		if x < l+25 {
			node.Toggle()
		}
	}
}

func getIntersectedNode(nodes []*TreeWidgetNode, x, y float64) *TreeWidgetNode {
	var intersectedNode *TreeWidgetNode
	for _, node := range nodes {
		if x > float64(node.box.left) &&
			x < float64(node.box.left+node.box.width) &&
			y > float64(node.box.top) &&
			y < float64(node.box.top+node.box.height) {

			if !node.isOpen {
				return node
			}

			intersectedNode = node
			childIntersectedNode := getIntersectedNode(node.Children, x, y)
			if childIntersectedNode != nil {
				intersectedNode = childIntersectedNode
			}
		}
	}

	return intersectedNode
}

func (tree *TreeWidget) SetWidth(width float64) {
	tree.box.width = width
	tree.fixedWidth = true
	tree.RequestReflow()
}

func (tree *TreeWidget) SetHeight(height float64) {
	tree.box.height = height
	tree.fixedHeight = true
	tree.RequestReflow()
}

func (tree *TreeWidget) SetFontSize(fontSize float64) {
	tree.fontSize = fontSize
	tree.needsRepaint = true
}

func (tree *TreeWidget) SetFontColor(fontColor string) {
	if len(fontColor) > 0 && string(fontColor[0]) == "#" {
		tree.fontColor = fontColor
		tree.needsRepaint = true
	}
}

func (tree *TreeWidget) SetBackgroundColor(backgroundColor string) {
	if len(backgroundColor) > 0 && string(backgroundColor[0]) == "#" {
		tree.backgroundColor = backgroundColor
		tree.needsRepaint = true
	}
}

func (tree *TreeWidget) draw() {
	context := tree.window.context
	top, left, width, height := tree.computedBox.GetCoords()

	context.SetHexColor(tree.backgroundColor)
	context.DrawRectangle(float64(left), float64(top), float64(width), float64(height))
	context.Fill()

	for _, node := range tree.nodes {
		flowNode(context, node, tree, 0)
		drawNode(context, node, tree, 0)
	}

	CopyWidgetToBuffer(tree, context.GetImage())
}

func flowNode(context *renderer_structs.Context, node *TreeWidgetNode, tree *TreeWidget, level int) {
	node.box.left = tree.fontSize * float64(level)
	node.box.height = tree.fontSize + 4
	node.box.width = tree.computedBox.width

	prevSibling := node.PreviousSibling()
	if node.Parent == nil {
		if prevSibling != nil {
			node.box.top = prevSibling.box.top + prevSibling.box.height
		} else {
			node.box.top = 0
		}
	} else {
		if prevSibling != nil {
			node.box.top = prevSibling.box.top + prevSibling.box.height
		} else {
			node.box.top = node.Parent.box.top + node.Parent.box.height
		}
	}

	if node.isOpen {
		for _, childNode := range node.Children {
			flowNode(context, childNode, tree, level+1)
			node.box.height += childNode.box.height
		}
	}

}

func drawNode(context *renderer_structs.Context, node *TreeWidgetNode, tree *TreeWidget, level int) {
	top, left, width, _ := node.box.GetCoords()

	if node.isSelected {
		context.SetHexColor("#7db1ff32")
		context.DrawRectangle(float64(tree.computedBox.left), float64(top), float64(width), tree.fontSize+4)
		context.Fill()

	} else {
		context.SetHexColor(tree.backgroundColor)
		context.DrawRectangle(float64(tree.computedBox.left), float64(top), float64(width), tree.fontSize+4)
		context.Fill()
	}

	context.SetHexColor(tree.fontColor)
	context.SetFont(tree.font, tree.fontSize)
	context.DrawString(node.Key, float64(left)+20+tree.fontSize/4, float64(top)+tree.fontSize*2/2)
	context.Fill()

	if len(node.Children) > 0 {
		if node.isOpen {
			context.DrawImage(tree.openIcon, int(left+4), int(top+1))

			for _, childNode := range node.Children {
				drawNode(context, childNode, tree, level+1)
			}
		} else {
			context.Push()
			context.Rotate(40)
			context.DrawImage(tree.closeIcon, int(left+4), int(top+1))
			context.Pop()
		}
	}

}
