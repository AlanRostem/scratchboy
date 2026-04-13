package register

import "github.com/AlanRostem/scratchboy/types"

type VirtualRegister types.Byte

const (
	VRegAF = VirtualRegister(iota)
	VRegBC
	VRegDE
	VRegHL
)
