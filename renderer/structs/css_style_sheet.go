package structs

type CSSStyleSheet struct {
	Color           *ColorRGBA
	BackgroundColor *ColorRGBA

	FontSize   float64
	FontWeight int

	Display  string
	Position string

	Top    float64
	Left   float64
	Width  float64
	Height float64
}

func CreateCSSStyleSheetByInitialStyle(nodeName string, css string) *CSSStyleSheet {
	styleSheet := CreateCSSStyleSheetByCSSString(css)
	if styleSheet.FontSize == float64(0) {
		fontSize := elementFontTable[nodeName]
		if fontSize != float64(0) {
			styleSheet.FontSize = fontSize
		} else {
			styleSheet.FontSize = float64(14)
		}
	}

	if styleSheet.Color == nil {
		color := elementColorTable[nodeName]
		if color != nil {
			styleSheet.Color = color
		} else {
			styleSheet.Color = &ColorRGBA{0, 0, 0, 1}
		}
	}

	if styleSheet.FontWeight == 0 {
		styleSheet.FontWeight = GetDefaultElementFontWeight(nodeName)
	}

	if styleSheet.Display == "" {
		styleSheet.Display = getDefaultElementDisplay(nodeName)
	}
	return styleSheet
}

func CreateCSSStyleSheetByCSSString(css string) *CSSStyleSheet {
	items := ParseDeclarationItems(css)
	return CreateCSSStyleSheetByCSSDeclarationItems(items...)
}
func CreateCSSStyleSheetByCSSDeclarationItems(items ...*CSSDeclarationItem) *CSSStyleSheet {
	styleSheet := &CSSStyleSheet{}
	for _, item := range items {
		styleSheet = mappingCSSDeclarationItemToStyleSheet(styleSheet, item)
	}
	return styleSheet
}

func mappingCSSDeclarationItemToStyleSheet(styleSheet *CSSStyleSheet, item *CSSDeclarationItem) *CSSStyleSheet {
	name := item.Name
	value := item.Value
	switch name {
	case "color":
		styleSheet.Color = ConvertColorToColorRGBA(value)
	case "background-color":
		styleSheet.BackgroundColor = ConvertColorToColorRGBA(value)
	case "display":
		styleSheet.Display = value
	case "position":
		styleSheet.Position = value
	case "width":
		styleSheet.Width = ConvertSizeValue(value)
	case "height":
		styleSheet.Height = ConvertSizeValue(value)
	}
	return styleSheet
}

func getDefaultElementDisplay(nodeName string) string {
	displayType := "block"

	switch nodeName {
	case "script", "style", "meta", "link", "head", "title":
		displayType = "none"
	case "li":
		displayType = "list-item"
	case "html:text", "a", "abbr", "acronym", "b", "bdo", "big", "br",
		"button", "cite", "code", "dfn", "em", "i", "img", "input", "kbd",
		"label", "map", "object", "output", "q", "samp", "select", "small",
		"span", "strong", "sub", "sup", "textarea", "time", "tt", "var", "font":
		displayType = "inline"
	}

	return displayType
}
