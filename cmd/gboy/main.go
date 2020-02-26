package main

import (
	"github.com/lthomasmp/GBoy/cpu"
	"github.com/lthomasmp/GBoy/mmu"
)

func main() {
	mem := mmu.NewMMU()
	cpu := cpu.NewCPU()

	go cpu.Dispatch(mem)
}
