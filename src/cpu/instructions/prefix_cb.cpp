#include <assert.h>

#include "util/bits.h"
#include "raw.h"
#include "instruction.h"
#include "cpu/cpu.h"

namespace scr
{
    enum class CBPrefixedRotationInstructions
    {
        RLC = 0,
        RRC,
        RL,
        RR,
        SLA,
        SRA,
        SWAP,
        SRL,
    };

    enum class CBPrefixedIndexInstructions
    {
        BIT = 1,
        RES,
        SET,
    };
}

void scr::instructions::PrefixCB(CPU *cpu, Word opcode)
{
    assert(opcode == opcodes::PrefixCB);
    CPU_IncrementPC(cpu);
    Word cbPrefixedOpcode = CPU_FetchOpcode(cpu);

    Word first2Bits = ExtractFirstNBits(cbPrefixedOpcode, 2);
    if (first2Bits > 0)
    {
        Word mnemonic = ExtractFirstNBits(cbPrefixedOpcode, 2);
        switch (IToE<CBPrefixedIndexInstructions>(mnemonic))
        {
        case CBPrefixedIndexInstructions::BIT:
            break;
        case CBPrefixedIndexInstructions::RES:
            break;
        case CBPrefixedIndexInstructions::SET:
            break;
        }
        return;
    }
    Word mnemonic = ExtractFirstNBits(cbPrefixedOpcode, 5);
    switch (IToE<CBPrefixedRotationInstructions>(mnemonic))
    {
    case CBPrefixedRotationInstructions::RLC:
        // TODO: Handle RLC
        break;
    case CBPrefixedRotationInstructions::RRC:
        // TODO: Handle RRC
        break;
    case CBPrefixedRotationInstructions::RL:
        // TODO: Handle RL
        break;
    case CBPrefixedRotationInstructions::RR:
        // TODO: Handle RR
        break;
    case CBPrefixedRotationInstructions::SLA:
        // TODO: Handle SLA
        break;
    case CBPrefixedRotationInstructions::SRA:
        // TODO: Handle SRA
        break;
    case CBPrefixedRotationInstructions::SWAP:
        // TODO: Handle SWAP
        break;
    case CBPrefixedRotationInstructions::SRL:
        // TODO: Handle SRL
        break;
    default:
        // TODO: Handle unknown instruction
        break;
    }
}
