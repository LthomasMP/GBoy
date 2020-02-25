package cpu

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
