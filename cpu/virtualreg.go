package cpu

import "github.com/AlanRostem/scratchboy/types"

type VirtualRegister types.Word

const (
	VRegAF = VirtualRegister(iota)
	VRegBC
	VRegDE
	VRegHL
)
