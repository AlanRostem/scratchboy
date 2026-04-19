package decode

const (
	b3ArithImmAdd = 0b00_000_000
	b3ArithImmAdc = 0b00_001_000
	b3ArithImmSub = 0b00_010_000
	b3ArithImmSbc = 0b00_011_000
	b3ArithImmAnd = 0b00_100_000
	b3ArithImmXor = 0b00_101_000
	b3ArithImmOr  = 0b00_110_000
	b3ArithImmCp  = 0b00_111_000
)

const (
	b3CBPrefix = 0b11001011
)

const (
	b3NoPatternLdhCA      = 0b000_00010
	b3NoPatternLdhImm8A   = 0b000_00000
	b3NoPatternLdhImm16A  = 0b000_01010
	b3NoPatternLdhAC      = 0b000_10010
	b3NoPatternLdhAImm8   = 0b000_10000
	b3NoPatternLdhAImm16  = 0b000_11010
	b3NoPatternAddSpImm8  = 0b000_01000
	b3NoPatternLdHlSpImm8 = 0b000_11000
	b3NoPatternLdSpHl     = 0b000_11001
)

const (
	b3StackPop  = 0b00_00_0001
	b3StackPush = 0b00_00_0101
)

const (
	b3ArithIdBits     = 0b00000_110
	b3InterruptIdBits = 0b00000_011
)

const (
	b3NoPatternBitMask = 0b000_11111
)

const (
	b3InterruptDi = 0b00_110_000
	b3InterruptEi = 0b00_111_000
)
