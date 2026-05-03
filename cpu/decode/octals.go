package decode

import "github.com/AlanRostem/scratchboy/nums"

func isInOctalRowInterval(opcodeByte, octalRowStart, octalRowEnd nums.Byte) bool {
	octalRow := opcodeByte.OctalRow()
	return octalRow >= octalRowStart && octalRow <= octalRowEnd
}

func isInOctalColInterval(opcodeByte, octalColStart, octalColEnd nums.Byte) bool {
	octalCol := opcodeByte.OctalCol()
	return octalCol >= octalColStart && octalCol <= octalColEnd
}
