package structs

import (
	"golang.org/x/image/math/fixed"
	"math"
)

type Point struct {
	X, Y float64
}

func (point Point) Fixed() fixed.Point26_6 {
	return Fixp(point.X, point.Y)
}

/*
*
获取两点距离
*/
func (point Point) Distance(p Point) float64 {
	return math.Hypot(point.X-p.X, point.Y-p.Y)
}

func (point Point) Interpolate(p1 Point, p2 Point) Point {
	x := point.X + (p1.X - p2.X)
	y := point.Y + (p1.Y - p2.Y)
	return Point{x, y}
}
