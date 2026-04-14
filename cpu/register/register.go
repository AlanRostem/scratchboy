package register

import "github.com/AlanRostem/scratchboy/nums"

type Register nums.Byte

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
