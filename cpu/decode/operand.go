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
