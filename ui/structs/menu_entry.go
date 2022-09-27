package structs

type MenuEntry struct {
	entryText string
	action    func()

	top  float64
	left float64

	width  float64
	height float64
}

func CreateMenuEntry(entryText string, action func()) *MenuEntry {
	return &MenuEntry{
		entryText: entryText,
		action: action,
	}
}