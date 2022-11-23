package structs

import renderer_structs "SimulationBrowser/renderer/structs"

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
		action:    action,
	}
}

func (entry *MenuEntry) PointIntersects(x, y float64) bool {
	top, left, width, height := entry.getCoords()
	if x > left &&
		x < left+width &&
		y > top &&
		y < top+height {
		return true
	}

	return false
}

func (entry *MenuEntry) getCoords() (float64, float64, float64, float64) {
	return entry.top, entry.left, entry.width, entry.height
}

func (entry *MenuEntry) setCoords(top, left, width, height float64) {
	entry.top, entry.left = top, left
	entry.width, entry.height = width, height
}

func prepEntry(ctx *renderer_structs.Context, entry string, width float64) string {
	w, _ := ctx.MeasureString(entry)

	if w < width {
		return entry
	}

	for i := 0; i < len(entry); i++ {
		nW, _ := ctx.MeasureString(entry[:len(entry)-i] + "...")

		if nW <= width {
			return entry[:len(entry)-i] + "..."
		}
	}

	return entry
}
