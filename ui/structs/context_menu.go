package structs

type ContextMenu struct {
	overlay       *Overlay
	entries       []*MenuEntry
	selectedEntry *MenuEntry
}

func CreateContextMenu() *ContextMenu {
	return &ContextMenu{
		overlay: &Overlay{
			ref: "contextMenu",
		},
	}
}

func (contextMenu *ContextMenu) AddContextMenuEntry(entryText string, action func()) {
	menuEntry := CreateMenuEntry(entryText,action)
	contextMenu.entries = append(contextMenu.entries,menuEntry)
}

func (contextMenu *ContextMenu) DestroyContextMenu() {
	contextMenu.entries = nil
	contextMenu.selectedEntry = nil
}
