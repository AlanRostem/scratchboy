#pragma once

#include <stdint.h>

namespace scr
{
    using Word = uint8_t;
    using VirtualWord = uint16_t;

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

    enum class VirtualRegister
    {
        BC,
        DE,
        HL,
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
        ADD,
    };

    using InstructionMethod = void(*)(CPU*, Register);

    /// @brief EtoI wraps static_cast<size_t>(e) for enum classes.
    /// @tparam E enum class type (implicit)
    /// @param e enum class value
    /// @return signed integer value of the enum class value
    template<typename E>
    constexpr inline int32_t EToI(E e)
    {
        return static_cast<int32_t>(e);
    }

    struct CPU
    {
        Word registers[EToI(Register::Count)];
    };

    void CPU_Init(CPU* self);

    void CPU_SetRegisterValue(CPU* self, Register reg, Word value);
    Word CPU_GetRegisterValue(CPU* self, Register reg);

    void CPU_SetVirtualRegisterValue(CPU* self, VirtualRegister reg, VirtualWord value);
    VirtualWord CPU_GetVirtualRegisterValue(CPU* self, VirtualRegister reg);

    void CPU_SetArithmeticFlag(CPU* self, ArithmeticFlag flag, bool value);

    void CPU_ExecuteInstruction(CPU* self, Instruction inst, Register target);
}