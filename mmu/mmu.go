package mmu

type Byte uint8

type Word uint16

type MMU interface {
	Read8(addr Word) Byte
	Read16(addr Word) Word
	Write8(addr Word, val Byte)
	Write16(addr Word, val Word)
}

type mmu struct {
}

func NewMMU() MMU {
	return &mmu{}
}

func (m *mmu) Read8(addr Word) Byte {
	var tmp Byte
	return tmp
}

func (m *mmu) Read16(addr Word) Word {
	var tmp Word
	return tmp
}

func (m *mmu) Write8(addr Word, val Byte) {

}

func (m *mmu) Write16(addr Word, val Word) {

}
