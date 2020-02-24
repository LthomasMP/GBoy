package cpu

import "time"

type Clock struct {
	LastInstructionDuration time.Duration
	TotalDuration time.Duration
}
