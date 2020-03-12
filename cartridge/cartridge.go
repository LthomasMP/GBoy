package cartridge

import "os"

type Cartridge struct {
	Memory []byte
}

func NewCartridge() *Cartridge {
	return &Cartridge{
		// TODO: check if adding capacity is necessary
		Memory: make([]byte, 0x200000),
	}
}

func (cart *Cartridge) LoadROM(file string) (int, error) {
	// TODO: sanitize file path before loading file
	// TODO: check if extension is zip
	rom, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	n, err := rom.Read(cart.Memory)
	if err != nil {
		return 0, err
	} else {
		return n, nil
	}
}
