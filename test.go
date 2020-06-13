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
	"C"
	"log"

	"github.com/inkyblackness/imgui-go"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)
import "github.com/dusk125/pixelui"

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

	context := imgui.CreateContext(nil)
	defer context.Destroy()
	ui := pixelui.NewUi(context)

	imgui.StyleColorsDark()

	for !win.Closed() {
		if win.JustReleased(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		win.Clear(colornames.Skyblue)

		imgui.NewFrame()
		imgui.ShowDemoWindow(nil)
		// imgui.Button("Hello")
		// imgui.Button("Button 2")

		ui.Draw(win)

		win.Update()
	}
}
