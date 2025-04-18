#include "opcode.h"

scr::BoolByte NewBoolByte(uint8_t value)
{
  scr::BoolByte byte;
  for (int i = 0; i < 8; i++)
  {
    byte.Bits[8-i] =  static_cast<bool>((1 << i) & value);
  } 
  return byte;
}

bool scr::BoolByte_Get(BoolByte* self, uint8_t n)
{
  if (n >= 8)
  {
    return false;
  }
  return self->Bits[n];
}