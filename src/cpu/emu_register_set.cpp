#include <string.h>

#include "emu_register_set.h"

namespace scr
{
    void EmuRegisterSet_setCombinedRegisters(EmuRegisterSet self, EmuRegister reg0, EmuRegister reg1, VirtualWord value);
    VirtualWord EmuRegisterSet_getCombinedRegisterValue(EmuRegisterSet self, EmuRegister reg0, EmuRegister reg1);
}

void scr::EmuRegisterSet_setCombinedRegisters(EmuRegisterSet self, EmuRegister reg0, EmuRegister reg1, VirtualWord value)
{
    Word combined[2];
    VirtualWord* combinedAsUint16 = reinterpret_cast<VirtualWord*>(combined);
    *combinedAsUint16 = value;
    self[EToI(reg0)] = combined[0];
    self[EToI(reg1)] = combined[1];
}

scr::VirtualWord scr::EmuRegisterSet_getCombinedRegisterValue(EmuRegisterSet self, EmuRegister reg0, EmuRegister reg1)
{
    Word combined[2];
    combined[0] = self[EToI(reg0)];
    combined[1] = self[EToI(reg1)];
    VirtualWord* combinedAsUint16 = reinterpret_cast<VirtualWord*>(combined);
    return *combinedAsUint16;
}

void scr::EmuRegisterSet_SetArithmeticFlag(EmuRegisterSet self, EmuArithmeticFlag flag, bool value)
{
    if (value)
    {
        // set to 1
        self[EToI(EmuRegister::F)] |= static_cast<Word>(flag);
        return;
    }
    // set to 0
    self[EToI(EmuRegister::F)] &= ~static_cast<Word>(flag);
}

void scr::EmuRegisterSet_Init(EmuRegisterSet self)
{
    constexpr auto size = sizeof(self);
    memset(self, 0, size);
}

void scr::EmuRegisterSet_SetValue(EmuRegisterSet self, EmuRegister reg, Word value)
{
    self[EToI(reg)] = value;
}

scr::Word scr::EmuRegisterSet_GetValue(EmuRegisterSet self, EmuRegister reg)
{
    return self[EToI(reg)];
}

void scr::EmuRegisterSet_SetValueVirtual(EmuRegisterSet self, EmuVirtualRegister reg, VirtualWord value)
{
    switch (reg)
    {
    case EmuVirtualRegister::AF:
        EmuRegisterSet_setCombinedRegisters(self, EmuRegister::A, EmuRegister::F, value);
        break;
    case EmuVirtualRegister::DE:
        EmuRegisterSet_setCombinedRegisters(self, EmuRegister::D, EmuRegister::E, value);
        break;
    case EmuVirtualRegister::BC:
        EmuRegisterSet_setCombinedRegisters(self, EmuRegister::B, EmuRegister::C, value);
        break;
    case EmuVirtualRegister::HL:
        EmuRegisterSet_setCombinedRegisters(self, EmuRegister::H, EmuRegister::L, value);
        break;
    }
}

scr::VirtualWord scr::EmuRegisterSet_GetValueVirtual(EmuRegisterSet self, EmuVirtualRegister reg)
{
    switch (reg)
    {
    case EmuVirtualRegister::AF:
        return EmuRegisterSet_getCombinedRegisterValue(self, EmuRegister::A, EmuRegister::F);
    case EmuVirtualRegister::DE:
        return EmuRegisterSet_getCombinedRegisterValue(self, EmuRegister::D, EmuRegister::E);
    case EmuVirtualRegister::BC:
        return EmuRegisterSet_getCombinedRegisterValue(self, EmuRegister::B, EmuRegister::C);
    case EmuVirtualRegister::HL:
        return EmuRegisterSet_getCombinedRegisterValue(self, EmuRegister::H, EmuRegister::L);
    }
}
