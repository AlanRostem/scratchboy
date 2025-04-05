#pragma once

#include <stdint.h>
#include "allocator/vtable_allocator.h"

namespace scr
{
    using T = int;

    struct BinarySearchTreeNode
    {
        T value;
        BinarySearchTreeNode *ChildLeft;
        BinarySearchTreeNode *ChildRight;
    };

    struct BinarySearchTreeSearchResult
    {
        bool IsFound;
        BinarySearchTreeNode *ParentNode;
        BinarySearchTreeNode *FoundNode;
    };

    struct BinarySearchTree
    {
        BinarySearchTreeNode *root;
        const VTableAllocator *allocator;
    };

    BinarySearchTreeNode *BinarySearchTree_NewNode(const VTableAllocator *allocator, T value)
    {
        void *nodeVoid = allocator->Alloc(sizeof(BinarySearchTreeNode));
        BinarySearchTreeNode *node = reinterpret_cast<BinarySearchTreeNode *>(nodeVoid);
        *node = BinarySearchTreeNode{
            .value = value,
            .ChildLeft = nullptr,
            .ChildRight = nullptr,
        };
        return node;
    }

    void BinarySearchTree_Init(BinarySearchTree *self, const VTableAllocator *allocator, T rootValue)
    {
        self->allocator = allocator;
        self->root = BinarySearchTree_NewNode(allocator, rootValue);
    }

    BinarySearchTree *NewBinarySearchTree(const VTableAllocator *allocator, T rootValue)
    {
        void *selfVoid = allocator->Alloc(sizeof(BinarySearchTree));
        BinarySearchTree *self = reinterpret_cast<BinarySearchTree *>(selfVoid);
        BinarySearchTree_Init(self, allocator, rootValue);
        return self;
    }

    BinarySearchTreeSearchResult BinarySearchTree_findRecursively(BinarySearchTree *tree, BinarySearchTreeNode *parent, BinarySearchTreeNode *node, T value)
    {
        if (node == nullptr)
        {
            return BinarySearchTreeSearchResult{
                .IsFound = false,
                .ParentNode = nullptr,
                .FoundNode = nullptr};
        }

        if (node->value == value)
        {
            return BinarySearchTreeSearchResult{
                .IsFound = true,
                .ParentNode = parent,
                .FoundNode = node};
        }

        if (value < node->value)
        {
            return BinarySearchTree_findRecursively(tree, node, node->ChildLeft, value);
        }

        return BinarySearchTree_findRecursively(tree, node, node->ChildRight, value);
    }

    BinarySearchTreeSearchResult BinarySearchTree_Find(BinarySearchTree *tree, T value)
    {
        return BinarySearchTree_findRecursively(tree, nullptr, tree->root, value);
    }

    void BinarySearchTree_insertRecursively(BinarySearchTree *tree, BinarySearchTreeNode *node, T value)
    {
        if (value < node->value)
        {
            if (node->ChildLeft == nullptr)
            {
                node->ChildLeft = BinarySearchTree_NewNode(tree->allocator, value);
                return;
            }
            BinarySearchTree_insertRecursively(tree, node->ChildLeft, value);
        }
        else if (value > node->value)
        {
            if (node->ChildRight == nullptr)
            {
                node->ChildRight = BinarySearchTree_NewNode(tree->allocator, value);
                return;
            }
            BinarySearchTree_insertRecursively(tree, node->ChildRight, value);
        }
    }

    void BinarySearchTree_Insert(BinarySearchTree *tree, T value)
    {
        if (tree->root == nullptr)
        {
            tree->root = BinarySearchTree_NewNode(tree->allocator, value);
            return;
        }
        BinarySearchTree_insertRecursively(tree, tree->root, value);
    }

    BinarySearchTreeNode *BinarySearchTree_Root(BinarySearchTree *tree)
    {
        return tree->root;
    }

    void BinarySearchTree_Delete(BinarySearchTree *tree, T value)
    {
        BinarySearchTreeSearchResult result = BinarySearchTree_Find(tree, value);
        if (!result.IsFound)
        {
            return;
        }

        // TODO: Implement deletion logic
    }
}
