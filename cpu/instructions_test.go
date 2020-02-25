package cpu

import (
	"testing"
	"time"
)

func TestCPU_Add8(t *testing.T) {
	cpu := NewCPU()
	cpu.Reg.A = 42
	cpu.Reg.B = 43
	// TODO: Create a structure Register to simply such operations
	cpu.Reg.A = cpu.Add8(cpu.Reg.A, cpu.Reg.B)
	if cpu.Reg.A != 85 {
		t.Errorf("expected %v in reg A", 85)
	}
	if cpu.Reg.B != 43 {
		t.Errorf("expected %v in reg B", 85)
	}
	if cpu.Reg.InstructionClock.TotalDuration != 4*time.Millisecond {
		t.Errorf("expected %v as duration", 4)
	}
	if cpu.Reg.InstructionClock.LastInstructionDuration != 1*time.Millisecond {
		t.Errorf("expected %v as duration", 1)
	}
}
