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
	"fmt"
	"log"

	"github.com/dusk125/pixelutils"

	_ "image/png"

	"github.com/dusk125/pixelui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/inkyblackness/imgui-go"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "PixelUi Test",
		Bounds: pixel.R(0, 0, 1920, 1080),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ui := pixelui.NewUI(win, 0)
	defer ui.Destroy()

	// Can also add a loaded sprite with 'ui.AddSprite'
	spriteId, sprite := ui.AddSpriteFromFile("golang.png")
	// spriteId, sprite := ui.AddSpriteFromFile("ship.png")

	ui.AddTTFFont("03b04.ttf", 16)

	// imgui.StyleColorsLight()

	ticker := pixelutils.NewTicker(60)
	for !win.Closed() {
		ui.NewFrame()
		if win.JustReleased(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		win.Clear(colornames.Skyblue)
		_, framerate := ticker.Tick()

		if ui.JustPressed(pixelgl.MouseButtonLeft) {
			fmt.Println("Left pressed")
		}

		if ui.JustReleased(pixelgl.MouseButtonLeft) {
			fmt.Println("Left released")
		}

		imgui.ShowDemoWindow(nil)

		imgui.Begin("Image Test")
		imgui.Text(fmt.Sprintf("%.2f", framerate))
		// Use the pixelui 'Image' helper function
		ui.Image("golang", 0.5)
		// Use the default imgui 'Image' function
		imgui.Image(spriteId, pixelui.IVec(sprite.Picture().Bounds().Size().Scaled(0.25)))
		imgui.End()

		ui.Draw(win)

		win.Update()
		ticker.Wait()
	}
}
