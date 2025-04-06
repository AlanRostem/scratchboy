#include "emulated_registers.h"

using namespace scr;
void scr_EmulatedRegisters_ExecuteADD(EmulatedRegisters* self, Register target)
{
    Word value = self->registers[EToI(target)];
    Word* aReg = &self->registers[EToI(Register::A)];
    // emulated overflow check:
    uint16_t overFlowCheckValue =
        static_cast<uint16_t>(*aReg) +
        static_cast<uint16_t>(value);
    *aReg += value; // add the value and allow overflow

    EmulatedRegisters_SetArithmeticFlag(self, ArithmeticFlag::Subtract, false);
    EmulatedRegisters_SetArithmeticFlag(self, ArithmeticFlag::Zero, *aReg == 0);
    EmulatedRegisters_SetArithmeticFlag(self, ArithmeticFlag::Carry, overFlowCheckValue > 255);
    EmulatedRegisters_SetArithmeticFlag(self, ArithmeticFlag::HalfCarry,
        ((*aReg & 0xF) + (value & 0xF)) > 0xF
    );
}
