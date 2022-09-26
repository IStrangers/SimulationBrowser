package structs

type WidgetType int

const (
	buttonWidget WidgetType = iota
	canvasWidget
	frameWidget
	imageWidget
	inputWidget
	labelWidget
	treeWidget
	scrollbarWidget
	textWidget
)
