package nums

type DByte uint16

func (db DByte) Split() (Byte, Byte) {
	left := Byte(db >> 8)
	right := Byte(db & 0x00FF)
	return left, right
}

func (db DByte) Uint16() uint16 {
	return uint16(db)
}
