#pragma once

#include <stdint.h>

#include "types/word.h"

namespace scr
{
    /// @brief ArithmeticFlag labels which emulated bit flag of the F register to flip or unflip.
    enum class ArithmeticFlag
    {
        Carry = 1 << (4),
        HalfCarry = 1 << (5),
        Subtract = 1 << (6),
        Zero = 1 << (7),
    };

    /// @brief Register labels which emulated register to read or write to in the RegisterBank type.
    enum class Register
    {
        None = -1,
        A, F,
        B, C,
        D, E,
        H, L,
        // Indicates the total number of emulated registers. 
        Count,
    };

    enum class VirtualRegister
    {
        AF,
        BC,
        DE,
        HL,
    };
    

    /// @brief RegisterBank stores the emulated state of each physical CPU register.
    struct RegisterBank
    {
        Word rawA;
        Word rawF;
        Word rawB;
        Word rawC;
        Word rawD;
        Word rawE;
        Word rawH;
        Word rawL;
    };
    
    void RegisterBank_Init(RegisterBank* self);

    void RegisterBank_SetValue(RegisterBank* self, Register reg, Word value);
    Word RegisterBank_GetValue(RegisterBank* self, Register reg);

    void RegisterBank_SetValueVirtual(RegisterBank* self, VirtualRegister reg, VirtualWord value);
    VirtualWord RegisterBank_GetValueVirtual(RegisterBank* self, VirtualRegister reg);

    void RegisterBank_SetArithmeticFlag(RegisterBank* self, ArithmeticFlag flag, bool value);
}