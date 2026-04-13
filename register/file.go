package register

import "github.com/AlanRostem/scratchboy/nums"

type RegisterFile struct {
	registers [RegisterCount]nums.Byte
}

func NewFile() RegisterFile {
	return RegisterFile{}
}

func (rf *RegisterFile) Set(reg Register, value nums.Byte) {
	rf.registers[reg] = value
}

func (rf *RegisterFile) Get(reg Register) nums.Byte {
	return rf.registers[reg]
}

func (rf *RegisterFile) VSet(vreg VirtualRegister, value nums.DByte) {
	r0, r1 := splitRegisters(vreg)
	v0, v1 := value.Split()
	rf.registers[r0] = v0
	rf.registers[r1] = v1
}

func (rf *RegisterFile) VGet(vreg VirtualRegister) nums.DByte {
	r0, r1 := splitRegisters(vreg)
	v0 := rf.registers[r0]
	v1 := rf.registers[r1]
	return v0.Concat(v1)
}

func splitRegisters(vreg VirtualRegister) (Register, Register) {
	switch vreg {
	case AF:
		return A, F
	case HL:
		return H, L
	case BC:
		return B, C
	case DE:
		return D, E
	default:
		return RegisterCount, RegisterCount
	}
}
