package cpu

import "github.com/AlanRostem/scratchboy/types"

type R8Operand types.Word

const (
	R8OperB = R8Operand(iota)
	R8OperC
	R8OperD
	R8OperE
	R8OperH
	R8OperL
	R8OperHL
	R8OperA
	R8OperInvalid = 255
)

func (r8oper R8Operand) ToReg() Register {
	switch r8oper {
	case R8OperB:
		return RegB
	case R8OperC:
		return RegC
	case R8OperD:
		return RegD
	case R8OperE:
		return RegE
	case R8OperH:
		return RegH
	case R8OperL:
		return RegL
	case R8OperHL:
		return R8OperInvalid
	case R8OperA:
		return RegA
	}
	return R8OperInvalid
}
