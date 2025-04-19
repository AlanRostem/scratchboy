#pragma once

#include <stdint.h>

namespace scr
{
  struct BoolByte
  {
    bool bits[8];
  };

  BoolByte NewBoolByte(uint8_t value);
  bool BoolByte_Get(BoolByte* self, uint8_t n);
}