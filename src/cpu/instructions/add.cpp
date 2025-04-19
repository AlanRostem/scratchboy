#include "add.h"
#include "util/bits.h"


void scr::CPU_ExecADD_A(CPU* cpu, Word opcode)
{
    Word r8 = ExtractLastNBits(opcode, 3);
    Register target;
    // see: https://gbdev.io/pandocs/CPU_Instruction_Set.html
    switch (r8)
    {
    case 0:
        target = Register::B;
        break;
    case 1:
        target = Register::C;
        break;
    case 2:
        target = Register::D;
        break;
    case 3:
        target = Register::E;
        break;
    case 4:
        target = Register::H;
        break;
    case 5:
        target = Register::L;
        break;
    case 6:
        // TODO: implement "add A, HL"
        return;
    case 7:
        target = Register::A;
        break;
    }


    RegisterFile* registers = CPU_GetRegisters(cpu);
    Word value = RegisterFile_GetValue(registers, target);
    Word aRegValue = RegisterFile_GetValue(registers, Register::A);
    // emulated overflow check:
    uint16_t overFlowCheckValue =
        static_cast<uint16_t>(aRegValue) +
        static_cast<uint16_t>(value);
    // add the value and allow overflow
    RegisterFile_SetValue(registers, Register::A, aRegValue + value);
    Word newAValue = RegisterFile_GetValue(registers, Register::A);
    RegisterFile_SetArithmeticFlag(registers, ArithmeticFlag::Subtract, false);
    RegisterFile_SetArithmeticFlag(registers, ArithmeticFlag::Zero, newAValue == 0);
    RegisterFile_SetArithmeticFlag(registers, ArithmeticFlag::Carry, overFlowCheckValue > 255);
    RegisterFile_SetArithmeticFlag(
        registers, ArithmeticFlag::HalfCarry,
        ((newAValue & 0xF) + (value & 0xF)) > 0xF);
}