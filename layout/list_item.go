package layout

import renderer_structs "renderer/structs"

func paintListItemElement(context *renderer_structs.Context, node *renderer_structs.NodeDOM) {
	style := node.Style
	renderBox := node.RenderBox

	context.DrawRectangle(renderBox.Left, renderBox.Top, renderBox.Width, renderBox.Height)
	context.SetRGBA(style.BackgroundColor.R, style.BackgroundColor.G, style.BackgroundColor.B, style.BackgroundColor.A)
	context.Fill()

	context.DrawCircle(renderBox.Left-15, renderBox.Top+style.FontSize/2, 3)
	context.SetRGBA(style.Color.R, style.Color.G, style.Color.B, style.Color.A)
	context.SetFont(sansSerif[style.FontWeight], style.FontSize)
	context.DrawStringWrapped(node.TextContent, renderBox.Left, renderBox.Top+1, 0, 0, renderBox.Width, 1.5, renderer_structs.AlignLeft)
	context.Fill()
}

func calculateListItemLayout(context *renderer_structs.Context, node *renderer_structs.NodeDOM, childIndex int) {
	style := node.Style
	renderBox := node.RenderBox

	if style.Width == 0 {
		renderBox.Width = node.Parent.RenderBox.Width - 30
	}

	if style.Height == 0 && len(node.TextContent) > 0 {
		context.SetFont(sansSerif[style.FontWeight], style.FontSize)
		renderBox.Height = context.MeasureStringWrapped(node.TextContent, renderBox.Width, 1.5) + 2 + context.FontHeight()*.5
	}

	if childIndex > 0 {
		prev := node.Parent.Children[childIndex-1]
		renderBox.Top = prev.RenderBox.Top + prev.RenderBox.Height
	} else {
		renderBox.Top = node.Parent.RenderBox.Top
	}

	renderBox.Left = 30
}
