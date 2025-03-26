#pragma once

#include <memory>

namespace scr
{
    using T = int;
    struct BinaryNode
    {
        T Value;
        std::shared_ptr<BinaryNode> Parent;
        std::shared_ptr<BinaryNode> ChildLeft;
        std::shared_ptr<BinaryNode> ChildRight;

        BinaryNode(std::shared_ptr<BinaryNode> parent, T value)
            :
            Value(value),
            Parent(parent),
            ChildLeft(nullptr),
            ChildRight(nullptr)
        {
        }
    };

    using T = int;
    class BinarySearchTree
    {
    private:
        using NodePtr = std::shared_ptr<BinaryNode>;
    public:
        BinarySearchTree(T rootValue)
            :
            m_root(std::make_shared<BinaryNode>(nullptr, rootValue))
        {
        }

        NodePtr Root()
        {
            return m_root;
        }

        void Insert(T value)
        {
            insertRecursively(m_root, value);
        }
    private:
        void insertRecursively(NodePtr node, const T& value)
        {
            if (node->Value > value)
            {
                if (node->ChildLeft == nullptr)
                {
                    node->ChildLeft = std::make_shared<BinaryNode>(node, value);
                    return;
                }
                insertRecursively(node->ChildLeft, value);
                return;
            }
            if (node->Value < value)
            {
                if (node->ChildRight == nullptr)
                {
                    node->ChildRight = std::make_shared<BinaryNode>(node, value);
                    return;
                }
                insertRecursively(node->ChildRight, value);
            }
        }

        NodePtr m_root;
    };
}