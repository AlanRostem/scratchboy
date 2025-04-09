# ScratchBoy

A GameBoy emulator programmed using a "minimal" C++.

## Coding Conventions

The emulator is written in C++, but only some C++ features are used, such as templates and **very limited** use of namespaces. The programming paradigm for this library is to program things like it was C, but with some minimal added features of C++. For example, a procedural approach is used instead of OOP, but one can still use templates. 

### Naming Conventions

- All type names use `PascalCase`.
  - Including `struct` and `using` statements.
- Virtual function table structs use the `VTable<name>` prefix
- Constants use `PascalCase` and are always inside nested namespaces.
- Standalone functions use `PascalCase` and are always inside nested namespaces.
- Namespaces use `snake_case`
- Stack-allocated variable names use `camelCase`.
  - I.e., local variables on the stack (both in blocks of code and parameters).
- Struct fields that are intended to be **public** use `PascalCase`.
- Struct fields that are intended to be **private** use `camelCase`.
- Struct methods that are intended to be **public** are always in the `.h` file and named as follows: `StructTypeName_PublicFunctionName(StructTypeName* self, ...)`
  - I.e., type-name is in `PascalCase`, followed up by `_` and the `PascalCase` "method" name
- Struct methods that are intended to be **private** are always in the `.cpp` file and named as follows: `StructTypeName_privateFunctionName(StructTypeName* self, ...)`
  - I.e., type-name is in `PascalCase`, followed up by `_` and the `camelCase` "method" name

### Singleton Namespaces

The implementation of "singletons" is done by creating a namespace for the singleton and defining its functions there.

Example:

```cpp
#include "vtable_allocator.h"

namespace scr
{
    namespace system_heap
    {
        const VTableAllocator* Get();
    }
}
```

### Allocators and Object Construction

Usage of standard allocation functions, such as `malloc` and `free`, is **strictly forbidden**. Please use the singleton `system_heap::Get()` defined in `src/allocator/heap.h`. This function wraps `malloc`, `free`, and `realloc` into a v-table struct. The reason behind this is to enforce a standard rule: *use allocator v-tables when constructing an object dynamically*. This makes all object constructors compatible with any type of allocator. 

Example 1: The `ArrayList` struct needs an allocator to store its elements and the `NewArrayList` function returns a pointer to an `ArrayList` allocated by the given allocator v-table.

```cpp
const scr::VTableAllocator *heap = scr::system_heap::Get();
scr::ArrayList<int> *list = scr::NewArrayList<int>(heap);
// do stuff...
scr::ArrayList_DeInit(list);
heap->Free(list);
```

Example 2: Same as Example 1, but the contents of `NewArrayList` is manually replicated here.

```cpp
const scr::VTableAllocator *heap = scr::system_heap::Get();
void* listAddr = heap->Alloc(sizeof(scr::ArrayList<int>));
scr::ArrayList<int> *list = reinterpret_cast<scr::ArrayList<int>*>(listAddr);
scr::ArrayList_Init(list, heap);
// do stuff...
scr::ArrayList_DeInit(list);
heap->Free(list);
```

This approach allows for usage of any type of allocator in the `ArrayList` struct when it needs to allocate memory for objects. Additionally, it clarifies in the scope that an allocator was used which (hopefully) reminds the developer to free the memory that was used.
