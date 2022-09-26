package structs

import (
	"regexp"
	"strconv"
	"strings"
)

var rgbaRegexp = regexp.MustCompile(`rgba?\([\.?\d?\.?\d?%?\s?,?]+\)`)
var rgbaParamsRegexp = regexp.MustCompile(`\([\.?\d?\.?\d?%?\s?,?]+\)`)

var colorTable = map[string]*ColorRGBA{
	"maroon":         {R: 0.5, G: 0.0, B: 0.0, A: 1.0},
	"red":            {R: 1.0, G: 0.0, B: 0.0, A: 1.0},
	"orange":         {R: 1.0, G: 0.6, B: 0.0, A: 1.0},
	"yellow":         {R: 1.0, G: 1.0, B: 0.0, A: 1.0},
	"olive":          {R: 0.5, G: 0.5, B: 0.0, A: 1.0},
	"green":          {R: 0.0, G: 0.5, B: 0.0, A: 1.0},
	"purple":         {R: 0.5, G: 0.0, B: 0.5, A: 1.0},
	"fuchsia":        {R: 1.0, G: 0.0, B: 1.0, A: 1.0},
	"lime":           {R: 0.0, G: 1.0, B: 0.0, A: 1.0},
	"teal":           {R: 0.0, G: 0.5, B: 0.5, A: 1.0},
	"aqua":           {R: 0.0, G: 1.0, B: 1.0, A: 1.0},
	"blue":           {R: 0.0, G: 0.0, B: 1.0, A: 1.0},
	"navy":           {R: 0.0, G: 0.0, B: 0.5, A: 1.0},
	"black":          {R: 0.0, G: 0.0, B: 0.0, A: 1.0},
	"gray":           {R: 0.5, G: 0.5, B: 0.5, A: 1.0},
	"silver":         {R: 0.7, G: 0.7, B: 0.7, A: 1.0},
	"white":          {R: 1.0, G: 1.0, B: 1.0, A: 1.0},
	"tomato":         {R: 1.0, G: 0.38, B: 0.27, A: 1.0},
	"crimson":        {R: 0.8, G: 0.07, B: 0.2, A: 1.0},
	"coral":          {R: 1.0, G: 0.5, B: 0.31, A: 1.0},
	"cornflowerblue": {R: 0.40, G: 0.58, B: 0.92, A: 1.0},
	"darkgreen":      {R: 0.0, G: 0.40, B: 0.0, A: 1.0},
}

var elementColorTable = map[string]*ColorRGBA{
	"a": {R: 0.0, G: 0.0, B: 1.0, A: 1.0},
}

type ColorRGBA struct {
	R float64
	G float64
	B float64
	A float64
}

func (colorRGBA *ColorRGBA) GetRGBA() (float64, float64, float64, float64) {
	return colorRGBA.R, colorRGBA.G, colorRGBA.B, colorRGBA.A
}

func ConvertColorToColorRGBA(color string) *ColorRGBA {
	if string(color[0]) == "#" {
		return HexStringToColorRGBA(color)
	} else if rgbaRegexp.MatchString(color) {
		return RGBAStringToColorRGBA(color)
	}
	return colorTable[color]
}

func hexToFloatInRange(hex string) float64 {
	number, err := strconv.ParseInt(hex, 16, 0)
	if err != nil {
		panic(err)
	}
	return float64(number / 255)
}

func HexStringToColorRGBA(color string) *ColorRGBA {
	color = strings.ToLower(color)
	colorLength := len(color)

	switch colorLength {
	case 9:
		return &ColorRGBA{
			R: hexToFloatInRange(color[1:3]),
			G: hexToFloatInRange(color[3:5]),
			B: hexToFloatInRange(color[5:7]),
			A: hexToFloatInRange(color[7:9]),
		}
	case 7:
		return &ColorRGBA{
			R: hexToFloatInRange(color[1:3]),
			G: hexToFloatInRange(color[3:5]),
			B: hexToFloatInRange(color[5:7]),
			A: 1,
		}
	case 5:
		return &ColorRGBA{
			R: hexToFloatInRange(color[1:2] + color[1:2]),
			G: hexToFloatInRange(color[2:3] + color[2:3]),
			B: hexToFloatInRange(color[3:4] + color[3:4]),
			A: hexToFloatInRange(color[4:5] + color[4:5]),
		}
	case 4:
		return &ColorRGBA{
			R: hexToFloatInRange(color[1:2] + color[1:2]),
			G: hexToFloatInRange(color[2:3] + color[2:3]),
			B: hexToFloatInRange(color[3:4] + color[3:4]),
			A: 1,
		}
	default:
		return &ColorRGBA{}
	}
}

func RGBAStringToColorRGBA(color string) *ColorRGBA {
	paramString := rgbaParamsRegexp.FindString(color)
	if paramString == "" {
		return nil
	}

	paramString = strings.Trim(paramString, "()")
	params := strings.Split(paramString, ",")
	paramsLength := len(params)
	if paramsLength < 3 {
		return nil
	}

	var red, green, blue, alpha float64
	if strings.HasSuffix(params[0], "%") {
		value, _ := strconv.ParseInt(strings.Trim(strings.TrimSpace(params[0]), "%"), 10, 0)
		red = float64(value / 100)
	} else if strings.Index(params[0], ".") != -1 {
		value, _ := strconv.ParseFloat(strings.TrimSpace(params[0]), 64)
		red = value
	} else {
		value, _ := strconv.Atoi(strings.TrimSpace(params[0]))
		red = float64(value / 255)
	}

	if strings.HasSuffix(params[1], "%") {
		value, _ := strconv.ParseInt(strings.Trim(strings.TrimSpace(params[1]), "%"), 10, 0)
		green = float64(value / 100)
	} else if strings.Index(params[1], ".") != -1 {
		value, _ := strconv.ParseFloat(strings.TrimSpace(params[1]), 64)
		green = value
	} else {
		value, _ := strconv.Atoi(strings.TrimSpace(params[1]))
		green = float64(value / 255)
	}

	if strings.HasSuffix(params[2], "%") {
		value, _ := strconv.ParseInt(strings.Trim(strings.TrimSpace(params[2]), "%"), 10, 0)
		blue = float64(value / 100)
	} else if strings.Index(params[2], ".") != -1 {
		value, _ := strconv.ParseFloat(strings.TrimSpace(params[2]), 64)
		blue = value
	} else {
		value, _ := strconv.Atoi(strings.TrimSpace(params[2]))
		blue = float64(value / 255)
	}

	alpha = 1
	return &ColorRGBA{
		R: red,
		G: green,
		B: blue,
		A: alpha,
	}
}
