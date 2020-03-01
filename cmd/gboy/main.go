package main

import (
	gb2 "github.com/lthomasmp/GBoy/gb/gb"
)

func main() {
	gb := gb2.NewGB()

	go gb.Dispatch()
}
