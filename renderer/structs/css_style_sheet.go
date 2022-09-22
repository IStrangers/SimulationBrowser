package structs

type CSSStyleSheet struct {
	StyleMap map[string]string
}

func CreateCSSStyleSheetByCSSString(css string) *CSSStyleSheet {
	items := ParseDeclarationItems(css)
	return CreateCSSStyleSheetByCSSDeclarationItems(items...)
}
func CreateCSSStyleSheetByCSSDeclarationItems(items ...*CSSDeclarationItem) *CSSStyleSheet {
	styleSheet := &CSSStyleSheet{
		StyleMap: make(map[string]string, 0),
	}
	for _, attr := range items {
		styleSheet.StyleMap[attr.Name] = attr.Value
	}
	return styleSheet
}
