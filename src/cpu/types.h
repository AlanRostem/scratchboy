#pragma once
#include <stdint.h>

namespace scr
{
    /// @brief EtoI wraps static_cast<size_t>(e) for enum classes.
    /// @tparam E enum class type (implicit)
    /// @param e enum class value
    /// @return signed integer value of the enum class value
    template<typename E>
    constexpr inline int32_t EToI(E e)
    {
        return static_cast<int32_t>(e);
    }

    using Word = uint8_t;
    using Opcode = Word;
    using VirtualWord = uint16_t;
}