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

	"github.com/dusk125/pixelui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)
import "time"

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

	// Currently, we NEED this shader in order for imgui to render properly.
	win.Canvas().SetFragmentShader(`
	#version 330 core
	in vec4  vColor;
	in vec2  vTexCoords;
	in float vIntensity;

	out vec4 fragColor;

	uniform vec4 uColorMask;
	uniform vec4 uTexBounds;
	uniform sampler2D uTexture;
	void main() {
		fragColor *= vColor * texture(uTexture, vTexCoords).a;
	}
	`)

	context := imgui.CreateContext(nil)
	defer context.Destroy()
	ui := pixelui.NewUI(context, win)

	// imgui.StyleColorsLight()

	open := true
	for !win.Closed() {
		ui.NewFrame()
		if win.JustReleased(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		win.Clear(colornames.Skyblue)

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
