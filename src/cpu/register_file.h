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

    // see: https://gbdev.io/pandocs/CPU_Instruction_Set.html
    enum class R8Operand
    {
        B = 0,
        C,
        D,
        E,
        H,
        L,
        HL,
        A
    };

    /// R8ToRegister converts the integer value that denotes a register, as found in an opcode, 
    /// to the Register enum value which can be use in the RegisterFile type to identify an emulated CPU register.
    /// Note that in instructions like "add A, (HL)" (where r8=6), this function returns Register::None.
    inline Register R8ToRegister(Word r8)
    {
        
        Register target = Register::None;
        R8Operand operand = IToE<R8Operand>(r8);
        switch (operand)
        {
        case R8Operand::B:
            target = Register::B;
            break;
        case R8Operand::C:
            target = Register::C;
            break;
        case R8Operand::D:
            target = Register::D;
            break;
        case R8Operand::E:
            target = Register::E;
            break;
        case R8Operand::H:
            target = Register::H;
            break;
        case R8Operand::L:
            target = Register::L;
            break;
        case R8Operand::HL:
            // this is the [HL] virtual register and is ignored
            break;
        case R8Operand::A:
            target = Register::A;
            break;
        }
        return target;
    }

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