#include <stdint.h>

namespace scr
{
    using Word = uint8_t;
    using Opcode = Word;
    using VirtualWord = uint16_t;

    /// @brief Register labels which register to read or write to in the RegisterSet type, including virtual ones.
    enum class Operand
    {
        None = -1,
        A, F,
        B, C,
        D, E,
        H, L,
        // Indicates the total number of physical registers. Any register number above this counts as virtual ones. 
        PhysicalRegisterCount, 
        AF,
        BC,
        DE,
        HL,
        // Indicates total number of registers, virtual or physical. Any number above this counts as immediate data operands.
        RegisterCount,
        
    };

    enum class Mnemonic {
        // Load / Store
        LD,
        LDI,
        LDD,
        PUSH,
        POP,

        // Arithmetic
        ADD,
        ADC,
        SUB,
        SBC,
        INC,
        DEC,
        AND,
        OR,
        XOR,
        CP,
        CPL,
        DAA,

        // Control flow
        JP,
        JR,
        CALL,
        RET,
        RETI,
        RST,

        // Rotate / Shift / Bit Ops (CB-prefixed)
        RLCA,
        RLA,
        RRCA,
        RRA,
        RLC,
        RL,
        RRC,
        RR,
        SLA,
        SRA,
        SWAP,
        SRL,
        BIT,
        SET,
        RES,

        // Flags / Special
        NOP,
        HALT,
        STOP,
        DI,
        EI,
        SCF,
        CCF,

        // Meta
        INVALID,  // Used for undefined opcodes
    };
}