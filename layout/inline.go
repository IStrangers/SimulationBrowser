package layout

import (
	"SimulationBrowser/filesystem"
	renderer_structs "SimulationBrowser/renderer/structs"
	"bytes"
	"fmt"
	"image"
)

func paintInlineElement(context *renderer_structs.Context, node *renderer_structs.NodeDOM) {
	style := node.Style
	renderBox := node.RenderBox

	context.DrawRectangle(renderBox.Left, renderBox.Top, renderBox.Width, renderBox.Height)
	context.SetRGBA(style.BackgroundColor.R, style.BackgroundColor.G, style.BackgroundColor.B, style.BackgroundColor.A)
	context.Fill()

	if node.NodeName == "img" {
		im, err := fetchNodeImage(node)

		if err != nil {
			fmt.Println(err)
			im, _, _ = image.Decode(bytes.NewReader([]byte("")))
		}
		context.DrawImage(im, int(renderBox.Left), int(renderBox.Top))
	}

	context.SetRGBA(style.Color.R, style.Color.G, style.Color.B, style.Color.A)
	context.SetFont(sansSerif[style.FontWeight], style.FontSize)
	context.DrawStringWrapped(node.TextContent, renderBox.Left, renderBox.Top, 0, 0, renderBox.Width, 1, renderer_structs.AlignLeft)
	context.Fill()
}

func fetchNodeImage(node *renderer_structs.NodeDOM) (image.Image, error) {
	imgPath := node.Attr("src")

	imgURL, err := node.Document.URL.Parse(imgPath)
	if err != nil {
		return nil, err
	}

	img, err := filesystem.GetImage(imgURL)
	if err != nil {
		return nil, err
	}
	im, _, err := image.Decode(bytes.NewReader(img.Data))

	if err != nil {
		return nil, err
	}
	return im, nil
}

func calculateInlineLayout(context *renderer_structs.Context, node *renderer_structs.NodeDOM, childIndex int) {
	style := node.Style
	renderBox := node.RenderBox

	context.SetFont(sansSerif[style.FontWeight], style.FontSize)

	if childIndex > 0 && node.Parent.Children[childIndex-1] != nil {
		prev := node.Parent.Children[childIndex-1]
		if prev.Style.Display == "inline" {
			renderBox.Top = prev.RenderBox.Top
			renderBox.Left = prev.RenderBox.Left + prev.RenderBox.Width
		} else {
			renderBox.Top = prev.RenderBox.Top + prev.RenderBox.Height
			renderBox.Left = node.Parent.RenderBox.Left
		}
	} else {
		renderBox.Top = node.Parent.RenderBox.Top
		renderBox.Left = node.Parent.RenderBox.Left
	}

	if node.NodeName == "img" {
		im, err := fetchNodeImage(node)
		if err != nil {
			fmt.Println(err)
			im, _, _ = image.Decode(bytes.NewReader([]byte("")))
		}
		imgSize := im.Bounds().Size()

		renderBox.Width = float64(imgSize.X)
		renderBox.Height = float64(imgSize.Y)
	} else {
		if renderBox.Width == 0 {
			renderBox.Width = node.Parent.RenderBox.Width
		}

		renderBox.Height = context.MeasureStringWrapped(node.TextContent, renderBox.Width, 1)
		mW, _ := context.MeasureString(node.TextContent)
		if mW < renderBox.Width {
			renderBox.Width = mW
		}
	}

	renderBox.Height++
}
