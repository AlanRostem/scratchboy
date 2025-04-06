#include <mem.h>

#include "cpu.h"

void scr::CPU_Init(CPU *self)
{
    constexpr auto size = sizeof(self->registers);
    memset(self->registers, 0, size);
}

void scr::CPU_SetRegister(CPU *self, Register reg, uint8_t value)
{
    self->registers[static_cast<size_t>(reg)] = value;
}

uint8_t scr::CPU_GetRegisterValue(CPU *self, Register reg)
{
    return self->registers[static_cast<size_t>(reg)];
}

void scr::CPU_SetCombinedRegisters(CPU* self, Register reg0, Register reg1, uint16_t value)
{
    uint8_t combined[2];
    uint16_t* combinedAsUint16 = reinterpret_cast<uint16_t*>(combined);
    *combinedAsUint16 = value;
    self->registers[static_cast<size_t>(reg0)] = combined[0];
    self->registers[static_cast<size_t>(reg1)] = combined[1];
}

uint16_t scr::CPU_GetCombinedRegisterValue(CPU *self, Register reg0, Register reg1)
{
    uint8_t combined[2];
    combined[0] = self->registers[static_cast<size_t>(reg0)];
    combined[1] = self->registers[static_cast<size_t>(reg1)];
    uint16_t* combinedAsUint16 = reinterpret_cast<uint16_t*>(combined);
    return *combinedAsUint16;
}

namespace scr
{   
    void CPU_SetCarryFlag(CPU* self, ArithmeticFlag flag, bool value)
    {
        if (value)
        {
            // set to 1
            self->registers[static_cast<size_t>(Register::F)] |= static_cast<uint8_t>(flag);
            return;
        }
        // set to 0
        self->registers[static_cast<size_t>(Register::F)] &= ~static_cast<uint8_t>(flag);
    }
}

void scr::CPU_ExecuteInstruction(CPU *self, Instruction inst, Register target)
{
    switch(inst)
    {
        case Instruction::Add:
        {
            uint8_t value = self->registers[static_cast<size_t>(target)];
            uint8_t* aReg = &self->registers[static_cast<size_t>(Register::A)];
            uint16_t overFlowCheckValue = *aReg + value;
            *aReg += value;

            CPU_SetCarryFlag(self, ArithmeticFlag::Subtract, true);
            if (*aReg == 0)
            {
                CPU_SetCarryFlag(self, ArithmeticFlag::Zero, true);
            }
            else if (overFlowCheckValue > 255) // overflow
            {
                CPU_SetCarryFlag(self, ArithmeticFlag::Carry, true);
            } 
            // TODO: check half-carry
        }
        break;
    }
}
