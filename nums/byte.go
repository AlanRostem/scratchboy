package nums

type Byte uint8

func (left Byte) Concat(right Byte) DByte {
	full := DByte(left) << 8
	full |= DByte(right)
	return full
}

func (b Byte) Uint8() uint8 {
	return uint8(b)
}
