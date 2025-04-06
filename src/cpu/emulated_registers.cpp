#include <string.h>

#include "emulated_registers.h"

namespace scr
{
    void EmulatedRegisters_setCombinedRegisters(EmulatedRegisters self, Register reg0, Register reg1, VirtualWord value);
    VirtualWord EmulatedRegisters_getCombinedRegisterValue(EmulatedRegisters self, Register reg0, Register reg1);
}

void scr::EmulatedRegisters_setCombinedRegisters(EmulatedRegisters self, Register reg0, Register reg1, VirtualWord value)
{
    Word combined[2];
    VirtualWord* combinedAsUint16 = reinterpret_cast<VirtualWord*>(combined);
    *combinedAsUint16 = value;
    self[EToI(reg0)] = combined[0];
    self[EToI(reg1)] = combined[1];
}

scr::VirtualWord scr::EmulatedRegisters_getCombinedRegisterValue(EmulatedRegisters self, Register reg0, Register reg1)
{
    Word combined[2];
    combined[0] = self[EToI(reg0)];
    combined[1] = self[EToI(reg1)];
    VirtualWord* combinedAsUint16 = reinterpret_cast<VirtualWord*>(combined);
    return *combinedAsUint16;
}

void scr::EmulatedRegisters_SetArithmeticFlag(EmulatedRegisters self, ArithmeticFlag flag, bool value)
{
    if (value)
    {
        // set to 1
        self[EToI(Register::F)] |= static_cast<Word>(flag);
        return;
    }
    // set to 0
    self[EToI(Register::F)] &= ~static_cast<Word>(flag);
}

void scr::EmulatedRegisters_Init(EmulatedRegisters self)
{
    constexpr auto size = sizeof(self);
    memset(self, 0, size);
}

void scr::EmulatedRegisters_SetValue(EmulatedRegisters self, Register reg, Word value)
{
    self[EToI(reg)] = value;
}

scr::Word scr::EmulatedRegisters_GetValue(EmulatedRegisters self, Register reg)
{
    return self[EToI(reg)];
}

void scr::EmulatedRegisters_SetValueVirtual(EmulatedRegisters self, VirtualRegister reg, VirtualWord value)
{
    switch (reg)
    {
    case VirtualRegister::DE:
        EmulatedRegisters_setCombinedRegisters(self, Register::D, Register::E, value);
        break;
    case VirtualRegister::BC:
        EmulatedRegisters_setCombinedRegisters(self, Register::B, Register::C, value);
        break;
    case VirtualRegister::HL:
        EmulatedRegisters_setCombinedRegisters(self, Register::H, Register::L, value);
        break;
    }
}

scr::VirtualWord scr::EmulatedRegisters_GetValueVirtual(EmulatedRegisters self, VirtualRegister reg)
{
    switch (reg)
    {
    case VirtualRegister::DE:
        return EmulatedRegisters_getCombinedRegisterValue(self, Register::D, Register::E);
    case VirtualRegister::BC:
        return EmulatedRegisters_getCombinedRegisterValue(self, Register::B, Register::C);
    case VirtualRegister::HL:
        return EmulatedRegisters_getCombinedRegisterValue(self, Register::H, Register::L);
    }
}
