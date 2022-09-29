package structs

import (
	"golang.org/x/image/math/fixed"
	"math"
)

type Point struct {
	X, Y float64
}

func (point Point) Fixed() fixed.Point26_6 {
	return fixp(point.X, point.Y)
}

/*
*
获取两点距离
*/
func (point Point) Distance(p Point) float64 {
	return math.Hypot(point.X-p.X, point.Y-p.Y)
}

func (point Point) Interpolate(p1 Point, t float64) Point {
	x := point.X + (p1.X-point.X)*t
	y := point.Y + (p1.Y-point.Y)*t
	return Point{x, y}
}
