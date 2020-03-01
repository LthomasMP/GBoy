package gb

type GB struct {
	Cpu *CPU
	Mmu *MMU
	Gpu *GPU
}

func NewGB() *GB {
	return &GB{
		Cpu: NewCPU(),
		Mmu: NewMMU(),
		Gpu: NewGPU(),
	}
}
