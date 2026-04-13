#pragma once

#include <assert.h>

#include "types/word.h"

namespace scr
{
    inline Word ExtractFirstNBits(Word value, uint8_t n)
    {
        return value >> (8-n);
    }

    inline Word ExtractLastNBits(Word value, uint8_t n)
    {
        constexpr Word allOnes = 0b11111111;
        Word mask = allOnes >> n;
        return mask & value;
    }

    inline Word ExtractRangedBits(Word value, uint8_t a, uint8_t b)
    {
        assert(a > b);
        
    }
}