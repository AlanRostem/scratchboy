package register

import "github.com/AlanRostem/scratchboy/nums"

type VirtualRegister nums.Byte

const (
	AF = VirtualRegister(iota)
	BC
	DE
	HL
)
