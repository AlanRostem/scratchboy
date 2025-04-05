#pragma once

#include <stdint.h>

#include "allocator/vtable_allocator.h"

namespace scr
{
    constexpr size_t ArrayList_InitialCapacity = 5;
    template <typename T>
    struct ArrayList
    {
        T *data;
        size_t size;
        size_t capacity;
        const VTableAllocator *allocator;
    };
    // using T = int;

    template <typename T>
    void ArrayList_Init(ArrayList<T> *self, const VTableAllocator *allocator)
    {
        self->size = 0;
        self->capacity = ArrayList_InitialCapacity;
        self->allocator = allocator;
        void *data = allocator->Alloc(sizeof(T) * self->capacity);
        self->data = reinterpret_cast<T *>(data);
    }

    template <typename T>
    ArrayList<T> *NewArrayList(const VTableAllocator *allocator)
    {
        void *selfVoid = allocator->Alloc(sizeof(ArrayList<T>));
        ArrayList<T> *self = reinterpret_cast<ArrayList<T> *>(selfVoid);
        ArrayList_Init(self, allocator);
        return self;
    }

    template <typename T>
    void ArrayList_DeInit(ArrayList<T> *self)
    {
        self->allocator->Free(self->data);
        self->data = nullptr;
        self->size = 0;
        self->capacity = 0;
        self->allocator = nullptr;
    }

    template <typename T>
    void ArrayList_expand(ArrayList<T> *self)
    {
        self->size++;
        if (self->size < self->capacity)
        {
            return;
        }
        self->capacity = self->capacity * 3 / 2;
        void *newElementsVoid = self->allocator->ReAlloc(self->data, self->capacity);
        T *newElements = reinterpret_cast<T *>(newElementsVoid);
        self->data = newElements;
    }

    template <typename T>
    void ArrayList_shiftRight(ArrayList<T> *self, size_t fromIndex)
    {
        if (fromIndex == self->size || self->size < 2) // TODO: handle when the size is 2
        {
            return;
        }
        for (size_t i = self->size; i >= fromIndex; i--)
        {
            self->data[i] = self->data[i - 1];
        }
    }

    template <typename T>
    void ArrayList_shiftLeftOverwrite(ArrayList<T> *self, size_t fromIndex)
    {
        if (fromIndex == 0)
        {
            return;
        }

        for (size_t i = fromIndex; i < self->size; i++)
        {
            self->data[i] = self->data[i + 1];
        }
    }

    template <typename T>
    void ArrayList_shrink(ArrayList<T> *self)
    {
        // no need to deallocate excess space
        self->size--;
    }

    template <typename T>
    void ArrayList_Insert(ArrayList<T> *self, size_t index, T value)
    {
        ArrayList_expand(self);
        ArrayList_shiftRight(self, index);
        self->data[index] = value;
    }

    template <typename T>
    void ArrayList_Append(ArrayList<T> *self, T value)
    {
        ArrayList_Insert(self, self->size, value);
    }

    template <typename T>
    T *ArrayList_Get(ArrayList<T> *self, size_t index)
    {
        if (index >= self->size)
        {
            return nullptr;
        }
        return self->data + index;
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
    void ArrayList_Remove(ArrayList<T> *self, size_t index)
    {
        if (index >= self->size)
        {
            return;
        }
        ArrayList_shiftLeftOverwrite(self, index);
        ArrayList_shrink(self);
    }

    template <typename T>
    size_t ArrayList_Size(ArrayList<T> *self)
    {
        return self->size;
    }

    template<typename T>
    struct ArrayListIterator
    {
        ArrayList<T>* arrayList;
        int64_t index;
    };

    template<typename T>
    bool ArrayListIterator_Next(ArrayListIterator<T>* iterator)
    {
        if (iterator->index == static_cast<int64_t>(ArrayList_Size(iterator->arrayList))-1)
        {
            return false;
        }

        iterator->index++;
        return true;
    }

    template<typename T>
    T* ArrayListIterator_Get(ArrayListIterator<T>* self)
    {
        return ArrayList_Get(self->arrayList, self->index);
    }

    template<typename T>
    ArrayListIterator<T> ArrayList_Iterate(ArrayList<T>* self)
    {
        return ArrayListIterator<T>{
            .arrayList = self,
            .index = -1,
        };
    }
}