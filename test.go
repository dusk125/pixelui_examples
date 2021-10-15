package main

/*
#include "imgui.h"
struct ImDrawVert {
    ImVec2  pos;
    ImVec2  uv;
    ImU32   col;
};
*/

import (
	"log"

	"github.com/dusk125/pixelutils"

	_ "image/png"

	"github.com/dusk125/pixelui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg := pixelgl.WindowConfig{
		Title:  "PixelUi Test",
		Bounds: pixel.R(0, 0, 1920, 1080),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	pixelui.Init(win, pixelui.NO_DEFAULT_FONT)

	// Can also add a loaded sprite with 'ui.AddSprite'
	// spriteId, sprite := ui.AddSpriteFromFile(0, "golang.png")
	// spriteId, sprite := ui.AddSpriteFromFile("ship.png")

	// ui.Packer().InsertFromFile(0, "golang.png", rectpack.AddInsertFlipped)

	// ui.AddTTFFont("03b04.ttf", 16)

	ticker := pixelutils.NewTicker(1000)
	for !win.Closed() {
		pixelui.NewFrame()
		if win.JustReleased(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		win.Clear(colornames.Skyblue)
		_, framerate := ticker.Tick()
		_ = framerate

		// ui.Packer().DrawSub(0, pixel.IM.Moved(win.Bounds().Center()))
		// ui.Packer().Draw(win)

		// imgui.ShowDemoWindow(nil)

		pixelui.Begin("Image Test")
		pixelui.Text("This is a test")
		pixelui.Textf("%.2f", framerate)

		if pixelui.Button("I am a button") {
			log.Println("A clicky button")
		}
		// Use the pixelui 'Image' helper function
		// ui.Image(0, 0.5)
		// Use the default imgui 'Image' function
		// imgui.Image(0, pixelui.IVec(sprite.Picture().Bounds().Size().Scaled(0.25)))
		pixelui.End()

		pixelui.Draw(win, pixel.IM)

		win.Update()
		ticker.Wait()
	}
}
