#include "register_set.h"

using namespace scr;
void scr_RegisterSet_ExecuteADD(RegisterSet self, Operand target)
{
    Word value = self[EToI(target)];
    Word* aReg = &self[EToI(Operand::A)];
    // emulated overflow check:
    uint16_t overFlowCheckValue =
        static_cast<uint16_t>(*aReg) +
        static_cast<uint16_t>(value);
    *aReg += value; // add the value and allow overflow

    RegisterSet_SetArithmeticFlag(self, ArithmeticFlag::Subtract, false);
    RegisterSet_SetArithmeticFlag(self, ArithmeticFlag::Zero, *aReg == 0);
    RegisterSet_SetArithmeticFlag(self, ArithmeticFlag::Carry, overFlowCheckValue > 255);
    RegisterSet_SetArithmeticFlag(self, ArithmeticFlag::HalfCarry,
        ((*aReg & 0xF) + (value & 0xF)) > 0xF
    );
}
