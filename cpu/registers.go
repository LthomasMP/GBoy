package cpu

type Register interface {

}

type Register8 uint8

type Register16 uint16

type Registers struct {
	A Register8
	B Register8
	C Register8
	D Register8
	E Register8
	H Register8
	L Register8
	F Register8
	PC Register16
	SP Register16
	InstructionClock Clock
}
