package cpu

import (
	"github.com/lthomasmp/GBoy/mmu"
)

func (cpu *CPU) Add8(A, B mmu.Byte) mmu.Byte {
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

func (cpu *CPU) NOP() {
	cpu.IncrementInternalClockNTime(1)
}

func (cpu *CPU) Push(m mmu.MMU, r mmu.Byte) {
	cpu.Reg.SP--
	m.Write8(cpu.Reg.SP, r)
	cpu.IncrementInternalClockNTime(2)
}

func (cpu *CPU) Pop(m mmu.MMU, r mmu.Byte) {
	r = m.Read8(cpu.Reg.SP)
	cpu.Reg.SP++
	cpu.IncrementInternalClockNTime(2)
}
