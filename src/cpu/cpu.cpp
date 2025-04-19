#include "cpu.h"

void scr::CPU_Init(CPU *cpu)
{
    RegisterFile_Init(&cpu->registers);
}

scr::RegisterFile *scr::CPU_GetRegisters(CPU *self)
{
    return &self->registers;
}
