#include "add.h"

void scr::CPU_ExecADD_A(CPU *cpu, Register target)
{
    RegisterBank *registers = CPU_GetRegisters(cpu);
    Word value = RegisterBank_GetValue(registers, target);
    Word aRegValue = RegisterBank_GetValue(registers, Register::A);
    // emulated overflow check:
    uint16_t overFlowCheckValue =
        static_cast<uint16_t>(aRegValue) +
        static_cast<uint16_t>(value);
    // add the value and allow overflow
    RegisterBank_SetValue(registers, Register::A, aRegValue + value);
    Word newAValue = RegisterBank_GetValue(registers, Register::A);
    RegisterBank_SetArithmeticFlag(registers, ArithmeticFlag::Subtract, false);
    RegisterBank_SetArithmeticFlag(registers, ArithmeticFlag::Zero, newAValue == 0);
    RegisterBank_SetArithmeticFlag(registers, ArithmeticFlag::Carry, overFlowCheckValue > 255);
    RegisterBank_SetArithmeticFlag(
        registers, ArithmeticFlag::HalfCarry,
        ((newAValue & 0xF) + (value & 0xF)) > 0xF);
}

void scr::CPU_ExecADD_A_HL(CPU *cpu)
{
    // TODO: Implement
}
