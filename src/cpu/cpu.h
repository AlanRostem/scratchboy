#pragma once

#include <stdint.h>

namespace scr
{
    enum class Register
    {
        A,
        B,
        C,
        D,
        E,
        F,
        H,
        L,
        Count,
    };

    enum class ArithmeticFlag
    {
        Carry = 1 << (4),
        HalfCarry = 1 << (5),
        Subtract = 1 << (6),
        Zero = 1 << (7),
    };

    enum class Instruction
    {
        Add,
    };

    struct CPU
    {
        uint8_t registers[static_cast<size_t>(Register::Count)];
    };

    void CPU_Init(CPU *self);

    void CPU_SetRegister(CPU *self, Register reg, uint8_t value);
    uint8_t CPU_GetRegisterValue(CPU *self, Register reg);

    void CPU_SetCombinedRegisters(CPU *self, Register reg0, Register reg1, uint16_t value);
    uint16_t CPU_GetCombinedRegisterValue(CPU *self, Register reg0, Register reg1);

    void CPU_ExecuteInstruction(CPU *self, Instruction inst, Register target);
}