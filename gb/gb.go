package gb

import (
	"github.com/lthomasmp/GBoy/cpu"
	"github.com/lthomasmp/GBoy/gpu"
	"github.com/lthomasmp/GBoy/mmu"
)

type GB struct {
	Cpu *cpu.CPU
	Mmu *mmu.MMU
	Gpu *gpu.GPU
}

func NewGB() *GB {
	return &GB{
		Cpu: cpu.NewCPU(),
		Mmu: mmu.NewMMU(),
		Gpu: gpu.NewGPU(),
	}
}

func (gb *GB) Read8(addr uint16) byte {
	switch addr & 0xF000 {

	// BIOS area
	case 0x0000:
		if gb.Mmu.InBios == 1 {
			if addr < 0x0100 {
				return gb.Mmu.Bios[addr]
			} else if gb.Cpu.Reg.PC == 0x0100 {
				gb.Mmu.InBios = 0
			}
			return gb.Mmu.Rom[addr]
		}

		// ROM 0 area
	case 0x1000, 0x2000, 0x3000:
		return gb.Mmu.Rom[addr]

		// ROM 1 area
	case 0x4000, 0x5000, 0x6000, 0x7000:
		return gb.Mmu.Rom[addr]

		//  Graphics VRAM
	case 0x8000, 0x9000:
		return gb.Gpu.Vram[addr]

		// External RAM
	case 0xa000, 0xb000:
		return gb.Mmu.Eram[addr&0x1FFF]

		// Working RAM
	case 0xc000, 0xd000:
		return gb.Mmu.Wram[addr&0x1FFF]

		// Working RAM shadow
	case 0xe000:
		return gb.Mmu.Wram[addr&0x1FFF]

		// Working RAM shadow, I/O, Zero-page RAM
	case 0xF000:
		switch addr & 0x0F00 {

		// Working RAM shadow
		case 0x000, 0x100, 0x200, 0x300, 0x400, 0x500, 0x600, 0x700, 0x800, 0x900, 0xa00, 0xb00, 0xc00, 0xd00:
			return gb.Mmu.Wram[addr&0xFF]

			// Graphics: object attribute memory
		case 0xe00:
			if addr < 0xFEA0 {
				return gb.Gpu.Oam[addr&0xFF]
			} else {
				return 0
			}

			//Zero-Page
		case 0xf00:
			if addr >= 0xff80 {
				return gb.Mmu.Zram[addr&0x7F]
			} else {
				// TODO: I/0 handling
				return 0
			}
		}
	}
	return 0
}

func (gb *GB) Read16(addr uint16) uint16 {
	// FIXME: shifting a 8-bit value of 8 bits ??
	return uint16(gb.Read8(addr) + (gb.Read8(addr+1) << 8))
}

func (gb *GB) push(m mmu.MMU, r byte) {
	gb.Cpu.Reg.SP--
	m.Write8(gb.Cpu.Reg.SP, r)
	gb.Cpu.IncrementInternalClockNTime(2)
}

func (gb *GB) pop(r byte) {
	r = gb.Read8(gb.Cpu.Reg.SP)
	gb.Cpu.Reg.SP++
	gb.Cpu.IncrementInternalClockNTime(2)
}

func (gb *GB) Dispatch() {
	opcode := gb.Read8(gb.Cpu.Reg.PC)
	opcodeToInstruction := gb.Cpu.MapOpcode()
	opcodeToInstruction[opcode]()
	gb.Cpu.Reg.PC &= 65535
	gb.Cpu.InternalClock.LastInstructionDuration += gb.Cpu.Reg.InstructionClock.LastInstructionDuration
	gb.Cpu.InternalClock.TotalDuration += gb.Cpu.Reg.InstructionClock.LastInstructionDuration
}
