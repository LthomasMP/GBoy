package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/p4zuu/GBoy/gb/gb"
)

func main() {
	gb := gb.NewGB()

	pixelgl.Run(gb.Screen.Run)
}
