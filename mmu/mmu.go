package mmu

type MMU struct {
	InBios int
	Rom    []byte
	Bios   []byte
	Wram   []byte
	Eram   []byte
	Zram   []byte
}

func NewMMU() *MMU {
	return &MMU{
		InBios: 1,
		Rom:    []byte{0},
		Bios:   []byte{0},
		Wram:   []byte{0},
		Eram:   []byte{0},
		Zram:   []byte{0},
	}
}

func (m *MMU) Write8(addr uint16, val byte) {
	// Do something
}

func (m *MMU) Write16(addr uint16, val uint16) {
	// Do something
}

func (m *MMU) LoadROM(file string) {
	// Do something
}
