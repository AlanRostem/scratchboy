#pragma once

#define SCR_CONCAT(a, b) SCR_CONCAT_INNER(a, b)
#define SCR_CONCAT_INNER(a, b) a##b

#define SCR_SCOPE_EXIT(x) auto SCR_CONCAT(__outOfScopeCaller, __LINE__) = scr::__OutOfScopeCaller([&]() { x; })

namespace scr
{
    template<typename TLambda>
    class __OutOfScopeCaller
    {
    public:
        __OutOfScopeCaller(TLambda lambdaFunc)
            : m_func(lambdaFunc)
        {

        }

        ~__OutOfScopeCaller()
        {
            m_func();
        }
    private:
        TLambda m_func;
    };
} // namespace scr
