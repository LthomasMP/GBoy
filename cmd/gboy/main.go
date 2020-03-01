package main

import (
	"golang.org/x/exp/shiny/driver"

	gb2 "github.com/lthomasmp/GBoy/gb/gb"
)

func main() {
	gb := gb2.NewGB()

	driver.Main(gb.Screen.Run)
}
