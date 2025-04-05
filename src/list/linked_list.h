#pragma once

#include "allocator/vtable_allocator.h"

namespace scr
{
    
    template<typename T>
    struct LinkedListNode
    {
        T Value;
        LinkedListNode* Next;
    };

    template<typename T>
    LinkedListNode<T>* NewLinkedListNode(VTableAllocator* allocator, T value)
    {
        void* valueVoid = allocator->Alloc(sizeof(LinkedListNode<T>));
        LinkedListNode<T>* nodePtr = reinterpret_cast<LinkedListNode<T>*>(valueVoid);
        *nodePtr = LinkedListNode<T>{
            .Value = value,
            .Next = nullptr,
        };
        return nodePtr;
    }
    
    template<typename T>
    struct LinkedList
    {
        LinkedListNode<T>* head;
        VTableAllocator* allocator;
    };
    
    using T = int;

    void LinkedList_Init(LinkedList<T>* self, VTableAllocator* allocator)
    {
        self->head = nullptr;
        self->allocator = allocator;
    }

    void LinkedList_Append(LinkedList<T>* self, T value)
    {
        if (self->head == nullptr)
        {
            self->head = NewLinkedListNode(self->allocator, value);
            return;
        }
        LinkedListNode<T>* current = self->head;
        while(current->Next != nullptr)
        {
            current = current->Next;
        }

        current->Next = NewLinkedListNode(self->allocator, value);
    }

    // TODO: Implement lookup and destruction.
}