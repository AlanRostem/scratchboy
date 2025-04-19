#include <string.h>

#include "register_file.h"

namespace scr
{
    using registerArray = Word[EToI(Register::Count)];
    void RegisterFile_setValueLikeArray(RegisterFile *self, Register reg, Word value)
    {
        registerArray* asArray = reinterpret_cast<registerArray*>(self);
        *asArray[EToI(reg)] = value;
    }

    Word RegisterFile_getValueLikeArray(RegisterFile *self, Register reg)
    {
        registerArray* asArray = reinterpret_cast<registerArray*>(self);
        return *asArray[EToI(reg)];
    }

    void RegisterFile_setCombinedRegisters(RegisterFile *self, Register reg0, Register reg1, DoubleWord value)
    {
        Word combined[2]{};
        DoubleWord *combinedAsUint16 = reinterpret_cast<DoubleWord *>(combined);
        *combinedAsUint16 = value;

        RegisterFile_setValueLikeArray(self, reg0, combined[0]);
        RegisterFile_setValueLikeArray(self, reg1, combined[1]);
    }
    DoubleWord RegisterFile_getCombinedRegisterValue(RegisterFile *self, Register reg0, Register reg1)
    {
        Word combined[2]{};
        RegisterFile_setValueLikeArray(self, reg0, combined[0]);
        RegisterFile_setValueLikeArray(self, reg1, combined[1]);
        DoubleWord *combinedAsUint16 = reinterpret_cast<DoubleWord *>(combined);
        return *combinedAsUint16;
    }
}

void scr::RegisterFile_SetArithmeticFlag(RegisterFile *self, ArithmeticFlag flag, bool value)
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

void scr::RegisterFile_Init(RegisterFile *self)
{
    constexpr auto size = sizeof(self);
    memset(self, 0, size);
}

void scr::RegisterFile_SetValue(RegisterFile *self, Register reg, Word value)
{
    RegisterFile_setValueLikeArray(self, reg, value);
}

scr::Word scr::RegisterFile_GetValue(RegisterFile *self, Register reg)
{
    return RegisterFile_getValueLikeArray(self, reg);
}

void scr::RegisterFile_SetValueVirtual(RegisterFile *self, VirtualRegister reg, DoubleWord value)
{
    switch (reg)
    {
    case VirtualRegister::AF:
        RegisterFile_setCombinedRegisters(self, Register::A, Register::F, value);
        break;
    case VirtualRegister::DE:
        RegisterFile_setCombinedRegisters(self, Register::D, Register::E, value);
        break;
    case VirtualRegister::BC:
        RegisterFile_setCombinedRegisters(self, Register::B, Register::C, value);
        break;
    case VirtualRegister::HL:
        RegisterFile_setCombinedRegisters(self, Register::H, Register::L, value);
        break;
    }
}

scr::DoubleWord scr::RegisterFile_GetValueVirtual(RegisterFile *self, VirtualRegister reg)
{
    switch (reg)
    {
    case VirtualRegister::AF:
        return RegisterFile_getCombinedRegisterValue(self, Register::A, Register::F);
    case VirtualRegister::DE:
        return RegisterFile_getCombinedRegisterValue(self, Register::D, Register::E);
    case VirtualRegister::BC:
        return RegisterFile_getCombinedRegisterValue(self, Register::B, Register::C);
    case VirtualRegister::HL:
        return RegisterFile_getCombinedRegisterValue(self, Register::H, Register::L);
    }
    return DoubleWord(-1); // this will never happen thanks to type safety
}
