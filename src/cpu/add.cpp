#include "emu_register_set.h"

using namespace scr;
void scr_EmuRegisterSet_ExecuteADD(EmuRegisterSet self, EmuRegister target)
{
    Word value = self[EToI(target)];
    Word* aReg = &self[EToI(EmuRegister::A)];
    // emulated overflow check:
    uint16_t overFlowCheckValue =
        static_cast<uint16_t>(*aReg) +
        static_cast<uint16_t>(value);
    *aReg += value; // add the value and allow overflow

    EmuRegisterSet_SetArithmeticFlag(self, EmuArithmeticFlag::Subtract, false);
    EmuRegisterSet_SetArithmeticFlag(self, EmuArithmeticFlag::Zero, *aReg == 0);
    EmuRegisterSet_SetArithmeticFlag(self, EmuArithmeticFlag::Carry, overFlowCheckValue > 255);
    EmuRegisterSet_SetArithmeticFlag(self, EmuArithmeticFlag::HalfCarry,
        ((*aReg & 0xF) + (value & 0xF)) > 0xF
    );
}
