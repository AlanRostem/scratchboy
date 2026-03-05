package cpu

type ArithmeticFlag uint8

const (
	ArithFlagCarry     = 1 << ArithmeticFlag(4)
	ArithFlagHalfCarry = 1 << ArithmeticFlag(5)
	ArithFlagSubtract  = 1 << ArithmeticFlag(6)
	ArithFlagZero      = 1 << ArithmeticFlag(7)
)
