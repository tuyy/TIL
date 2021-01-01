### makeString
* formatting 하여 std::string을 만드는 함수 제작

```C++
#include <iostream>
#include <memory>
#include <cstdarg>

#define BUFF_SIZE 32

// 1) 임의로 Buf size를 할당하여 makeString
std::string makeString(const char *aFmt, ...)
{
    constexpr int32_t sLen = 4096;
    va_list sAp;

    if (aFmt == nullptr)
    {
        return std::string("");
    }

    auto sBuf = std::make_unique<char[]>(sLen + 1);

    va_start(sAp, aFmt);
    (void)std::vsnprintf(sBuf.get(), sLen + 1, aFmt, sAp);
    va_end(sAp);

    return std::string(sBuf.get(), sLen);
}

// 2) Buf size를 구하여 makeString
std::string makeString2(const char *aFmt, ...)
{
    int32_t sLen;
    va_list sAp;

    if (aFmt == nullptr)
    {
        return std::string("");
    }

    va_start(sAp, aFmt);
    sLen = std::vsnprintf(nullptr, 0, aFmt, sAp);
    va_end(sAp);

    auto sBuf = std::make_unique<char[]>(sLen + 1);

    va_start(sAp, aFmt);
    (void)std::vsnprintf(sBuf.get(), sLen + 1, aFmt, sAp);
    va_end(sAp);

    return std::string(sBuf.get(), sLen);
}

auto main() -> int32_t
{
    for (size_t i = 0; i < 1000000; ++i) // 100만번 반복
    {
        /* 아래 주석 하나씩 풀면서 elapsed time을 체크해보자 */
        // std::string sRz = makeString("%s %s %d %d %d %f %f", "hello", "world", 100, 200, 300, 0.5);
        //std::string sRz = makeString2("%s %s %d %d %d %f", "hello", "world", 100, 200, 300, 0.5);
        //std::string sRz = "hello world " + std::to_string(100) + std::to_string(200) + std::to_string(300) + std::to_string(0.5);
    }

    return 0;
}
```

#### 1) 임의로 Buf size를 할당하여 makeString 함수 확인
```
[irteam@dev-tuyy-ncl 017_snprintf_ex]$ g++ -std=c++17 -o main PrintTest.cpp ;time ./main

real    0m4.197s
user    0m4.192s
sys     0m0.000s
```

#### 2) Buf size를 계산한 이후에 makeString 함수 확인
```
[irteam@dev-tuyy-ncl 017_snprintf_ex]$ g++ -std=c++17 -o main PrintTest.cpp ;time ./main

real    0m1.637s
user    0m1.621s
sys     0m0.003s
```

#### 3) std::to_string 함수와 std::string operator= 함수 연산을 이용하여 계산
```
[irteam@dev-tuyy-ncl 017_snprintf_ex]$ g++ -std=c++17 -o main PrintTest.cpp ;time ./main

real    0m2.449s
user    0m2.445s
sys     0m0.003s
```

#### 결론
* 3번은 확실히 느리다. 물론 std::string을 더하는 값이 별로 없다면 복사가 덜 발생하므로 더 빠를 수 있다.
* 사실 buf size를 우연히 비슷하게 맞춘다면 1번이 2번보다 성능이 더 좋다.(당연한가?)
* 반대로 임의의 buf size가 더 작다면? 원하는 결과를 얻을 수 없다.
* 결국 2번 방법으로 문제를 해결하는게 가장 합리적으로 보인다.
