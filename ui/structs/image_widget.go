package structs

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"image"
	"log"
	renderer_structs "renderer/structs"
)

type ImageWidget struct {
	BaseWidget

	path string
	img  image.Image
}

func CreateImageWidget(path []byte) *ImageWidget {
	var widgets []Widget

	img, err := renderer_structs.LoadAsset(path)
	if err != nil {
		log.Fatal(err)
	}

	return &ImageWidget{
		BaseWidget: BaseWidget{

			needsRepaint: true,
			widgets:      widgets,

			widgetType: imageWidget,

			cursor: glfw.CreateStandardCursor(glfw.ArrowCursor),

			backgroundColor: "#fff",
		},

		img: img,
	}
}

func (image *ImageWidget) SetWidth(width float64) {
	image.box.width = width
	image.fixedWidth = true
	image.RequestReflow()
}

func (image *ImageWidget) SetHeight(height float64) {
	image.box.height = height
	image.fixedHeight = true
	image.RequestReflow()
}

func (image *ImageWidget) draw() {
	top, left, _, _ := image.computedBox.GetCoords()
	image.window.context.DrawImage(image.img, int(left)+15, int(top)+3)

	CopyWidgetToBuffer(image, image.window.context.GetImage())
}
