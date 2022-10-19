package structs

import (
	"github.com/goki/freetype/raster"
	"github.com/goki/freetype/truetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/f64"
	"image"
	"image/color"
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
	return CreateContextByRGBA(imageToRGBA(im))
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
	r, g, b, a := parseHexColor(backgroundColor)
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

func (context *Context) Width() int {
	return context.width
}

func (context *Context) Height() int {
	return context.height
}

func (context *Context) Clear() {
	src := image.NewUniform(context.color)
	draw.Draw(context.im, context.im.Bounds(), src, image.Point{}, draw.Src)
}

func (context *Context) SetFont(font *truetype.Font, points float64) {
	context.fontFace = truetype.NewFace(font, &truetype.Options{Size: points})
	context.fontHeight = points * 72 / 96
}

func (context *Context) MeasureString(s string) (w, h float64) {
	d := &font.Drawer{
		Face: context.fontFace,
	}
	a := d.MeasureString(s)
	return float64(a >> 6), context.fontHeight
}

func (context *Context) DrawStringAnchored(s string, x, y, ax, ay float64) {
	w, h := context.MeasureString(s)
	x -= ax * w
	y += ay * h
	if context.mask == nil {
		context.drawString(context.im, s, x, y)
	} else {
		im := image.NewRGBA(image.Rect(0, 0, context.width, context.height))
		context.drawString(im, s, x, y)
		draw.DrawMask(context.im, context.im.Bounds(), im, image.Point{}, context.mask, image.Point{}, draw.Over)
	}
}

func (context *Context) drawString(im *image.RGBA, s string, x, y float64) {
	d := &font.Drawer{
		Dst:  im,
		Src:  image.NewUniform(context.color),
		Face: context.fontFace,
		Dot:  fixp(x, y),
	}
	prevC := rune(-1)
	for _, c := range s {
		if prevC >= 0 {
			d.Dot.X += d.Face.Kern(prevC, c)
		}
		dr, mask, maskp, advance, ok := d.Face.Glyph(d.Dot, c)
		if !ok {
			continue
		}
		sr := dr.Sub(dr.Min)

		transformer := draw.BiLinear
		fx, fy := float64(dr.Min.X), float64(dr.Min.Y)
		m := context.matrix.Translate(fx, fy)
		s2d := f64.Aff3{m.XX, m.XY, m.X0, m.YX, m.YY, m.Y0}
		transformer.Transform(d.Dst, s2d, d.Src, sr, draw.Over, &draw.Options{
			SrcMask:  mask,
			SrcMaskP: maskp,
		})
		d.Dot.X += advance
		prevC = c
	}
}

func (context *Context) DrawString(s string, x, y float64) {
	context.DrawStringAnchored(s, x, y, 0, 0)
}

func (context *Context) SetLineWidth(lineWidth float64) {
	context.lineWidth = lineWidth
}

func (context *Context) SetLineJoinRound() {
	context.lineJoin = LineJoinRound
}

func (context *Context) capper() raster.Capper {
	switch context.lineCap {
	case LineCapButt:
		return raster.ButtCapper
	case LineCapRound:
		return raster.RoundCapper
	case LineCapSquare:
		return raster.SquareCapper
	}
	return nil
}

func (context *Context) joiner() raster.Joiner {
	switch context.lineJoin {
	case LineJoinBevel:
		return raster.BevelJoiner
	case LineJoinRound:
		return raster.RoundJoiner
	}
	return nil
}

func (context *Context) stroke(painter raster.Painter) {
	path := context.strokePath
	if len(context.dashes) > 0 {
		path = dashed(path, context.dashes, context.dashOffset)
	} else {
		path = rasterPath(flattenPath(path))
	}
	r := context.rasterizer
	r.UseNonZeroWinding = true
	r.Clear()
	r.AddStroke(path, fix(context.lineWidth), context.capper(), context.joiner())
	r.Rasterize(painter)
}

func (context *Context) StrokePreserve() {
	var painter raster.Painter
	if context.mask == nil {
		if pattern, ok := context.strokePattern.(*SolidPattern); ok {
			p := raster.NewRGBAPainter(context.im)
			p.SetColor(pattern.color)
			painter = p
		}
	}
	if painter == nil {
		painter = CreatePatternPainter(context.im, context.mask, context.strokePattern)
	}
	context.stroke(painter)
}

func (context *Context) Stroke() {
	context.StrokePreserve()
	context.ClearPath()
}

func (context *Context) Push() {
	x := *context
	context.stack = append(context.stack, &x)
}

func (context *Context) Rotate(angle float64) {
	context.matrix = context.matrix.Rotate(angle)
}

func (context *Context) Pop() {
	before := *context
	s := context.stack
	x, s := s[len(s)-1], s[:len(s)-1]
	*context = *x
	context.mask = before.mask
	context.strokePath = before.strokePath
	context.fillPath = before.fillPath
	context.start = before.start
	context.current = before.current
	context.hasCurrent = before.hasCurrent
}

func (context *Context) SetRGBA(r, g, b, a float64) {
	context.color = color.NRGBA{
		R: uint8(r * 255),
		G: uint8(g * 255),
		B: uint8(b * 255),
		A: uint8(a * 255),
	}
	context.SetFillAndStrokeColor(context.color)
}

func (context *Context) SetRGB(r, g, b float64) {
	context.SetRGBA(r, g, b, 1)
}

func (context *Context) WordWrap(s string, w float64) []string {
	return wordWrap(context, s, w)
}

func (context *Context) MeasureStringWrapped(s string, width, lineSpacing float64) float64 {
	lines := context.WordWrap(s, width)

	h := float64(len(lines)) * context.fontHeight * lineSpacing
	h -= (lineSpacing - 1) * context.fontHeight

	return h
}

func (context *Context) FontHeight() float64 {
	return context.fontHeight
}

func (context *Context) DrawStringWrapped(s string, x, y, ax, ay, width, lineSpacing float64, align Align) {
	lines := context.WordWrap(s, width)

	h := float64(len(lines)) * context.fontHeight * lineSpacing
	h -= (lineSpacing - 1) * context.fontHeight

	x -= ax * width
	y -= ay * h
	switch align {
	case AlignLeft:
		ax = 0
	case AlignCenter:
		ax = 0.5
		x += width / 2
	case AlignRight:
		ax = 1
		x += width
	}
	ay = 1
	for _, line := range lines {
		context.DrawStringAnchored(line, x, y, ax, ay)
		y += context.fontHeight * lineSpacing
	}
}
