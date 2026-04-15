package decode

import (
	"fmt"

	num "github.com/AlanRostem/scratchboy/nums"
)

type Block0Opcode num.Byte

func (o Block0Opcode) Decode() (Info, error) {
	return Info{}, fmt.Errorf("opcode could not be decoded: %X", o)
}
