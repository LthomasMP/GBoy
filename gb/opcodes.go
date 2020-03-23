package gb

func (cpu *CPU) MapOpcode() [0x100]func() {
	opcodeToInstruction := [0x100]func(){
		0x00: func() {
			cpu.nop()
		},
	}

	return opcodeToInstruction
}
