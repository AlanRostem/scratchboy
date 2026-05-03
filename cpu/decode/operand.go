package decode

const (
	R8B = iota
	R8C
	R8D
	R8E
	R8H
	R8L
	R8HL
	R8A
)

const (
	R16BC = iota
	R16DE
	R16HL
	R16SP
)

const (
	R16StkBC = iota
	R16StkDE
	R16StkHL
	R16StkAF
)

const (
	R16MemBC = iota
	R16MemDE
	R16MemHLI
	R16MemHLD
)

const (
	CondNz = iota
	CondZ
	CondNc
	CondC
)

const (
	Tgt3Vec00 = iota
	Tgt3Vec08
	Tgt3Vec10
	Tgt3Vec18
	Tgt3Vec20
	Tgt3Vec28
	Tgt3Vec30
	Tgt3Vec38
)
