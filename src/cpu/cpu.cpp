#include <string.h>

#include "cpu.h"

#include "instructions.h"

namespace scr
{
    void CPU_setCombinedRegisters(CPU* self, Register reg0, Register reg1, VirtualWord value);
    VirtualWord CPU_getCombinedRegisterValue(CPU* self, Register reg0, Register reg1);
}

void scr::CPU_setCombinedRegisters(CPU* self, Register reg0, Register reg1, VirtualWord value)
{
    Word combined[2];
    VirtualWord* combinedAsUint16 = reinterpret_cast<VirtualWord*>(combined);
    *combinedAsUint16 = value;
    self->registers[EToI(reg0)] = combined[0];
    self->registers[EToI(reg1)] = combined[1];
}

scr::VirtualWord scr::CPU_getCombinedRegisterValue(CPU* self, Register reg0, Register reg1)
{
    Word combined[2];
    combined[0] = self->registers[EToI(reg0)];
    combined[1] = self->registers[EToI(reg1)];
    VirtualWord* combinedAsUint16 = reinterpret_cast<VirtualWord*>(combined);
    return *combinedAsUint16;
}

void scr::CPU_SetArithmeticFlag(CPU* self, ArithmeticFlag flag, bool value)
{
    if (value)
    {
        // set to 1
        self->registers[EToI(Register::F)] |= static_cast<Word>(flag);
        return;
    }
    // set to 0
    self->registers[EToI(Register::F)] &= ~static_cast<Word>(flag);
}

void scr::CPU_Init(CPU* self)
{
    constexpr auto size = sizeof(self->registers);
    memset(self->registers, 0, size);
}

void scr::CPU_SetRegisterValue(CPU* self, Register reg, Word value)
{
    self->registers[EToI(reg)] = value;
}

scr::Word scr::CPU_GetRegisterValue(CPU* self, Register reg)
{
    return self->registers[EToI(reg)];
}

void scr::CPU_SetVirtualRegisterValue(CPU* self, VirtualRegister reg, VirtualWord value)
{
    switch (reg)
    {
    case VirtualRegister::DE:
        CPU_setCombinedRegisters(self, Register::D, Register::E, value);
        break;
    case VirtualRegister::BC:
        CPU_setCombinedRegisters(self, Register::B, Register::C, value);
        break;
    case VirtualRegister::HL:
        CPU_setCombinedRegisters(self, Register::H, Register::L, value);
        break;
    }
}

scr::VirtualWord scr::CPU_GetVirtualRegisterValue(CPU* self, VirtualRegister reg)
{
    switch (reg)
    {
    case VirtualRegister::DE:
        return CPU_getCombinedRegisterValue(self, Register::D, Register::E);
    case VirtualRegister::BC:
        return CPU_getCombinedRegisterValue(self, Register::B, Register::C);
    case VirtualRegister::HL:
        return CPU_getCombinedRegisterValue(self, Register::H, Register::L);
    }
}

void scr::CPU_ExecuteInstruction(CPU* self, Instruction inst, Register target)
{
    switch (inst)
    {
    case Instruction::ADD:
        CPU_ExecuteADD(self, target);
        break;
    }
}
