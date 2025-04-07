#pragma once

#include <stdint.h>

#include "types.h"

namespace scr
{
    /// @brief EmuArithmeticFlag labels which emulated bit flag of the F register to flip or unflip.
    enum class EmuArithmeticFlag
    {
        Carry = 1 << (4),
        HalfCarry = 1 << (5),
        Subtract = 1 << (6),
        Zero = 1 << (7),
    };

    /// @brief EmuRegister labels which emulated register to read or write to in the EmuEmuRegisterSet type.
    enum class EmuRegister
    {
        None = -1,
        A, F,
        B, C,
        D, E,
        H, L,
        // Indicates the total number of emulated registers. 
        Count,
    };

    enum class EmuVirtualRegister
    {
        AF,
        BC,
        DE,
        HL,
    };
    

    /// @brief EmuRegisterSet is a static array of the EmuRegister enum class type until Register::Count. This type stores
    /// the emulated state of each physical CPU register.
    using EmuRegisterSet = Word[EToI(EmuRegister::Count)];
    
    void EmuRegisterSet_Init(EmuRegisterSet self);

    void EmuRegisterSet_SetValue(EmuRegisterSet self, EmuRegister reg, Word value);
    Word EmuRegisterSet_GetValue(EmuRegisterSet self, EmuRegister reg);

    void EmuRegisterSet_SetValueVirtual(EmuRegisterSet self, EmuVirtualRegister reg, VirtualWord value);
    VirtualWord EmuRegisterSet_GetValueVirtual(EmuRegisterSet self, EmuVirtualRegister reg);

    void EmuRegisterSet_SetArithmeticFlag(EmuRegisterSet self, EmuArithmeticFlag flag, bool value);
}