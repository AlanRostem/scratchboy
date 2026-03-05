package cpu

import "github.com/AlanRostem/scratchboy/types"

type Register types.Word

const (
	RegA = Register(iota)
	RegF
	RegB
	RegC
	RegD
	RegE
	RegH
	RegL
	RegCount
)
