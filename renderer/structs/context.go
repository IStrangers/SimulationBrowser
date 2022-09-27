package structs

import (
	"github.com/goki/freetype/raster"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"image"
	"image/color"
	"image/draw"
)

var (
	defaultFillStyle   = CreateSolidPattern(color.White)
	defaultStrokeStyle = CreateSolidPattern(color.Black)
)

type Context struct {
	width         int
	height        int
	rasterizer    *raster.Rasterizer
	im            *image.RGBA
	mask          *image.Alpha
	color         color.Color
	fillPattern   Pattern
	strokePattern Pattern
	strokePath    raster.Path
	fillPath      raster.Path
	start         Point
	current       Point
	hasCurrent    bool
	dashes        []float64
	dashOffset    float64
	lineWidth     float64
	lineCap       LineCap
	lineJoin      LineJoin
	fillRule      FillRule
	fontFace      font.Face
	fontHeight    float64
	matrix        Matrix
	stack         []*Context
}

func CreateContext(width, height int) *Context {
	return CreateContextByRGBA(image.NewRGBA(image.Rect(0, 0, width, height)))
}

func CreateContextByImage(im image.Image) *Context {
	return CreateContextByRGBA(ImageToRGBA(im))
}

func CreateContextByRGBA(im *image.RGBA) *Context {
	w := im.Bounds().Size().X
	h := im.Bounds().Size().Y
	return &Context{
		width:         w,
		height:        h,
		rasterizer:    raster.NewRasterizer(w, h),
		im:            im,
		color:         color.Transparent,
		fillPattern:   defaultFillStyle,
		strokePattern: defaultStrokeStyle,
		lineWidth:     1,
		fillRule:      FillRuleWinding,
		fontFace:      basicfont.Face7x13,
		fontHeight:    13,
		matrix:        Identity(),
	}
}

func (context *Context) SetHexColor(backgroundColor string) {
	r, g, b, a := ParseHexColor(backgroundColor)
	context.SetRGBA255(r, g, b, a)
}

func (context *Context) DrawRectangle(x float64, y float64, w float64, h float64) {
	context.CreateSubPath()
	context.MoveTo(x, y)
	context.LineTo(x+w, y)
	context.LineTo(x+w, y+h)
	context.LineTo(x, y+h)
	context.ClosePath()
}

func (context *Context) Fill() {
	context.FillPreserve()
	context.ClearPath()
}

func (context *Context) GetImage() image.Image {
	return context.im
}

func (context *Context) SetRGBA255(r int, g int, b int, a int) {
	context.color = color.NRGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
	context.SetFillAndStrokeColor(context.color)
}

func (context *Context) SetFillAndStrokeColor(color color.Color) {
	context.color = color
	context.fillPattern = CreateSolidPattern(color)
	context.strokePattern = CreateSolidPattern(color)
}

/*
创建绘制路径
*/
func (context *Context) CreateSubPath() {
	if context.hasCurrent {
		context.fillPath.Add1(context.start.Fixed())
	}
	context.hasCurrent = false
}

/*
移动路径
*/
func (context *Context) MoveTo(x float64, y float64) {
	if context.hasCurrent {
		context.fillPath.Add1(context.start.Fixed())
	}
	x, y = context.TransformPoint(x, y)
	point := Point{x, y}
	context.strokePath.Start(point.Fixed())
	context.fillPath.Start(point.Fixed())
	context.start = point
	context.current = point
	context.hasCurrent = false
}

/*
连接路径
*/
func (context *Context) LineTo(x float64, y float64) {
	if context.hasCurrent {
		x, y = context.TransformPoint(x, y)
		point := Point{x, y}
		context.strokePath.Add1(point.Fixed())
		context.fillPath.Add1(point.Fixed())
		context.current = point
	} else {
		context.MoveTo(x, y)
	}
}

/*
结束绘制路径
*/
func (context *Context) ClosePath() {
	if context.hasCurrent {
		context.strokePath.Add1(context.start.Fixed())
		context.fillPath.Add1(context.start.Fixed())
		context.current = context.start
	}
}

func (context *Context) TransformPoint(x float64, y float64) (tx, ty float64) {
	return context.matrix.TransformPoint(x, y)
}

func (context *Context) FillPreserve() {
	var painter raster.Painter
	if context.mask == nil {
		if pattern, ok := context.fillPattern.(*SolidPattern); ok {
			p := raster.NewRGBAPainter(context.im)
			p.SetColor(pattern.color)
			painter = p
		}
	}
	if painter == nil {
		painter = CreatePatternPainter(context.im, context.mask, context.fillPattern)
	}
	context.FillByPainter(painter)
}

func (context *Context) ClearPath() {
	context.strokePath.Clear()
	context.fillPath.Clear()
	context.hasCurrent = false
}

func (context *Context) FillByPainter(painter raster.Painter) {
	path := context.fillPath
	if context.hasCurrent {
		path = make(raster.Path, len(context.fillPath))
		copy(path, context.fillPath)
		path.Add1(context.start.Fixed())
	}
	rasterizer := context.rasterizer
	rasterizer.UseNonZeroWinding = context.fillRule == FillRuleWinding
	rasterizer.Clear()
	rasterizer.AddPath(path)
	rasterizer.Rasterize(painter)
}

func (context *Context) DrawImage(src image.Image, x int, y int) {
	dest, _ := context.GetImage().(draw.Image)
	draw.Draw(dest, src.Bounds().Add(image.Pt(x, y)), src, image.Point{}, draw.Over)
}
