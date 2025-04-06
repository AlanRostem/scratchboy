#include <string.h>

#include "cpu.h"


namespace scr
{
    void CPU_setCarryFlag(CPU* self, ArithmeticFlag flag, bool value)
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

    void CPU_setCombinedRegisters(CPU* self, Register reg0, Register reg1, VirtualWord value)
    {
        Word combined[2];
        VirtualWord* combinedAsUint16 = reinterpret_cast<VirtualWord*>(combined);
        *combinedAsUint16 = value;
        self->registers[EToI(reg0)] = combined[0];
        self->registers[EToI(reg1)] = combined[1];
    }

    VirtualWord CPU_getCombinedRegisterValue(CPU* self, Register reg0, Register reg1)
    {
        Word combined[2];
        combined[0] = self->registers[EToI(reg0)];
        combined[1] = self->registers[EToI(reg1)];
        VirtualWord* combinedAsUint16 = reinterpret_cast<VirtualWord*>(combined);
        return *combinedAsUint16;
    }
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
    case Instruction::Add:
    {
        Word value = self->registers[EToI(target)];
        Word* aReg = &self->registers[EToI(Register::A)];
        // emulated overflow check:
        uint16_t overFlowCheckValue = static_cast<uint16_t>(*aReg) + static_cast<uint16_t>(value);
        *aReg += value; // add the value and allow overflow

        CPU_setCarryFlag(self, ArithmeticFlag::Subtract, true);
        if (*aReg == 0)
        {
            CPU_setCarryFlag(self, ArithmeticFlag::Zero, true);
        }
        else if (overFlowCheckValue > 255) // overflow check
        {
            CPU_setCarryFlag(self, ArithmeticFlag::Carry, true);
        }
        // TODO: check half-carry
    }
    break;
    }
}
