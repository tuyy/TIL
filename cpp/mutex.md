## std::mutex
* thread 간 공유 자원에 대해 깃발 꼽기!
* std::lock_guard는 생성시점에 m.lock()과 소멸자 호출할 때 m.unlock()이 호출된다. (RAII)
* std::mutex를 사용하면 확실히 성능적으로 느리다;
* 멤버 함수
```C++
m.lock()
m.unlock()
m.try_lock() // lock 획득 시도, 성공시 true 반환 및 lock 획득하고 실패시 false 반환
```

```C++
#include <iostream>
#include <thread>
#include <mutex>
#include <algorithm>
#include <vector>
#include <functional>

template <typename T>
T boo(T aVal)
{
    aVal *= 999;
    return aVal + 10;
}

void foo(std::mutex &aMutex, const std::string &aName)
{
    for (int32_t i = 0; i < 10000000; ++i)
    {
        //std::unique_lock sLock(aMutex);
        auto k = i * i * i * i;
        boo(k);
    }
}

auto main() -> int32_t
{
    std::cout << "Start main()" << std::endl;

    std::mutex sMutex;

    std::vector<std::thread> sVec;
    sVec.emplace_back(foo, std::ref(sMutex), "first");
    sVec.emplace_back(foo, std::ref(sMutex), "second");
    sVec.emplace_back(foo, std::ref(sMutex), "third");
    sVec.emplace_back(foo, std::ref(sMutex), "fourth");
    sVec.emplace_back(foo, std::ref(sMutex), "fifth");

    std::this_thread::sleep_for(std::chrono::seconds(1));

    if (sMutex.try_lock() == true)
    {
        std::cout << "main thread가 lock 획득 성공" << std::endl;
        sMutex.unlock();
    }
    else
    {
        std::cout << "main thread가 lock 획득 실패" << std::endl;
    }

    std::for_each(sVec.begin(), sVec.end(), [](auto &v){ v.join(); });

    return 0;
}
```
