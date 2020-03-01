package gb

type GPU struct {
	Vram []byte
	Oam  []byte
}

func NewGPU() *GPU {
	return &GPU{
		Vram: []byte{0},
		Oam:  []byte{0},
	}
}
