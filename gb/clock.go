package gb

import (
	"time"
)

type Clock struct {
	LastInstructionDuration time.Duration
	TotalDuration           time.Duration
}

func NewClock() *Clock {
	return &Clock{
		LastInstructionDuration: 0,
		TotalDuration:           0,
	}
}

func (cpu *CPU) IncrementInternalClockNTime(n time.Duration) {
	cpu.Reg.InstructionClock.LastInstructionDuration = n * time.Millisecond
	cpu.Reg.InstructionClock.TotalDuration = 4 * time.Millisecond
}
