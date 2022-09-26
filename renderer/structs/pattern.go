package structs

import (
	"github.com/goki/freetype/raster"
	"image"
	"image/color"
)

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

type PatternPainter struct {
	im   *image.RGBA
	mask *image.Alpha
	p    Pattern
}

func (r *PatternPainter) Paint(ss []raster.Span, done bool) {
	b := r.im.Bounds()
	for _, s := range ss {
		if s.Y < b.Min.Y {
			continue
		}
		if s.Y >= b.Max.Y {
			return
		}
		if s.X0 < b.Min.X {
			s.X0 = b.Min.X
		}
		if s.X1 > b.Max.X {
			s.X1 = b.Max.X
		}
		if s.X0 >= s.X1 {
			continue
		}
		const m = 1<<16 - 1
		y := s.Y - r.im.Rect.Min.Y
		x0 := s.X0 - r.im.Rect.Min.X

		i0 := (s.Y-r.im.Rect.Min.Y)*r.im.Stride + (s.X0-r.im.Rect.Min.X)*4
		i1 := i0 + (s.X1-s.X0)*4
		for i, x := i0, x0; i < i1; i, x = i+4, x+1 {
			ma := s.Alpha
			if r.mask != nil {
				ma = ma * uint32(r.mask.AlphaAt(x, y).A) / 255
				if ma == 0 {
					continue
				}
			}
			c := r.p.ColorAt(x, y)
			cr, cg, cb, ca := c.RGBA()
			dr := uint32(r.im.Pix[i+0])
			dg := uint32(r.im.Pix[i+1])
			db := uint32(r.im.Pix[i+2])
			da := uint32(r.im.Pix[i+3])
			a := (m - (ca * ma / m)) * 0x101
			r.im.Pix[i+0] = uint8((dr*a + cr*ma) / m >> 8)
			r.im.Pix[i+1] = uint8((dg*a + cg*ma) / m >> 8)
			r.im.Pix[i+2] = uint8((db*a + cb*ma) / m >> 8)
			r.im.Pix[i+3] = uint8((da*a + ca*ma) / m >> 8)
		}
	}
}

func CreatePatternPainter(im *image.RGBA, mask *image.Alpha, p Pattern) *PatternPainter {
	return &PatternPainter{im, mask, p}
}
