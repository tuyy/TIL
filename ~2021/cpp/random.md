## random
* https://en.cppreference.com/w/cpp/header/random

```C++
#include <map>
#include <random>

int32_t main()
{
    std::random_device sRd; // seed 생성기
    std::mt19937 sGen(sRd()); // 난수 생성기
    std::uniform_int_distribution<int32_t> sDis(0, 99); // 난수를 고르게 분포하도록함

    std::map<int32_t, size_t> sMap;

    for (size_t i = 0; i < 10; ++i)
    {
        auto sVal = sDis(sGen);
        if (auto sIter = sMap.find(sVal); sIter != sMap.end())
        {
            sIter->second += 1;
        }
        else
        {
            sMap[sVal] = 1;
        }
    }

    for (auto &[k, v] : sMap)
    {
        std::cout << "NUM:" << k << " COUNT:" << v << std::endl;
    }
    return 0;
}
```


```
$ ./main
난수생성:2970692652
NUM:18 COUNT:1
NUM:25 COUNT:1
NUM:31 COUNT:1
NUM:40 COUNT:1
NUM:58 COUNT:1
NUM:60 COUNT:2
NUM:64 COUNT:1
NUM:77 COUNT:1
NUM:82 COUNT:1
```
