package structs

import "image/color"

type Pattern interface {
	ColorAt(x, y int) color.Color
}

type SolidPattern struct {
	color color.Color
}

func (pattern *SolidPattern) ColorAt(x, y int) color.Color {
	return pattern.color
}

func CreateSolidPattern(color color.Color) Pattern {
	return &SolidPattern{
		color: color,
	}
}
