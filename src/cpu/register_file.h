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

    /// @brief Register labels which emulated register to read or write to in the RegisterFile type.
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
    

    /// @brief RegisterFile stores the emulated state of each physical CPU register.
    struct RegisterFile
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
    
    void RegisterFile_Init(RegisterFile* self);

    void RegisterFile_SetValue(RegisterFile* self, Register reg, Word value);
    Word RegisterFile_GetValue(RegisterFile* self, Register reg);

    void RegisterFile_SetValueVirtual(RegisterFile* self, VirtualRegister reg, DoubleWord value);
    DoubleWord RegisterFile_GetValueVirtual(RegisterFile* self, VirtualRegister reg);

    void RegisterFile_SetArithmeticFlag(RegisterFile* self, ArithmeticFlag flag, bool value);
}