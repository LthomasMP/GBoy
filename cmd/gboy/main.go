package main

import (
	"github.com/lthomasmp/GBoy/gb"
)

func main() {
	gb := gb.NewGB()

	go gb.Dispatch()
}
