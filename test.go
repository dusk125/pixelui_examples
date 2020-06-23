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

	"github.com/inkyblackness/imgui-go"

	"time"

	"github.com/dusk125/pixelui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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

	ui.AddTTFFont("03b04.ttf", 16)

	// imgui.StyleColorsLight()

	open := true
	for !win.Closed() {
		ui.NewFrame()
		if win.JustReleased(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		win.Clear(colornames.Skyblue)

		if ui.JustPressed(pixelgl.MouseButtonLeft) {
			fmt.Println("Left pressed")
		}

		if ui.JustReleased(pixelgl.MouseButtonLeft) {
			fmt.Println("Left released")
		}

		if open {
			imgui.ShowDemoWindow(&open)
		}
		// imgui.Button("Hello")
		// imgui.Button("Button 2")

		ui.Draw(win)

		win.Update()
		<-time.NewTimer(time.Second / time.Duration(60)).C
	}
}
