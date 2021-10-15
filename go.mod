module github.com/dusk125/pixelui_examples

go 1.17

require (
	github.com/dusk125/pixelui v0.1.0
	github.com/dusk125/pixelutils v1.1.0
	github.com/dusk125/rectpack v1.1.1
	github.com/faiface/pixel v0.11.0-beta
	github.com/inkyblackness/imgui-go v1.12.0
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d
)

require (
	github.com/faiface/glhf v0.0.0-20211013000516-57b20770c369 // indirect
	github.com/faiface/mainthread v0.0.0-20171120011319-8b78f0a41ae3 // indirect
	github.com/go-gl/gl v0.0.0-20210905235341-f7a045908259 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20210727001814-0db043d8d5be // indirect
	github.com/go-gl/mathgl v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pzsz/voronoi v0.0.0-20130609164533-4314be88c79f // indirect
)

replace github.com/faiface/pixel => ../../faiface/pixel
replace github.com/dusk125/pixelui => ../pixelui

replace github.com/dusk125/pixelutils => ../pixelutils

replace github.com/dusk125/rectpack => ../rectpack
