package structs

type ContextMenu struct {
	overlay       *Overlay
	entries       []*MenuEntry
	selectedEntry *MenuEntry
}
