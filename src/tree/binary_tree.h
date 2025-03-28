#pragma once

#include <memory>

namespace scr
{
    using T = int;
    class BinarySearchTree
    {
    public:
        struct Node;
        using NodePtr = std::shared_ptr<Node>;
        
        struct Node
        {
            T Value;
            NodePtr Parent;
            NodePtr ChildLeft;
            NodePtr ChildRight;
    
            Node(const NodePtr& parent, T value)
                :
                Value(value),
                Parent(parent),
                ChildLeft(nullptr),
                ChildRight(nullptr)
            {
            }
        };
    
        struct SearchResult
        {
            bool IsFound;
            NodePtr Predecessor;
            NodePtr Successor;

            SearchResult()
                :
                IsFound(false),
                Predecessor(nullptr),
                Successor(nullptr)
            {
            }
        };
    public:
        BinarySearchTree(T rootValue)
            :
            m_root(std::make_shared<Node>(nullptr, rootValue))
        {
        }

        NodePtr Root()
        {
            return m_root;
        }

        void Insert(const T& value)
        {
            insertRecursively(m_root, value);
        }

        SearchResult Find(const T& value)
        {
            return findRecursively(m_root, value);
        }
    private:
    SearchResult findRecursively(const NodePtr& node, const T& value)
        {
            if (node->Value == value)
            {
                SearchResult result;
                result.IsFound = true;
                result.Predecessor = node->Parent;
                // TODO: implement successor
                return result;
            }

            if (node->Value > value)
            {
                if (node->ChildLeft != nullptr)
                {
                    return findRecursively(node->ChildLeft, value);
                }
                return SearchResult();
            }

            if (node->Value < value)
            {
                if (node->ChildRight != nullptr)
                {
                    return findRecursively(node->ChildRight, value);
                }
                return SearchResult();
            }
            return SearchResult();
        }

        void insertRecursively(const NodePtr& node, const T& value)
        {
            if (node->Value > value)
            {
                if (node->ChildLeft == nullptr)
                {
                    node->ChildLeft = std::make_shared<Node>(node, value);
                    node->ChildLeft->Parent = node;
                    return;
                }
                insertRecursively(node->ChildLeft, value);
                return;
            }
            if (node->Value < value)
            {
                if (node->ChildRight == nullptr)
                {
                    node->ChildRight = std::make_shared<Node>(node, value);
                    node->ChildRight->Parent = node;
                    return;
                }
                insertRecursively(node->ChildRight, value);
            }
        }

        NodePtr m_root;
    };
}