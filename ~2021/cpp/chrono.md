### chrono
* C++11 이후 지원하는 date and time library
* ```#include <chrono>```

#### clock
```C++
#include <chrono>

// 1. steady_clock
// monotonic_clock 란 별칭을 가지고 있다.
// os 시간을 돌리면 영향을 받지 않는다. 계속 증가만한다.
auto sStart = std::chrono::steady_clock::now();

// 2. system_clock
// 가장 일반적인 clock 클래스이고, C time 구조체와 호환된다.
// os 시간을 돌리면 영향을 받는다.
// auto sStart = std::chrono::system_clock::now();

// 3. high_resolution_clock 
// system_clock 의 typedef 이다.
// auto sStart = std::chrono::high_resolution_clock::now();

// doit..

auto sEnd = std::chrono::steady_clock::now();

std::chrono::duration<double> sDuration = sEnd - sStart;

std::cout << "Duration:" << sDuration.count() << std::endl;

```
