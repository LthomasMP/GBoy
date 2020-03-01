package gb

import (
	"fmt"
	"image"
	"log"

	"golang.org/x/mobile/event/paint"

	"golang.org/x/mobile/event/size"

	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
)

type Screen struct {
	sizeFactor int
}

func NewScreen() *Screen {
	return &Screen{
		sizeFactor: 4,
	}
}

func (scr *Screen) Run(s screen.Screen) {
	currentWidth := ScreenWidth * scr.sizeFactor
	currentHeight := ScreenHeight * scr.sizeFactor
	w, err := s.NewWindow(&screen.NewWindowOptions{
		Width:  currentWidth,
		Height: currentHeight,
		Title:  fmt.Sprintf(Title, GBoyVersion),
	})
	if err != nil {
		log.Fatalf("error while creating window: %e", err)
	}
	defer w.Release()

	s0 := image.Point{currentWidth, currentHeight}
	b, err := s.NewBuffer(s0)
	if err != nil {
		log.Fatalf("error while creating buffer: %e", err)
	}
	defer b.Release()

	t, err := s.NewTexture(s0)
	if err != nil {
		log.Fatalf("error while creating texture: %e", err)
	}
	defer t.Release()

	for {
		switch e := w.NextEvent().(type) {

		case size.Event:
			if e.WidthPx == 0 && e.HeightPx == 0 {
				return
			}

		case paint.Event:
			w.Publish()

		case lifecycle.Event:
			if e.To == lifecycle.StageDead {
				return
			}
		}
	}
}
