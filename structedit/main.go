package main

import (
	"log"

	"github.com/dusk125/pixelui"
	"github.com/dusk125/pixelui/structedit"
	"github.com/dusk125/pixelutils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/inkyblackness/imgui-go"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(Run)
}

type Thing struct {
	String      string
	Float32     float32
	Float64     float64
	Bool, Bool2 bool
	Int         int
	Int8        int8
	Int16       int16
	Int32       int32
	Int64       int64
	UInt        uint
	UInt8       uint8
	UInt16      uint16
	UInt32      uint32
	UInt64      uint64
	Pointer     *Thing
	Struct      struct {
		String string
	}
	Chan chan interface{}
}

func Run() {
	var (
		err error
		cfg = pixelgl.WindowConfig{
			Title:  "StructEdit",
			Bounds: pixel.R(0, 0, 1280, 720),
		}
		ticker = pixelutils.NewTicker(60)
		win    *pixelgl.Window
		ui     *pixelui.UI
		thing  = Thing{
			String:  "Hello there",
			Float64: 1.5,
			Bool:    false,
			Bool2:   true,
		}
	)

	if win, err = pixelgl.NewWindow(cfg); err != nil {
		log.Fatalln(err)
	}

	ui = pixelui.NewUI(win, 0)
	defer ui.Destroy()

	for !win.Closed() {
		ui.NewFrame()
		ticker.Tick()

		win.Clear(colornames.Aliceblue)

		if win.JustReleased(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		structedit.Render("Thing##point", &thing)

		if imgui.Begin("StructView") {
			structedit.Render("Thing##struct", thing)
		}
		imgui.End()

		ui.Draw(win)

		win.Update()
		ticker.Wait()
	}
}
