package structs

import (
	"assets"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/goki/freetype/truetype"
	"image"
	"image/draw"
	"log"
	"os"
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
	}

	//开启菜单
	window.EnableContextMenus()
	//重新创建上下文
	window.RecreateContext()
	glw.MakeContextCurrent()

	window.backend = renderer_structs.CreateGLBackend()
	window.addEvents()
	window.generateTexture()
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
	window.context = renderer_structs.CreateContext(window.width, window.height)
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

func (window *Window) SetCursor(cursorType CursorType) {
	switch cursorType {
	case PointerCursor:
		window.glw.SetCursor(window.pointerCursor)
	default:
		window.glw.SetCursor(window.defaultCursor)
	}
}

func (window *Window) AddOverlay(overlay *Overlay) {
	window.overlays = append(window.overlays, overlay)
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

func (window *Window) AddStaticOverlay(overlay *Overlay) {
	window.staticOverlays = append(
		window.staticOverlays,
		overlay,
	)

	window.hasStaticOverlay = true
}

func (window *Window) RemoveStaticOverlay(ref string) {
	for idx, cOverlay := range window.staticOverlays {
		if cOverlay.ref == ref {
			window.staticOverlays = append(window.staticOverlays[:idx], window.staticOverlays[idx+1:]...)
		}
	}

	if len(window.staticOverlays) < 1 {
		window.hasStaticOverlay = false
	}
}

func CreateStaticOverlay(ref string, ctx *renderer_structs.Context, position image.Point) *Overlay {
	buffer := ctx.GetImage().(*image.RGBA)

	return &Overlay{
		ref:    ref,
		active: true,

		top:  float64(position.Y),
		left: float64(position.X),

		width:  float64(buffer.Rect.Max.X),
		height: float64(buffer.Rect.Max.Y),

		position: position,
		buffer:   buffer,
	}
}

func (window *Window) EnableContextMenus() {
	window.contextMenu = CreateContextMenu()
}

func (window *Window) AddContextMenuEntry(entryText string, action func()) {
	window.contextMenu.AddContextMenuEntry(entryText, action)
}

func (window *Window) DestroyContextMenu() {
	window.RemoveOverlay(window.contextMenu.overlay)
	window.contextMenu.DestroyContextMenu()
	window.SetCursor(DefaultCursor)
}

func (window *Window) refreshContextMenu() {
	ctx := renderer_structs.CreateContextByRGBA(window.contextMenu.overlay.buffer)
	menuWidth := float64(window.contextMenu.overlay.buffer.Rect.Max.X)
	textLeft := 4.
	ctx.SetHexColor("#eee")
	ctx.Clear()

	font, _ := truetype.Parse(assets.OpenSans(400))
	ctx.SetHexColor("#222")
	ctx.SetFont(font, 16)

	for idx, entry := range window.contextMenu.entries {
		if window.contextMenu.selectedEntry == entry {
			ctx.DrawRectangle(0, float64(idx*20), menuWidth, 20)
			ctx.SetHexColor("#ccc")
			ctx.Fill()
		}

		ctx.SetHexColor("#222")
		ctx.DrawString(prepEntry(ctx, entry.entryText, menuWidth-textLeft), textLeft, 16+float64(idx*20))
		ctx.Fill()
	}
}

func (window *Window) SelectEntry(entry *MenuEntry) {
	window.contextMenu.selectedEntry = entry
	window.refreshContextMenu()
	window.SetCursor(PointerCursor)
}

func (window *Window) DeselectEntries() {
	if window.contextMenu.selectedEntry != nil {
		window.contextMenu.selectedEntry = nil
		window.refreshContextMenu()
		window.SetCursor(DefaultCursor)
	}
}

func (window *Window) addEvents() {
	window.glw.SetFocusCallback(func(w *glfw.Window, focused bool) {
	})

	window.glw.SetSizeCallback(func(w *glfw.Window, width, height int) {
		xScale, yScale := float32(1), float32(1)
		if window.hiDPI {
			xScale, yScale = w.GetContentScale()
		}

		swidth := int(float32(width) / xScale)
		sheight := int(float32(height) / yScale)

		window.width, window.height = swidth, sheight
		window.RecreateContext()
		//window.RecreateOverlayContext()
		window.needsReflow = true
	})

	window.glw.SetCursorPosCallback(func(w *glfw.Window, x, y float64) {
		xScale, yScale := float32(1), float32(1)
		if window.hiDPI {
			xScale, yScale = w.GetContentScale()
		}

		window.cursorX, window.cursorY = x/float64(xScale), y/float64(yScale)
		window.ProcessPointerPosition()
	})

	window.glw.SetCharCallback(func(w *glfw.Window, char rune) {
		if window.activeInput != nil {
			inputVal, cursorPos := window.activeInput.value, window.activeInput.cursorPosition

			window.activeInput.value = inputVal[:len(inputVal)+cursorPos] + string(char) + inputVal[len(inputVal)+cursorPos:]
			window.activeInput.needsRepaint = true
		}
	})

	window.glw.SetCloseCallback(func(w *glfw.Window) {
		os.Exit(0)
	})

	window.glw.SetKeyCallback(func(w *glfw.Window, key glfw.Key, sc int, action glfw.Action, mods glfw.ModifierKey) {
		switch key {
		case glfw.KeyBackspace:
			if action == glfw.Repeat || action == glfw.Release {
				if window.activeInput != nil && len(window.activeInput.value) > 0 {
					if window.activeInput.cursorPosition == 0 {
						window.activeInput.value = window.activeInput.value[:len(window.activeInput.value)-1]
					} else {
						inputVal, cursorPos := window.activeInput.value, window.activeInput.cursorPosition

						if cursorPos+len(inputVal) > 0 {
							window.activeInput.value = inputVal[:len(inputVal)+cursorPos-1] + inputVal[len(inputVal)+cursorPos:]
						}
					}
					window.activeInput.needsRepaint = true
				}
			}
			break
		case glfw.KeyEscape:
			if action == glfw.Release {
				window.DestroyContextMenu()

				if window.activeInput != nil {
					window.activeInput.active = false
					window.activeInput.selected = false
					window.activeInput.needsRepaint = true
					window.activeInput = nil
				}
			}

			break
		case glfw.KeyUp:
			if action == glfw.Release || action == glfw.Repeat {
				window.ProcessArrowKeys("up")
			}
			break
		case glfw.KeyDown:
			if action == glfw.Release || action == glfw.Repeat {
				window.ProcessArrowKeys("down")
			}
			break
		case glfw.KeyLeft:
			if action == glfw.Release || action == glfw.Repeat {
				window.ProcessArrowKeys("left")
			}
			break
		case glfw.KeyRight:
			if action == glfw.Release || action == glfw.Repeat {
				window.ProcessArrowKeys("right")
			}
			break
		case glfw.KeyV:
			if action == glfw.Release && mods == glfw.ModControl {
				if window.activeInput != nil {
					if window.activeInput.cursorPosition == 0 {
						window.activeInput.value = window.activeInput.value + glfw.GetClipboardString()
					} else {
						inputVal, cursorPos := window.activeInput.value, window.activeInput.cursorPosition
						window.activeInput.value = inputVal[:len(inputVal)+cursorPos] + glfw.GetClipboardString() + inputVal[len(inputVal)+cursorPos:]
					}
					window.activeInput.needsRepaint = true
				}
			}
			break
		case glfw.KeyEnter:
			if action == glfw.Release {
				window.ProcessReturnKey()
			}
			break
		}
	})

	window.glw.SetMouseButtonCallback(func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
		if action == glfw.Release {
			window.ProcessPointerClick(button)
		}
	})

	window.glw.SetScrollCallback(func(w *glfw.Window, x, y float64) {
		window.ProcessScroll(x, y)
	})
}

