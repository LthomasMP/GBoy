package cpu

func (cpu CPU) Add(A, B Register8) {
	cpu.Reg.A = A
	cpu.Reg.B = B
	cpu.Reg.A += cpu.Reg.B
	cpu.Reg.F = 0
	if cpu.Reg.A == 0 {
		cpu.Reg.F |= 0x80
	}
	if cpu.Reg.A > 255 {
		cpu.Reg.F |= 0x10
	}
	cpu.Reg.A &= 255
	cpu.Reg.InstructionClock.LastInstructionDuration = 1
	cpu.Reg.InstructionClock.TotalDuration = 4
}
