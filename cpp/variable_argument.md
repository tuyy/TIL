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
```

* http://www.cplusplus.com/reference/cstdarg/va_start/
