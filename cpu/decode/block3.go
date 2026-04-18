package decode

import "github.com/AlanRostem/scratchboy/nums"

type block3Opcode nums.Byte

func (o block3Opcode) Decode() (Info, error) {
	if o&0b00000_111 == 0b00000_110 {

	}
	return Info{}, nil
}
