package cpu

import "github.com/lthomasmp/GBoy/mmu"

type Registers struct {
	A                mmu.Byte
	B                mmu.Byte
	C                mmu.Byte
	D                mmu.Byte
	E                mmu.Byte
	H                mmu.Byte
	L                mmu.Byte
	F                mmu.Byte
	PC               mmu.Word
	SP               mmu.Word
	InstructionClock Clock
}

func NewRegisters() *Registers {
	return &Registers{
		A:                0,
		B:                0,
		C:                0,
		D:                0,
		E:                0,
		H:                0,
		L:                0,
		F:                0,
		PC:               0,
		SP:               0,
		InstructionClock: *NewClock(),
	}
}

func (cpu *CPU) ResetRegisters() {
	cpu.Reg.A = 0
	cpu.Reg.B = 0
	cpu.Reg.C = 0
	cpu.Reg.D = 0
	cpu.Reg.E = 0
	cpu.Reg.H = 0
	cpu.Reg.L = 0
	cpu.Reg.F = 0
	cpu.Reg.PC = 0
	cpu.Reg.SP = 0
	cpu.Reg.A = 0
	cpu.InternalClock.LastInstructionDuration = 0
	cpu.InternalClock.TotalDuration = 0
}
