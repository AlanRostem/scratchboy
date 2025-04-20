#include "add.h"
#include "util/bits.h"

void scr::instructions::add_A_r8(CPU* cpu, Word opcode)
{
    Word r8 = ExtractLastNBits(opcode, 3);
    Register target = RegisterFromR8(r8);
    

    RegisterFile* registers = CPU_GetRegisters(cpu);
    Word targetValue = RegisterFile_GetValue(registers, target);
    Word aRegValue = RegisterFile_GetValue(registers, Register::A);
    // emulated overflow check:
    uint16_t overFlowCheckValue =
        static_cast<uint16_t>(aRegValue) +
        static_cast<uint16_t>(targetValue);
    // add the value and allow overflow
    RegisterFile_SetValue(registers, Register::A, aRegValue + targetValue);
    Word newAValue = RegisterFile_GetValue(registers, Register::A);
    RegisterFile_SetArithmeticFlag(registers, ArithmeticFlag::Subtract, false);
    RegisterFile_SetArithmeticFlag(registers, ArithmeticFlag::Zero, newAValue == 0);
    RegisterFile_SetArithmeticFlag(registers, ArithmeticFlag::Carry, overFlowCheckValue > 255);
    RegisterFile_SetArithmeticFlag(
        registers, ArithmeticFlag::HalfCarry,
        ((newAValue & 0xF) + (targetValue & 0xF)) > 0xF);
}