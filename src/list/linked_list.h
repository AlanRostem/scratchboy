#pragma once

#include <memory>

namespace scr
{
    using T = int;
    class LinkedList
    {
    private:
        struct Node;
        using NodePtr = std::shared_ptr<Node>;

        struct Node
        {
            T Value;
            NodePtr Next;
            Node(T value)
                :
                Value(value),
                Next(nullptr)
            {

            }
        };
    public:
        LinkedList()
            : m_head(nullptr)
        {
        }

        void Append(const T& value)
        {
            if (m_head == nullptr)
            {
                m_head = std::make_shared<Node>(value);
                return;
            }

            auto current = m_head;
            while (current->Next != nullptr)
            {
                current = current->Next;
            }
            current->Next = std::make_shared<Node>(value);
        }

        
    private:
        NodePtr m_head;
    };
}