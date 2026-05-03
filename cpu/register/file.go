package register

import "github.com/AlanRostem/scratchboy/nums"

type File struct {
	registers [RegisterCount]nums.Byte
}

func NewFile() *File {
	return &File{}
}

func (rf *File) Set(reg Register, value nums.Byte) {
	rf.registers[reg] = value
}

func (rf *File) Get(reg Register) nums.Byte {
	return rf.registers[reg]
}

func (rf *File) VSet(vreg VirtualRegister, value nums.DByte) {
	r0, r1 := splitRegisters(vreg)
	v0, v1 := value.Split()
	rf.registers[r0] = v0
	rf.registers[r1] = v1
}

func (rf *File) VGet(vreg VirtualRegister) nums.DByte {
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
