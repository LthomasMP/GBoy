package gb

type GB struct {
	Cpu    *CPU
	Mmu    *MMU
	Gpu    *GPU
	Screen *Screen
}

func NewGB() *GB {
	return &GB{
		Cpu:    NewCPU(),
		Mmu:    NewMMU(),
		Gpu:    NewGPU(),
		Screen: NewScreen(),
	}
}
