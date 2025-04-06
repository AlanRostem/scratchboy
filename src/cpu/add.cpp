#include "instructions.h"

void scr::CPU_ExecuteADD(CPU* self, Register target)
{
    Word value = self->registers[EToI(target)];
    Word* aReg = &self->registers[EToI(Register::A)];
    // emulated overflow check:
    uint16_t overFlowCheckValue =
        static_cast<uint16_t>(*aReg) +
        static_cast<uint16_t>(value);
    *aReg += value; // add the value and allow overflow

    CPU_SetArithmeticFlag(self, ArithmeticFlag::Subtract, false);
    CPU_SetArithmeticFlag(self, ArithmeticFlag::Zero, *aReg == 0);
    CPU_SetArithmeticFlag(self, ArithmeticFlag::Carry, overFlowCheckValue > 255);
    CPU_SetArithmeticFlag(self, ArithmeticFlag::HalfCarry,
        ((*aReg & 0xF) + (value & 0xF)) > 0xF
    );
}
