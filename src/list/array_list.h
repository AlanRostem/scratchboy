#pragma once

#include <cstdint>
#include <memory>

namespace scr
{
    using T = int;
    class ArrayList
    {
    private:
        using ArrayType = T[];
        static constexpr std::size_t InitialCapacity = 5;
    public:
        ArrayList()
            :
            ArrayList(InitialCapacity)
        {

        }

        ArrayList(std::size_t reservedCapacity)
            :
            m_size(0),
            m_capacity(reservedCapacity),
            m_elements(std::make_unique<ArrayType>(reservedCapacity))
        {

        }

        void Insert(std::size_t index, T value)
        {
            expand();
            shiftRight(index);
            m_elements[index] = value;
        }

        void Append(T value)
        {
            Insert(m_size, value);
        }
    private:
        void expand()
        {
            if (m_size < m_capacity)
            {
                return;
            }
        }

        void shiftRight(std::size_t fromIndex)
        {
            if (fromIndex == m_size)
            {
                return;
            }

            for (std::size_t i = fromIndex; i < m_size; i++)
            {
                
            }
        }

        std::size_t m_capacity;
        std::size_t m_size;
        std::unique_ptr<ArrayType> m_elements;
    };
}