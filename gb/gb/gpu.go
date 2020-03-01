package gb

type GPU struct {
	Vram      []byte
	Oam       []byte
	Mode      int
	Modeclock int
	Line      int
}

func NewGPU() *GPU {
	return &GPU{
		Vram:      []byte{0},
		Oam:       []byte{0},
		Mode:      0,
		Modeclock: 0,
		Line:      0,
	}
}

func (gpu *GPU) ClockStep() {

}
