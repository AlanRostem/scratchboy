#include <string.h>

#include "register_set.h"

namespace scr
{
    void RegisterSet_setCombinedRegisters(RegisterSet self, Operand reg0, Operand reg1, VirtualWord value);
    VirtualWord RegisterSet_getCombinedRegisterValue(RegisterSet self, Operand reg0, Operand reg1);
}

void scr::RegisterSet_setCombinedRegisters(RegisterSet self, Operand reg0, Operand reg1, VirtualWord value)
{
    Word combined[2];
    VirtualWord* combinedAsUint16 = reinterpret_cast<VirtualWord*>(combined);
    *combinedAsUint16 = value;
    self[EToI(reg0)] = combined[0];
    self[EToI(reg1)] = combined[1];
}

scr::VirtualWord scr::RegisterSet_getCombinedRegisterValue(RegisterSet self, Operand reg0, Operand reg1)
{
    Word combined[2];
    combined[0] = self[EToI(reg0)];
    combined[1] = self[EToI(reg1)];
    VirtualWord* combinedAsUint16 = reinterpret_cast<VirtualWord*>(combined);
    return *combinedAsUint16;
}

void scr::RegisterSet_SetArithmeticFlag(RegisterSet self, ArithmeticFlag flag, bool value)
{
    if (value)
    {
        // set to 1
        self[EToI(Operand::F)] |= static_cast<Word>(flag);
        return;
    }
    // set to 0
    self[EToI(Operand::F)] &= ~static_cast<Word>(flag);
}

void scr::RegisterSet_Init(RegisterSet self)
{
    constexpr auto size = sizeof(self);
    memset(self, 0, size);
}

void scr::RegisterSet_SetValue(RegisterSet self, Operand reg, Word value)
{
    self[EToI(reg)] = value;
}

scr::Word scr::RegisterSet_GetValue(RegisterSet self, Operand reg)
{
    return self[EToI(reg)];
}

void scr::RegisterSet_SetValueVirtual(RegisterSet self, Operand reg, VirtualWord value)
{
    switch (reg)
    {
    case Operand::AF:
        RegisterSet_setCombinedRegisters(self, Operand::A, Operand::F, value);
        break;
    case Operand::DE:
        RegisterSet_setCombinedRegisters(self, Operand::D, Operand::E, value);
        break;
    case Operand::BC:
        RegisterSet_setCombinedRegisters(self, Operand::B, Operand::C, value);
        break;
    case Operand::HL:
        RegisterSet_setCombinedRegisters(self, Operand::H, Operand::L, value);
        break;
    }
}

scr::VirtualWord scr::RegisterSet_GetValueVirtual(RegisterSet self, Operand reg)
{
    switch (reg)
    {
    case Operand::AF:
        return RegisterSet_getCombinedRegisterValue(self, Operand::A, Operand::F);
    case Operand::DE:
        return RegisterSet_getCombinedRegisterValue(self, Operand::D, Operand::E);
    case Operand::BC:
        return RegisterSet_getCombinedRegisterValue(self, Operand::B, Operand::C);
    case Operand::HL:
        return RegisterSet_getCombinedRegisterValue(self, Operand::H, Operand::L);
    }
}
