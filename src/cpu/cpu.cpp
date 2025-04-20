#include <assert.h>

#include "instructions/table.h"
#include "instructions/opcodes.h"
#include "util/bits.h"
#include "cpu.h"

void scr::CPU_Init(CPU* cpu)
{
    cpu->programCounter = 0;
    cpu->stackPointer = 0;
    RegisterFile_Init(&cpu->registers);
    MemoryBus_Init(&cpu->bus);

    RegisterFile_SetValue(&cpu->registers, scr::Register::A, 120);
    MemoryBus_Write(&cpu->bus, 0, 0x87); // hardcoding "add A, A" instruction to first byte in memory bus.
}

namespace scr
{
    bool CPU_decodeAndExecuteStandard(CPU* self, Word opcode)
    {
        Instruction inst = instructions::Table[opcode];
        if (inst == nullptr)
        {
            return false;
        }
        (inst)(self, opcode);
        return true;
    }

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

    void CPU_decodeAndExecutePrefixedCB(CPU* self, Word cbPrefixedOpcode)
    {
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
}

void scr::CPU_Step(CPU* self)
{
    Word instructionByte = MemoryBus_Read(&self->bus, self->programCounter);
    if (instructionByte == EToI(Opcodes::PrefixCB))
    {
        instructionByte = MemoryBus_Read(&self->bus, self->programCounter + 1);
        CPU_decodeAndExecutePrefixedCB(self, instructionByte);
        self->programCounter++; // TODO: determine if this is correct behavior
        return;
    }

    if (CPU_decodeAndExecuteStandard(self, instructionByte))
    {
        self->programCounter++;
        return;
    }
}

scr::RegisterFile* scr::CPU_GetRegisters(CPU* self)
{
    return &self->registers;
}
