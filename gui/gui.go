package gui

import (
	"fmt"
	"runtime"
	"time"

	"github.com/liamg/glfont"

	"github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"gitlab.com/liamg/raft/buffer"
	"gitlab.com/liamg/raft/config"
	"gitlab.com/liamg/raft/terminal"
	"go.uber.org/zap"
)

type GUI struct {
	window      *glfw.Window
	logger      *zap.SugaredLogger
	config      *config.Config
	terminal    *terminal.Terminal
	width       int //window width in pixels
	height      int //window height in pixels
	font        *glfont.Font
	fontScale   int32
	renderer    Renderer
	colourAttr  uint32
	renderState *RenderState
}

func New(config *config.Config, terminal *terminal.Terminal, logger *zap.SugaredLogger) *GUI {

	//logger.
	return &GUI{
		config:      config,
		logger:      logger,
		width:       600,
		height:      300,
		terminal:    terminal,
		fontScale:   15.0,
		renderState: NewRenderState(),
	}
}

// inspired by https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-1-hello-opengl

// can only be called on OS thread
func (gui *GUI) resize(w *glfw.Window, width int, height int) {

	if width == gui.width && height == gui.height {
		//return
	}

	gui.logger.Debugf("Initiating GUI resize to %dx%d", width, height)

	gui.width = width
	gui.height = height

	gui.logger.Debugf("Updating font resolution...")
	if gui.font != nil {
		gui.font.UpdateResolution((width), (height))
	}

	gui.logger.Debugf("Setting renderer area...")
	gui.renderer.SetArea(0, 0, gui.width, gui.height)

	gui.logger.Debugf("Calculating size in cols/rows...")
	cols, rows := gui.renderer.GetTermSize()

	gui.logger.Debugf("Resizing internal terminal...")
	if err := gui.terminal.SetSize(cols, rows); err != nil {
		gui.logger.Errorf("Failed to resize terminal to %d cols, %d rows: %s", cols, rows, err)
	}

	gui.logger.Debugf("Resetting render state...")
	gui.renderState.Reset()

	gui.logger.Debugf("Setting viewport size...")
	gl.Viewport(0, 0, int32(gui.width), int32(gui.height))

	gui.logger.Debugf("Resize complete!")

}

func (gui *GUI) glfwScrollCallback(w *glfw.Window, xoff float64, yoff float64) {
	//fmt.Printf("Scroll x=%f y=%f\n", xoff, yoff)
	if yoff > 0 {
		gui.terminal.ScrollUp(1)
	} else {
		gui.terminal.ScrollDown(1)
	}
}

func (gui *GUI) getTermSize() (uint, uint) {
	if gui.renderer == nil {
		return 0, 0
	}
	return gui.renderer.GetTermSize()
}

func (gui *GUI) Close() {
	gui.window.SetShouldClose(true)
}

