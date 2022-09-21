package structs

type CSSDeclaration struct {
	Content  string
	Items    []*CSSDeclarationItem
	Location *Selection
}
