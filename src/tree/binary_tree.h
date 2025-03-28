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
            NodePtr ChildLeft;
            NodePtr ChildRight;

            Node(T value)
                : Value(value),
                  ChildLeft(nullptr),
                  ChildRight(nullptr)
            {
            }
        };

        struct SearchResult
        {
            bool IsFound;
            NodePtr ParentNode;
            NodePtr FoundNode;
            SearchResult()
                : IsFound(false),
                  ParentNode(nullptr),
                  FoundNode(nullptr)
            {
            }
        };

    public:
        BinarySearchTree(T rootValue)
            : m_root(std::make_shared<Node>(rootValue))
        {
        }

        NodePtr Root()
        {
            return m_root;
        }

        void Insert(const T &value)
        {
            insertRecursively(m_root, value);
        }

        void Delete(const T &value)
        {
            auto result = Find(value);
            if (!result.IsFound)
            {
                return;
            }
            // TODO: Implement
        }

        SearchResult Find(const T &value)
        {
            return findRecursively(nullptr, m_root, value);
        }

    private:
        SearchResult findRecursively(const NodePtr &parent, const NodePtr &node, const T &value)
        {
            if (node->Value == value)
            {
                SearchResult result;
                result.IsFound = true;
                result.ParentNode = parent;
                result.FoundNode = node;
                return result;
            }
            if (node->Value > value)
            {
                if (node->ChildLeft != nullptr)
                {
                    return findRecursively(node, node->ChildLeft, value);
                }
                return SearchResult();
            }
            if (node->Value < value)
            {
                if (node->ChildRight != nullptr)
                {
                    return findRecursively(node, node->ChildRight, value);
                }
                return SearchResult();
            }
            return SearchResult();
        }

        void insertRecursively(const NodePtr &node, const T &value)
        {
            if (node->Value > value)
            {
                if (node->ChildLeft == nullptr)
                {
                    node->ChildLeft = std::make_shared<Node>(value);
                    return;
                }
                insertRecursively(node->ChildLeft, value);
                return;
            }
            if (node->Value < value)
            {
                if (node->ChildRight == nullptr)
                {
                    node->ChildRight = std::make_shared<Node>(value);
                    return;
                }
                insertRecursively(node->ChildRight, value);
            }
        }

        NodePtr m_root;
    };
}