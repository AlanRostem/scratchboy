#pragma once

namespace scr
{
    enum class InstructionType
    {
        Misc,
        Call,
        Arith8,
        Arith16,
        Access8, // load, store, or move
        Access16, // load, store, or move
        Shift,
    };

    enum class Operand
    {

    };

    enum class ArithOp
    {
        Add,
        Sub,
        And,
        Or,
    };

    struct InstructionArith8
    {
        ArithOp Op;
        
    };

    struct Instruction
    {
        InstructionType Type;
        union
        {
            InstructionArith8 Arith8;
        } UnionSelf;
    };
}