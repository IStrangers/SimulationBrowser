package structs

type MenuEntry struct {
	entryText string
	action    func()

	top  float64
	left float64

	width  float64
	height float64
}
