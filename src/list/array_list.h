#pragma once

#include <stdint.h>

#include "allocator/vtable_allocator.h"

namespace scr
{
    constexpr size_t ArrayList_InitialCapacity = 5;
    template <typename T>
    struct ArrayList
    {
        T *Data;
        size_t Size;
        size_t Capacity;
        const VTableAllocator *Allocator;
    };
    // using T = int;

    template <typename T>
    void ArrayList_Ctor(ArrayList<T> *self, const VTableAllocator *allocator)
    {
        self->Size = 0;
        self->Capacity = ArrayList_InitialCapacity;
        self->Allocator = allocator;
        void *data = allocator->Alloc(sizeof(T) * self->Capacity);
        self->Data = reinterpret_cast<T *>(data);
    }

    template <typename T>
    ArrayList<T> *NewArrayList(VTableAllocator *allocator)
    {
        void *selfVoid = allocator->Alloc(sizeof(ArrayList<T>));
        ArrayList<T> *self = reinterpret_cast<ArrayList<T> *>(selfVoid);
        ArrayList_Ctor(self, allocator);
        return self;
    }

    template <typename T>
    void ArrayList_Dtor(ArrayList<T> *self)
    {
        self->Allocator->Free(self->Data);
        self->Data = nullptr;
        self->Size = 0;
        self->Capacity = 0;
        self->Allocator = nullptr;
    }

    template <typename T>
    void ArrayList_expand(ArrayList<T> *self)
    {
        self->Size++;
        if (self->Size < self->Capacity)
        {
            return;
        }
        self->Capacity = self->Capacity * 3 / 2;
        void *newElementsVoid = self->Allocator->ReAlloc(self->Data, self->Capacity);
        T *newElements = reinterpret_cast<T *>(newElementsVoid);
        self->Data = newElements;
    }

    template <typename T>
    void ArrayList_shiftRight(ArrayList<T> *self, std::size_t fromIndex)
    {
        if (fromIndex == self->Size || self->Size < 2) // TODO: handle when the size is 2
        {
            return;
        }
        for (std::size_t i = self->Size; i >= fromIndex; i--)
        {
            self->Data[i] = self->Data[i - 1];
        }
    }

    template <typename T>
    void ArrayList_shiftLeftOverwrite(ArrayList<T> *self, std::size_t fromIndex)
    {
        if (fromIndex == 0)
        {
            return;
        }

        for (std::size_t i = fromIndex; i < self->Size; i++)
        {
            self->Data[i] = self->Data[i + 1];
        }
    }

    template <typename T>
    void ArrayList_shrink(ArrayList<T> *self)
    {
        // no need to deallocate excess space
        self->Size--;
    }

    template <typename T>
    void ArrayList_Insert(ArrayList<T> *self, size_t index, T value)
    {
        ArrayList_expand(self);
        ArrayList_shiftRight(self, index);
        self->Data[index] = value;
    }

    template <typename T>
    void ArrayList_Append(ArrayList<T> *self, T value)
    {
        ArrayList_Insert(self, self->Size, value);
    }

    template <typename T>
    T *ArrayList_Get(ArrayList<T> *self, size_t index)
    {
        if (index >= self->Size)
        {
            return nullptr;
        }
        return self->Data + index;
    }

    template <typename T>
    void ArrayList_Set(ArrayList<T> *self, size_t index, T value)
    {
        T *ptr = ArrayList_Get(self, index);
        if (ptr == nullptr)
        {
            return;
        }
        *ptr = value;
    }

    template <typename T>
    void ArrayList_Remove(ArrayList<T> *self, std::size_t index)
    {
        if (index >= self->Size)
        {
            throw std::invalid_argument("index exceeds array bounds");
        }
        ArrayList_shiftLeftOverwrite(self, index);
        ArrayList_shrink(self);
    }

    template <typename T>
    size_t ArrayList_Size(ArrayList<T> *self)
    {
        return self->Size;
    }
}