#pragma once

#include <stdint.h>

#include "types.h"

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

    /// @brief ArithmeticFlag labels which bit flag of the F register to flip or unflip.
    enum class ArithmeticFlag
    {
        Carry = 1 << (4),
        HalfCarry = 1 << (5),
        Subtract = 1 << (6),
        Zero = 1 << (7),
    };

    /// @brief RegisterSet is a static array of the Register enum class type until Register::PhysicalCount. This type stores
    /// the emulated state of each physical CPU register.
    using RegisterSet = Word[EToI(Operand::PhysicalRegisterCount)];
    
    void RegisterSet_Init(RegisterSet self);

    void RegisterSet_SetValue(RegisterSet self, Operand reg, Word value);
    Word RegisterSet_GetValue(RegisterSet self, Operand reg);

    void RegisterSet_SetValueVirtual(RegisterSet self, Operand reg, VirtualWord value);
    VirtualWord RegisterSet_GetValueVirtual(RegisterSet self, Operand reg);

    void RegisterSet_SetArithmeticFlag(RegisterSet self, ArithmeticFlag flag, bool value);
}