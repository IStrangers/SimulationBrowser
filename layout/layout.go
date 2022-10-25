package layout

import renderer_structs "renderer/structs"

func LayoutDOM(context *renderer_structs.Context, node *renderer_structs.NodeDOM, childIndex int) {
	nodeChildren := getNodeChildren(node)

	node.RenderBox = &renderer_structs.RenderBox{}
	calculateNode(context, node, childIndex)

	for i := 0; i < len(nodeChildren); i++ {
		LayoutDOM(context, nodeChildren[i], i)
		node.RenderBox.Height += nodeChildren[i].RenderBox.Height
	}

	paintNode(context, node)
}

func getNodeChildren(node *renderer_structs.NodeDOM) []*renderer_structs.NodeDOM {
	var newChildren []*renderer_structs.NodeDOM
	for _, child := range node.Children {
		if child.Style.Display == "none" {
			continue
		}
		newChildren = append(newChildren, child)
	}
	return newChildren
}

func calculateNode(context *renderer_structs.Context, node *renderer_structs.NodeDOM, position int) {
	switch node.Style.Display {
	case "block":
		calculateBlockLayout(context, node, position)
		break
	case "inline":
		calculateInlineLayout(context, node, position)
		break
	case "list-item":
		calculateListItemLayout(context, node, position)
		break
	}
}

func paintNode(context *renderer_structs.Context, node *renderer_structs.NodeDOM) {
	switch node.Style.Display {
	case "block":
		paintBlockElement(context, node)
		break
	case "inline":
		paintInlineElement(context, node)
		break
	case "list-item":
		paintListItemElement(context, node)
		break
	}
}
