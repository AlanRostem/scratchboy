package register

import "github.com/AlanRostem/scratchboy/types"

type Register types.Byte

const (
	A = Register(iota)
	F
	B
	C
	D
	E
	H
	L
	RegisterCount
)
