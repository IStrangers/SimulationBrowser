package layout

import renderer_structs "renderer/structs"

func LayoutDocument(context *renderer_structs.Context, document *renderer_structs.Document) error {
	body, _ := document.DOM.FindChildByName("body")

	document.DOM.RenderBox.Width = float64(context.Width())
	document.DOM.RenderBox.Height = float64(context.Height())

	context.SetRGB(1, 1, 1)
	context.Clear()

	layoutDOM(context, body, 0)
	return nil
}

func getNodeChildren(node *renderer_structs.NodeDOM) []*renderer_structs.NodeDOM {
	return node.Children
}

func calculateNode(context *renderer_structs.Context, node *renderer_structs.NodeDOM, position int) {
	if node.Style == nil {
		return
	}
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
	if node.Style == nil {
		return
	}
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

func layoutDOM(context *renderer_structs.Context, node *renderer_structs.NodeDOM, childIndex int) {
	nodeChildren := getNodeChildren(node)

	node.RenderBox = &renderer_structs.RenderBox{}
	calculateNode(context, node, childIndex)

	for i := 0; i < len(nodeChildren); i++ {
		layoutDOM(context, nodeChildren[i], i)
		node.RenderBox.Height += nodeChildren[i].RenderBox.Height
	}

	paintNode(context, node)
}
