#include "cpu.h"

void scr::CPU_Init(CPU *cpu)
{
    RegisterBank_Init(&cpu->registers);
}

scr::RegisterBank *scr::CPU_GetRegisters(CPU *self)
{
    return &self->registers;
}
