#pragma once

#include <stdint.h>

#include "allocator/vtable_allocator.h"

namespace scr
{
    constexpr size_t ArrayList_InitialCapacity = 5;
    template<typename T>
    struct ArrayList
    {
        T* Data;
        size_t Size;
        size_t Capacity;
        const VTableAllocator* Allocator;
    };
    // using T = int;

    namespace array_list
    {
        template<typename T>
        void Ctor(ArrayList<T>* arrayList, const VTableAllocator* allocator)
        {
            arrayList->Size = 0;
            arrayList->Capacity = ArrayList_InitialCapacity;
            arrayList->Allocator = allocator;
            void* data = allocator->Alloc(sizeof(T) * arrayList->Capacity);
            arrayList->Data = reinterpret_cast<T*>(data);
        }
        
        template<typename T>
        ArrayList<T>* New(VTableAllocator* allocator)
        {
            void* selfVoid = allocator->Alloc(sizeof(ArrayList<T>));
            ArrayList<T>* self = reinterpret_cast<ArrayList<T>*>(selfVoid);
            Ctor(self, allocator);
            return self;
        }

        template<typename T>
        void Dtor(ArrayList<T>* arrayList)
        {
            arrayList->Allocator->Free(arrayList->Data);
            arrayList->Data = nullptr;
            arrayList->Size = 0;
            arrayList->Capacity = 0;
            arrayList->Allocator = nullptr;
        }
    
        template<typename T>
        void expand(ArrayList<T>* arrayList)
        {
            arrayList->Size++;
            if (arrayList->Size < arrayList->Capacity)
            {
                return;
            }
            arrayList->Capacity = arrayList->Capacity * 3 / 2;
            void* newElementsVoid = arrayList->Allocator->ReAlloc(arrayList->Data, arrayList->Capacity);
            T* newElements = reinterpret_cast<T*>(newElementsVoid);
            arrayList->Data = newElements;
        }
        
        template<typename T>
        void shiftRight(ArrayList<T>* arrayList, std::size_t fromIndex)
        {
            if (fromIndex == arrayList->Size || arrayList->Size < 2) // TODO: handle when the size is 2 
            {
                return;
            }
            for (std::size_t i = arrayList->Size; i >= fromIndex; i--)
            {
                arrayList->Data[i] = arrayList->Data[i-1];
            }
        }

        template<typename T>
        void shiftLeftOverwrite(ArrayList<T>* self, std::size_t fromIndex)
        {
            if (fromIndex == 0)
            {
                return;
            }

            for (std::size_t i = fromIndex; i < self->Size; i++)
            {
                self->Data[i] = self->Data[i+1];
            }
        }
        
        template<typename T>
        void shrink(ArrayList<T>* self)
        {
            // no need to deallocate excess space
            self->Size--;
        }

        template<typename T>
        void Insert(ArrayList<T>* self, size_t index, T value)
        {
            expand(self);
            shiftRight(self, index);
            self->Data[index] = value;
        }

        template<typename T>
        void Append(ArrayList<T>* self, T value)
        {
            Insert(self, self->Size, value);
        }

        template<typename T>
        T* Get(ArrayList<T>* self, size_t index)
        {
            if (index >= self->Size)
            {
                return nullptr;
            }
            return self->Data + index;
        }

        template<typename T>
        void Set(ArrayList<T>* self, size_t index, T value)
        {
            T* ptr = Get(self, index);
            if (ptr == nullptr)
            {
                return;
            }
            *ptr = value;
        }
        
        template<typename T>
        void Remove(ArrayList<T>* self, std::size_t index)
        {
            if (index >= self->Size)
            {
                throw std::invalid_argument("index exceeds array bounds");
            }
            shiftLeftOverwrite(self, index);
            shrink(self);
        }
        
        template<typename T>
        size_t Size(ArrayList<T>* self)
        {
            return self->Size;
        }
    }
}