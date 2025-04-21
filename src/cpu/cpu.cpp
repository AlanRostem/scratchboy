#include <assert.h>

#include "instructions/raw.h"
#include "instructions/table.h"
#include "util/bits.h"
#include "cpu.h"

void scr::CPU_Init(CPU* cpu)
{
    cpu->programCounter = 0;
    cpu->stackPointer = 0;
    RegisterFile_Init(&cpu->registers);
    MemoryBus_Init(&cpu->bus);

    RegisterFile_SetValue(&cpu->registers, scr::Register::A, 120);
    MemoryBus_Write(&cpu->bus, 0, opcodes::ADD_A_A); // hardcoding "add A, A" instruction to first byte in memory bus for testing.
}

namespace scr
{
    void CPU_decodeAndExecute(CPU* self, Word opcode)
    {
        Instruction inst = instructions::Table[opcode];
        assert(inst != nullptr);
        (inst)(self, opcode);
    }
}

void scr::CPU_Step(CPU* self)
{
    Word instructionByte = CPU_FetchOpcode(self);
    CPU_decodeAndExecute(self, instructionByte);
    CPU_IncrementPC(self);
}

void scr::CPU_IncrementPC(CPU *self)
{
    self->programCounter++;
}

scr::Word scr::CPU_FetchOpcode(CPU *self)
{
    return MemoryBus_Read(&self->bus, self->programCounter);
}

scr::RegisterFile* scr::CPU_GetRegisters(CPU* self)
{
    return &self->registers;
}
