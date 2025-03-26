#pragma once

#include <cstdint>
#include <memory>
#include <stdexcept>

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
            m_elements(new T[reservedCapacity])
        {

        }

        void Insert(std::size_t index, T value)
        {
            // TODO: optimize by expanding and shifting to the right in the same operation
            expand();
            shiftRight(index);
            m_elements[index] = value;
        }

        void Append(T value)
        {
            Insert(m_size, value);
        }

        T Get(std::size_t index) const
        {
            if (index >= m_size)
            {
                throw std::invalid_argument("index exceeds array bounds");
            }
            return m_elements[index];
        }

        void Set(std::size_t index, T value)
        {
            if (index >= m_size)
            {
                throw std::invalid_argument("index exceeds array bounds");
            }
            m_elements[index] = value;
        }

        void Remove(std::size_t index)
        {
            if (index >= m_size)
            {
                throw std::invalid_argument("index exceeds array bounds");
            }
            shiftLeftOverwrite(index);
            shrink();
        }

        std::size_t Size() const
        {
            return m_size;
        }
    private:
        void expand()
        {
            m_size++;
            if (m_size < m_capacity)
            {
                return;
            }
            m_capacity = m_capacity * 3 / 2;
            auto newElements = std::shared_ptr<T[]>(new T[m_capacity]);
            for (std::size_t i = 0; i < m_size; i++)
            {
                newElements[i] = m_elements[i];
            }
            m_elements = newElements;
        }

        void shrink()
        {
            // no need to deallocate excess space
            m_size--;
        }

        void shiftRight(std::size_t fromIndex)
        {
            if (fromIndex == m_size || m_size < 2) // TODO: handle when the size is 2 
            {
                return;
            }
            for (std::size_t i = m_size; i >= fromIndex; i--)
            {
                m_elements[i] = m_elements[i-1];
            }
        }

        void shiftLeftOverwrite(std::size_t fromIndex)
        {
            if (fromIndex == 0)
            {
                return;
            }

            for (std::size_t i = fromIndex; i < m_size; i++)
            {
                m_elements[i] = m_elements[i+1];
            }
        }

        std::size_t m_size;
        std::size_t m_capacity;
        std::shared_ptr<ArrayType> m_elements;
    };
}