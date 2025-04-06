#pragma once

#include <stdint.h>

namespace scr
{
    /// @brief EtoI wraps static_cast<size_t>(e) for enum classes.
    /// @tparam E enum class type (implicit)
    /// @param e enum class value
    /// @return signed integer value of the enum class value
    template<typename E>
    constexpr inline int32_t EToI(E e)
    {
        return static_cast<int32_t>(e);
    }

    using Word = uint8_t;
    using VirtualWord = uint16_t;

    /// @brief Register labels which register to read or write to in the EmulatedRegisters type.
    enum class Register
    {
        A = 0,
        B,
        C,
        D,
        E,
        F,
        H,
        L,
        Count,
    };

    /// @brief VirtualRegister labels which virtual (combined) registers to read or write to in the EmulatedRegisters type.
    enum class VirtualRegister
    {
        AF = 0,
        BC,
        DE,
        HL,
    };

    /// @brief ArithmeticFlag labels which bit flag of the F register to flip/unflip.
    enum class ArithmeticFlag
    {
        Carry = 1 << (4),
        HalfCarry = 1 << (5),
        Subtract = 1 << (6),
        Zero = 1 << (7),
    };

    /// @brief EmulatedRegisters is a static array of the Register enum class.
    using EmulatedRegisters = Word[EToI(Register::Count)];
    
    void EmulatedRegisters_Init(EmulatedRegisters self);

    void EmulatedRegisters_SetValue(EmulatedRegisters self, Register reg, Word value);
    Word EmulatedRegisters_GetValue(EmulatedRegisters self, Register reg);

    void EmulatedRegisters_SetValueVirtual(EmulatedRegisters self, VirtualRegister reg, VirtualWord value);
    VirtualWord EmulatedRegisters_GetValueVirtual(EmulatedRegisters self, VirtualRegister reg);

    void EmulatedRegisters_SetArithmeticFlag(EmulatedRegisters self, ArithmeticFlag flag, bool value);
}