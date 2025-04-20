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
    instructions::Table[opcode - 0x80](self, opcode);
}
