package main

import (
	"flag"
	"log"

	"github.com/faiface/pixel/pixelgl"
	"github.com/p4zuu/GBoy/cartridge"
	"github.com/p4zuu/GBoy/gb/gb"
)

var romPath = flag.String("f", "", "ROM file path")

func main() {
	cart := cartridge.NewCartridge()
	_, err := cart.LoadROM(*romPath)
	if err != nil {
		log.Fatalf("error while loading rom file: %s", err.Error())
	}

	gb := gb.NewGB()
	pixelgl.Run(gb.Screen.Run)
}
