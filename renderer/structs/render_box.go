package structs

type RenderBox struct {
	Node *NodeDOM

	Top  float64
	Left float64

	Width  float64
	Height float64

	PaddingTop    float64
	PaddingLeft   float64
	PaddingRight  float64
	PaddingBottom float64

	BorderTop    float64
	BorderLeft   float64
	BorderRight  float64
	BorderBottom float64

	MarginTop    float64
	MarginLeft   float64
	MarginRight  float64
	MarginBottom float64
}

func (box *RenderBox) GetRect() (float64, float64, float64, float64) {
	return box.Top, box.Left, box.Width, box.Height
}
