package structs

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"image"
	"log"
	renderer_structs "renderer/structs"
)

type Window struct {
	title string

	width  int
	height int

	hiDPI bool

	needsReflow bool
	visible     bool

	asyncFlag bool

	glw         *glfw.Window
	context     *renderer_structs.Context
	backend     *renderer_structs.GLBackend
	frameBuffer *image.RGBA

	defaultCursor  *glfw.Cursor
	pointerCursor  *glfw.Cursor
	selectedWidget Widget

	registeredTrees   []*TreeWidget
	registeredButtons []*ButtonWidget
	registeredInputs  []*InputWidget
	activeInput       *InputWidget
	rootFrame         *Frame

	cursorX float64
	cursorY float64

	pointerPositionEventListeners []func(float64, float64)
	scrollEventListeners          []func(int)
	clickEventListeners           []func(MouseKey)

	overlays         []*Overlay
	hasActiveOverlay bool

	staticOverlays   []*Overlay
	hasStaticOverlay bool

	contextMenu *ContextMenu

	history *History
}

func CreateWindow(title string, width int, height int, hiDPI bool) *Window {
	glfw.Init()
	setGLFWHints()
	//创建窗口
	glw, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	//限制窗口大小
	glw.SetSizeLimits(300, 200, glfw.DontCare, glfw.DontCare)

	xScale, yScale := float32(1), float32(1)
	if hiDPI {
		xScale, yScale = glw.GetContentScale()
	}

	window := &Window{
		title:  title,
		width:  int(float32(width) / xScale),
		height: int(float32(height) / yScale),
		hiDPI:  hiDPI,
		glw:    glw,

		defaultCursor: glfw.CreateStandardCursor(glfw.ArrowCursor),
		pointerCursor: glfw.CreateStandardCursor(glfw.HandCursor),

		history: CreateHistory(),
	}

	//开启菜单
	window.EnableContextMenus()
	//重新创建上下文
	window.RecreateContext()
	glw.MakeContextCurrent()

	return window
}

func setGLFWHints() {
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
}

/*
显示窗口
*/
func (window *Window) Show() {
	window.needsReflow = true
	window.visible = true
	window.glw.Show()
}

/*
销毁窗口
*/
func (window *Window) Destroy() {
	window.visible = false
	window.glw.Destroy()

	window.glw = nil
	window.context = nil
	window.backend = nil
	window.frameBuffer = nil

	window.defaultCursor = nil
	window.pointerCursor = nil

	window.registeredButtons = nil
	window.registeredInputs = nil
	window.activeInput = nil

	window.rootFrame = nil

	window = nil
}

/*
重新创建上下文
*/
func (window *Window) RecreateContext() {
	window.context = renderer_structs.CreateContext(window.width,window.height)
}

func (window *Window) GetContext() *renderer_structs.Context {
	return window.context
}

func (window *Window) SetNeedsReflow(needsReflow bool) *Window {
	window.needsReflow = needsReflow
	return window
}

func (window *Window) SetRootFrame(rootFrame *Frame) *Window {
	window.rootFrame = rootFrame
	return window
}

func (window *Window) SetCursor(cursorType CursorType)  {
	switch cursorType {
	case PointerCursor:
		window.glw.SetCursor(window.pointerCursor)
	default:
		window.glw.SetCursor(window.defaultCursor)
	}
}

func (window *Window) AddOverlay(overlay *Overlay) {
	window.overlays = append(window.overlays,overlay)
	window.hasActiveOverlay = true
}

func (window *Window) RemoveOverlay(overlay *Overlay) {
	for idx, cOverlay := range window.overlays {
		if cOverlay == overlay {
			window.overlays = append(window.overlays[:idx], window.overlays[idx+1:]...)
			break
		}
	}
	if len(window.overlays) < 1 {
		window.hasActiveOverlay = false
	}
}


func (window *Window) EnableContextMenus() {
	window.contextMenu = CreateContextMenu()
}

func (window *Window) AddContextMenuEntry(entryText string,action func()) {
	window.contextMenu.AddContextMenuEntry(entryText,action)
}

func (window *Window) DestroyContextMenu() {
	window.RemoveOverlay(window.contextMenu.overlay)
	window.contextMenu.DestroyContextMenu()
	window.SetCursor(DefaultCursor)
}