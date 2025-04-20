#include <assert.h>

#include "instructions/table.h"
#include "cpu.h"

void scr::CPU_Init(CPU *cpu)
{
    RegisterFile_Init(&cpu->registers);
}

scr::RegisterFile *scr::CPU_GetRegisters(CPU *self)
{
    return &self->registers;
}

void scr::CPU_DecodeAndExecute(CPU* self, Word opcode)
{
    Instruction inst = instructions::Table[opcode];
    assert(inst != nullptr);
    (inst)(self, opcode);
}
