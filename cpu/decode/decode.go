package decode

import num "github.com/AlanRostem/scratchboy/nums"

type Block num.Byte

const (
	Block0 = Block(iota)
	Block1
	Block2
	Block3
	BlockInvalid
)
