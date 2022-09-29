package structs

import (
	renderer_structs "renderer/structs"
	ui_structs "ui/structs"
)

type Debugger struct {
	DebugFlag       bool
	DebugWindow     *ui_structs.Window
	DebugTree       *ui_structs.TreeWidget
	SelectedElement *renderer_structs.NodeDOM
}

func CreateDebugger(debugFlag bool) *Debugger {
	return &Debugger{
		DebugFlag: debugFlag,
	}
}