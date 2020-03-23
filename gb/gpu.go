package gb

import "time"

type GPU struct {
	Vram      []byte
	Oam       []byte
	Mode      int
	ModeClock time.Duration
	Line      int
}

func NewGPU() *GPU {
	// TODO: add correct sizes
	return &GPU{
		Vram:      []byte{0},
		Oam:       []byte{0},
		Mode:      0,
		ModeClock: 0,
		Line:      0,
	}
}

func (gb *GB) Step() {
	gb.Gpu.ModeClock += gb.Cpu.Reg.InstructionClock.TotalDuration
	switch gb.Gpu.Mode {

	case 2:
		if gb.Gpu.ModeClock >= 80*time.Millisecond {
			gb.Gpu.ModeClock = 0
			gb.Gpu.Mode = 3
		}
		break

	case 3:
		if gb.Gpu.ModeClock >= 172*time.Millisecond {
			gb.Gpu.ModeClock = 0
			gb.Gpu.Mode = 0

			gb.Renderscan()
		}

	case 0:
		if gb.Gpu.ModeClock >= 204 {
			gb.Gpu.ModeClock = 0
			gb.Gpu.Line++

			if gb.Gpu.Line == 143 {
				gb.Gpu.Mode = 1
				gb.Screen.Reset()
			} else {
				gb.Gpu.Mode = 2
			}
		}
		break

	case 1:
		if gb.Gpu.ModeClock >= 456 {
			gb.Gpu.ModeClock = 0
			gb.Gpu.Line++
			if gb.Gpu.Line > 153 {
				gb.Gpu.Mode = 2
				gb.Gpu.Line = 0
			}
		}
	}
}

func (gb *GB) Renderscan() {
	//Do something
}
