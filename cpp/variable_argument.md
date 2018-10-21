#### variable_argument
* ```#include <cstdarg>```
* 가변인자 처리방법
```C++

#include <string>
#include <cstdarg>
#include <cstdio>

int32_t sumInteger(int32_t count, ...)
{
    int32_t sum = 0;
    va_list args;
    va_start(args, count);
    
    for (int32_t i = 0; i < count; ++i)
    {
      sum += va_arg(args, int);
    }
    
    va_end(args);
    return sum;
}

std::string makeString(const char *format, ...)
{
      va_list sAp;
      va_start(sAp, aFormat);
      int32_t sLen = std::vsnprintf(nullptr, 0, aFormat, sAp);
      va_end(sAp);

      std::unique_ptr<char[]> sBuffer(new char[sLen + 1]);

      va_start(sAp, aFormat);
      (void)std::vsnprintf(sBuffer.get(), sLen + 1, aFormat, sAp);
      va_end(sAp);

      return std::string(sBuffer.get(), sLen);
}

// Parameter pack. C++11 이후부터 가능
template <typename ... Args>
std::string makeString(const char *aFormat, Args ... aArgs)
{
    if (aFormat == nullptr)
    {
        return std::string();
    }
    else
    {
        int32_t sLen = std::snprintf(nullptr, 0, aFormat, aArgs...);
        std::cout << "length:" << sLen << std::endl;

        auto sBuffer = std::make_unique<char[]>(sLen + 1);

        (void)std::snprintf(sBuffer.get(), sLen + 1, aFormat, aArgs...);

        return std::string(sBuffer.get(), sLen);
    }
}
```

* https://en.cppreference.com/w/cpp/language/parameter_pack
* http://www.cplusplus.com/reference/cstdarg/va_start/