func (window *Window) generateTexture() {
	gl.DeleteTextures(1, &window.backend.Texture)
	window.frameBuffer = window.context.GetImage().(*image.RGBA)

	if window.rootFrame != nil {
		compositeAll(window.frameBuffer, window.rootFrame)
	}

	if window.hasStaticOverlay {
		nBuffer := image.NewRGBA(window.frameBuffer.Bounds())
		draw.Draw(nBuffer, window.frameBuffer.Bounds(), window.frameBuffer, image.Point{}, draw.Over)
		window.frameBuffer = nBuffer

		for _, overlay := range window.staticOverlays {
			draw.Draw(window.frameBuffer, overlay.buffer.Bounds().Add(overlay.position), overlay.buffer, image.Point{}, draw.Over)
		}
	}

	if window.hasActiveOverlay {
		nBuffer := image.NewRGBA(window.frameBuffer.Bounds())
		draw.Draw(nBuffer, window.frameBuffer.Bounds(), window.frameBuffer, image.Point{}, draw.Over)
		window.frameBuffer = nBuffer

		for _, overlay := range window.overlays {
			draw.Draw(window.frameBuffer, overlay.buffer.Bounds().Add(overlay.position), overlay.buffer, image.Point{}, draw.Over)
		}
	}

	gl.GenTextures(1, &window.backend.Texture)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, window.backend.Texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	gl.TexImage2D(
		gl.TEXTURE_2D, 0, gl.RGBA,
		int32(window.frameBuffer.Rect.Size().X), int32(window.frameBuffer.Rect.Size().Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(window.frameBuffer.Pix),
	)
}

func (window *Window) RegisterInput(input *InputWidget) {
	window.registeredInputs = append(window.registeredInputs, input)
}

func (window *Window) RegisterTree(tree *TreeWidget) {
	window.registeredTrees = append(window.registeredTrees, tree)
}

func (window *Window) RegisterButton(button *ButtonWidget, callback func()) {
	button.onClick = callback
	window.registeredButtons = append(window.registeredButtons, button)
}

func (window *Window) RegisterScrollEventListener(callback func(direction int)) {
	window.scrollEventListeners = append(window.scrollEventListeners, callback)
}

func (window *Window) RegisterClickEventListener(callback func(MouseKey)) {
	window.clickEventListeners = append(window.clickEventListeners, callback)
}

func (window *Window) GetCursorPosition() (float64, float64) {
	return window.cursorX, window.cursorY
}

func (window *Window) SetContextMenuOverlay(overlay *Overlay) {
	window.contextMenu.overlay = overlay
	window.AddOverlay(overlay)
}

func (window *Window) DrawContextMenu() {
	menuWidth := float64(200)
	menuHeight := float64(len(window.contextMenu.entries) * 20)

	menuTop := window.cursorY
	menuLeft := window.cursorX

	if menuLeft+menuWidth > float64(window.width) {
		menuLeft = float64(window.width) - menuWidth
	}

	if menuTop+menuHeight > float64(window.height) {
		menuTop = float64(window.height) - menuHeight
	}

	ctx := renderer_structs.CreateContext(int(menuWidth), int(menuHeight))
	ctx.DrawRectangle(0, 0, menuWidth, menuHeight)
	ctx.SetHexColor("#eee")
	ctx.Fill()

	font, _ := truetype.Parse(assets.OpenSans(400))
	ctx.SetHexColor("#222")
	ctx.SetFont(font, 16)

	textLeft := 4.

	for idx, entry := range window.contextMenu.entries {
		top, left := 16+float64(idx*20), 0.

		entry.setCoords(menuTop+top-16, menuLeft+left, menuWidth, 20)
		ctx.DrawString(prepEntry(ctx, entry.entryText, menuWidth-textLeft), textLeft, top)
		ctx.Fill()
	}

	ctx.DrawRectangle(0, 0, menuWidth, menuHeight)
	ctx.SetHexColor("#ddd")
	ctx.Stroke()

	overlay := extractOverlay(
		ctx.GetImage().(*image.RGBA),
		image.Point{
			X: int(menuLeft),
			Y: int(menuTop),
		})

	window.SetContextMenuOverlay(overlay)
}

func (window *Window) IsVisible() bool {
	return window.visible
}

func (window *Window) GetGLW() *glfw.Window {
	return window.glw
}

func (window *Window) ProcessFrame() {

}
