package structs

type Box struct {
	top    float64
	left   float64
	width  float64
	height float64
}

func (box *Box) SetCoords(top, left, width, height float64) {
	box.top = top
	box.left = left
	box.width = width
	box.height = height
}

func (box *Box) GetCoords() (float64, float64, float64, float64) {
	return box.top, box.left, box.width, box.height
}
