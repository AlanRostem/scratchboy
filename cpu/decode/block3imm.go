package decode

var b3ImmediateByteCounts = map[InstructionId]int{
	LdhImm8A:   1,
	LdImm16A:   2,
	LdhAImm8:   1,
	LdAImm16:   2,
	AddSpImm8:  1,
	LdHlSpImm8: 1,
}
