#include <string.h>

#include "register_bank.h"

namespace scr
{
    using registerArray = Word[EToI(Register::Count)];
    void RegisterBank_setValueLikeArray(RegisterBank *self, Register reg, Word value)
    {
        registerArray* asArray = reinterpret_cast<registerArray*>(self);
        *asArray[EToI(reg)] = value;
    }

    Word RegisterBank_getValueLikeArray(RegisterBank *self, Register reg)
    {
        registerArray* asArray = reinterpret_cast<registerArray*>(self);
        return *asArray[EToI(reg)];
    }

    void RegisterBank_setCombinedRegisters(RegisterBank *self, Register reg0, Register reg1, VirtualWord value)
    {
        Word combined[2]{};
        VirtualWord *combinedAsUint16 = reinterpret_cast<VirtualWord *>(combined);
        *combinedAsUint16 = value;

        RegisterBank_setValueLikeArray(self, reg0, combined[0]);
        RegisterBank_setValueLikeArray(self, reg1, combined[1]);
    }
    VirtualWord RegisterBank_getCombinedRegisterValue(RegisterBank *self, Register reg0, Register reg1)
    {
        Word combined[2]{};
        RegisterBank_setValueLikeArray(self, reg0, combined[0]);
        RegisterBank_setValueLikeArray(self, reg1, combined[1]);
        VirtualWord *combinedAsUint16 = reinterpret_cast<VirtualWord *>(combined);
        return *combinedAsUint16;
    }
}

void scr::RegisterBank_SetArithmeticFlag(RegisterBank *self, ArithmeticFlag flag, bool value)
{
    if (value)
    {
        // set to 1
        self->rawF |= static_cast<Word>(flag);
        return;
    }
    // set to 0
    self->rawF &= ~static_cast<Word>(flag);
}

void scr::RegisterBank_Init(RegisterBank *self)
{
    constexpr auto size = sizeof(self);
    memset(self, 0, size);
}

void scr::RegisterBank_SetValue(RegisterBank *self, Register reg, Word value)
{
    RegisterBank_setValueLikeArray(self, reg, value);
}

scr::Word scr::RegisterBank_GetValue(RegisterBank *self, Register reg)
{
    return RegisterBank_getValueLikeArray(self, reg);
}

void scr::RegisterBank_SetValueVirtual(RegisterBank *self, VirtualRegister reg, VirtualWord value)
{
    switch (reg)
    {
    case VirtualRegister::AF:
        RegisterBank_setCombinedRegisters(self, Register::A, Register::F, value);
        break;
    case VirtualRegister::DE:
        RegisterBank_setCombinedRegisters(self, Register::D, Register::E, value);
        break;
    case VirtualRegister::BC:
        RegisterBank_setCombinedRegisters(self, Register::B, Register::C, value);
        break;
    case VirtualRegister::HL:
        RegisterBank_setCombinedRegisters(self, Register::H, Register::L, value);
        break;
    }
}

scr::VirtualWord scr::RegisterBank_GetValueVirtual(RegisterBank *self, VirtualRegister reg)
{
    switch (reg)
    {
    case VirtualRegister::AF:
        return RegisterBank_getCombinedRegisterValue(self, Register::A, Register::F);
    case VirtualRegister::DE:
        return RegisterBank_getCombinedRegisterValue(self, Register::D, Register::E);
    case VirtualRegister::BC:
        return RegisterBank_getCombinedRegisterValue(self, Register::B, Register::C);
    case VirtualRegister::HL:
        return RegisterBank_getCombinedRegisterValue(self, Register::H, Register::L);
    }
    return VirtualWord(-1); // this will never happen thanks to type safety
}
