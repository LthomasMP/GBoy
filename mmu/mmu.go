package mmu

type Byte uint8

type Word uint16

type MMU struct {
	InBios int
	Rom    []Byte
	Bios   []Byte
	Wram   []Byte
	Eram   []Byte
	Zram   []Byte
}

func NewMMU() *MMU {
	return &MMU{
		InBios: 1,
		Rom:    []Byte{0},
		Bios:   []Byte{0},
		Wram:   []Byte{0},
		Eram:   []Byte{0},
		Zram:   []Byte{0},
	}
}

func (m *MMU) Write8(addr Word, val Byte) {
	// Do something
}

func (m *MMU) Write16(addr Word, val Word) {
	// Do something
}

func (m *MMU) LoadROM(file string) {
	// DO something
}