func (gui *GUI) Render() error {

	gui.logger.Debugf("Locking OS thread...")
	runtime.LockOSThread()

	gui.logger.Debugf("Creating window...")
	var err error
	gui.window, err = gui.createWindow(500, 300)
	if err != nil {
		return fmt.Errorf("Failed to create window: %s", err)
	}
	defer glfw.Terminate()

	gui.logger.Debugf("Initialising OpenGL and creating program...")
	program, err := gui.createProgram()
	if err != nil {
		return fmt.Errorf("Failed to initialise OpenGL: %s", err)
	}

	gui.colourAttr = uint32(gl.GetAttribLocation(program, gl.Str("inColour\x00")))
	gl.BindFragDataLocation(program, 0, gl.Str("outColour\x00"))

	gui.logger.Debugf("Loading font...")
	if err := gui.loadFont("./fonts/envypn-15.ttf"); err != nil {
		//if err := gui.loadFont("./fonts/Roboto.ttf"); err != nil {
		return fmt.Errorf("Failed to load font: %s", err)
	}

	changeChan := make(chan bool, 1)
	titleChan := make(chan bool, 1)

	gui.renderer = NewOpenGLRenderer(gui.config, gui.font, gui.fontScale, 0, 0, gui.width, gui.height, gui.colourAttr, program)

	gui.window.SetFramebufferSizeCallback(gui.resize)
	gui.window.SetKeyCallback(gui.key)
	gui.window.SetCharCallback(gui.char)
	gui.window.SetScrollCallback(gui.glfwScrollCallback)
	gui.window.SetMouseButtonCallback(gui.mouseButtonCallback)
	gui.window.SetRefreshCallback(func(w *glfw.Window) {
		select {
		case changeChan <- true:
		default:
		}
	})
	gui.window.SetFocusCallback(func(w *glfw.Window, focused bool) {
		if focused {
			select {
			case changeChan <- true:
			default:
			}
		}
	})
	w, h := gui.window.GetSize()
	gui.resize(gui.window, w, h)

	gui.logger.Debugf("Starting pty read handling...")

	go func() {
		err := gui.terminal.Read()
		if err != nil {
			gui.logger.Errorf("Read from pty failed: %s", err)
		}
		gui.Close()
	}()

	gui.logger.Debugf("Starting render...")

	gl.UseProgram(program)

	// stop smoothing fonts

	gl.Disable(gl.DEPTH_TEST)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.ClearColor(
		gui.config.ColourScheme.Background[0],
		gui.config.ColourScheme.Background[1],
		gui.config.ColourScheme.Background[2],
		1.0,
	)

	gui.terminal.AttachTitleChangeHandler(titleChan)
	gui.terminal.AttachDisplayChangeHandler(changeChan)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	defaultCell := buffer.NewBackgroundCell(gui.config.ColourScheme.Background)

	var lastCursorX uint
	var lastCursorY uint

	for !gui.window.ShouldClose() {

		dirty := false

		select {

		case <-ticker.C:
			gui.logger.Sync()
		case <-changeChan:
			dirty = true
		case <-titleChan:
			gui.window.SetTitle(gui.terminal.GetTitle())
		default:
		}

		gl.UseProgram(program)

		if dirty || gui.terminal.CheckDirty() {

			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)

			if gui.terminal.Modes().ShowCursor {
				cx := uint(gui.terminal.GetLogicalCursorX())
				cy := uint(gui.terminal.GetLogicalCursorY())
				cy = cy + uint(gui.terminal.GetScrollOffset())

				if lastCursorX != cx || lastCursorY != cy {
					gui.renderState.SetDirty(lastCursorX, lastCursorY)
					dirty = true
				}
			}

			lines := gui.terminal.GetVisibleLines()
			lineCount := gui.terminal.ActiveBuffer().ViewHeight()
			colCount := gui.terminal.ActiveBuffer().ViewWidth()
			for y := 0; y < int(lineCount); y++ {

				for x := 0; x < int(colCount); x++ {

					cell := defaultCell

					if y < len(lines) {
						cells := lines[y].Cells()
						if x < len(cells) {
							cell = cells[x]
							if cell.Rune() == 0 {
								cell = defaultCell
							}
						}
					}

					gui.renderer.DrawCell(cell, uint(x), uint(y))

				}
			}

			if gui.terminal.Modes().ShowCursor {
				cx := uint(gui.terminal.GetLogicalCursorX())
				cy := uint(gui.terminal.GetLogicalCursorY())
				cy = cy + uint(gui.terminal.GetScrollOffset())
				gui.renderer.DrawCursor(cx, cy, gui.config.ColourScheme.Cursor)
			}

			gui.window.SwapBuffers()
		}

		//glfw.PollEvents()
		glfw.WaitEventsTimeout(0.02) // up to 50fps on no input, otherwise higher
	}

	gui.logger.Debugf("Stopping render...")
	return nil

}

func (gui *GUI) loadFont(path string) error {
	font, err := glfont.LoadFont(path, gui.fontScale, gui.width, gui.height)
	if err != nil {
		return fmt.Errorf("LoadFont: %v", err)
	}
	gui.font = font
	return nil
}

func (gui *GUI) createWindow(width int, height int) (*glfw.Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMajor, 3) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Terminal", nil, nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()

	return window, nil
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func (gui *GUI) createProgram() (uint32, error) {
	if err := gl.Init(); err != nil {
		return 0, fmt.Errorf("Failed to initialise OpenGL: %s", err)
	}
	gui.logger.Infof("OpenGL version %s", gl.GoStr(gl.GetString(gl.VERSION)))

	gui.logger.Debugf("Compiling shaders...")

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)

	return prog, nil
}
