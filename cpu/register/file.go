package register

import "github.com/AlanRostem/scratchboy/nums"

type File struct {
	pc             nums.DByte
	sp             nums.DByte
	generalPurpose [GeneralPurposeRegisterCount]nums.Byte
}

func NewFile() *File {
	return &File{}
}

func (rf *File) IncSP() {
	rf.sp++
}

func (rf *File) DecSP() {
	rf.sp--
}

func (rf *File) SP() nums.DByte {
	return rf.sp
}

func (rf *File) IncPC() {
	rf.pc++
}

func (rf *File) SetPC(pc nums.DByte) {
	rf.pc = pc
}

func (rf *File) PC() nums.DByte {
	return rf.pc
}

func (rf *File) Set(reg GeneralPurposeRegister, value nums.Byte) {
	rf.generalPurpose[reg] = value
}

func (rf *File) Get(reg GeneralPurposeRegister) nums.Byte {
	return rf.generalPurpose[reg]
}

func (rf *File) VSet(vreg VirtualRegister, value nums.DByte) {
	r0, r1 := splitRegisters(vreg)
	v0, v1 := value.Split()
	rf.generalPurpose[r0] = v0
	rf.generalPurpose[r1] = v1
}

func (rf *File) VGet(vreg VirtualRegister) nums.DByte {
	r0, r1 := splitRegisters(vreg)
	v0 := rf.generalPurpose[r0]
	v1 := rf.generalPurpose[r1]
	return v0.Concat(v1)
}

func splitRegisters(vreg VirtualRegister) (GeneralPurposeRegister, GeneralPurposeRegister) {
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
		return GeneralPurposeRegisterCount, GeneralPurposeRegisterCount
	}
}
