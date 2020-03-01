package gb

import (
	"fmt"
	"log"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Screen struct {
	sizeFactor float64
	win        *pixelgl.Window
	winCfg     pixelgl.WindowConfig
}

func NewScreen() *Screen {
	return &Screen{
		sizeFactor: 4,
	}
}

func (scr *Screen) Reset() {
	scr.win.Clear(colornames.White)
}

func (scr *Screen) Run() {
	currentWidth := ScreenWidth * scr.sizeFactor
	currentHeight := ScreenHeight * scr.sizeFactor
	scr.winCfg = pixelgl.WindowConfig{
		Title:  fmt.Sprintf(Title, GBoyVersion),
		Bounds: pixel.R(0, 0, currentWidth, currentHeight),
		VSync:  true,
	}

	var err error
	scr.win, err = pixelgl.NewWindow(scr.winCfg)
	if err != nil {
		log.Fatalf("error while creating window: %e", err)
	}

	for !scr.win.Closed() {
		scr.win.Update()
	}
}
