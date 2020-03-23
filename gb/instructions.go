package gb

func (cpu *CPU) add8(A, B byte) byte {
	A += B
	cpu.Reg.F = 0
	if A&255 != 0 {
		cpu.Reg.F |= 0x80
	}
	if A > 255 {
		cpu.Reg.F |= 0x10
	}
	A &= 255
	cpu.IncrementInternalClockNTime(1)
	return A
}

func (cpu *CPU) nop() {
	cpu.IncrementInternalClockNTime(1)
}

func (cpu *CPU) push(m MMU, r byte) {
	cpu.Reg.SP--
	m.Write8(cpu.Reg.SP, r)
	cpu.IncrementInternalClockNTime(2)
}

func (gb *GB) pop(r byte) {
	r = gb.Read8(gb.Cpu.Reg.SP)
	gb.Cpu.Reg.SP++
	gb.Cpu.IncrementInternalClockNTime(2)
}
