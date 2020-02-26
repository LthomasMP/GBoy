package cpu

import "github.com/lthomasmp/GBoy/mmu"

type CPU struct {
	InternalClock Clock
	Reg           Registers
}

func NewCPU() *CPU {
	return &CPU{
		InternalClock: *NewClock(),
		Reg:           *NewRegisters(),
	}
}

func (cpu *CPU) Dispatch(m mmu.MMU) {
	opcode := m.Read8(cpu.Reg.PC)
	opcodeToInstruction := cpu.mapOpcode()
	opcodeToInstruction[opcode]()
	cpu.Reg.PC &= 65535
	cpu.InternalClock.LastInstructionDuration += cpu.Reg.InstructionClock.LastInstructionDuration
	cpu.InternalClock.TotalDuration += cpu.Reg.InstructionClock.LastInstructionDuration
}
