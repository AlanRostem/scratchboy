#pragma once

#include <stdint.h>

namespace scr
{
  struct BoolByte
  {
    bool Bits[8];
  };
  
  BoolByte NewBoolByte(uint8_t value);
  uint8_t BoolByte_Get(BoolByte* self, uint8_t n);
  
  enum class Operand
  {
    R8,
    R16,
    
  };
  
  struct Opcode
  {
    int Block;
    int Command;
    
  };
}