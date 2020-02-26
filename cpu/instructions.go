package cpu

import (
	"github.com/lthomasmp/GBoy/mmu"
)

func (cpu *CPU) add8(A, B mmu.Byte) mmu.Byte {
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
