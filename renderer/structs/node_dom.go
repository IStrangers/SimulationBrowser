package structs

type NodeDOM struct {
	NodeName   string
	Attributes *[]Attribute
	Children   []*NodeDOM
}
