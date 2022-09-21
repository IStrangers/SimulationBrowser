package structs

type CSSRule struct {
	Selector    *CSSSelector
	Declaration *CSSDeclaration
	Location    *Selection
}
