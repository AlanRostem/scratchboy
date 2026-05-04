package register

import "github.com/AlanRostem/scratchboy/nums"

type GeneralPurposeRegister nums.Byte

const (
	A = GeneralPurposeRegister(iota)
	F
	B
	C
	D
	E
	H
	L
	GeneralPurposeRegisterCount
)
