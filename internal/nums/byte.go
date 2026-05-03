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

func (b Byte) OctalRow() Byte {
	const octalMaskRow = 0b11111000
	octalRow := octalMaskRow & b
	octalRow >>= 3 // to get a full value
	return octalRow
}

func (b Byte) OctalCol() Byte {
	const octalMaskCol = 0b00000111
	octalCol := octalMaskCol & b
	return octalCol
}
