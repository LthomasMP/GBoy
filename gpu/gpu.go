package gpu

import "github.com/lthomasmp/GBoy/mmu"

type GPU struct {
	Vram []mmu.Byte
	Oam  []mmu.Byte
}

func NewGPU() *GPU {
	return &GPU{
		Vram: []mmu.Byte{0},
		Oam:  []mmu.Byte{0},
	}
}
