## std::this_thread
* Since C++11, ```#include <thread>```
* sleep, get_id, yield 함수가 있다.

### 1) std::this_thread::sleep_for(..), std::this_thread::sleep_until(..)
* std::chrono 의 시간 값을 인자로 받는다.
```C++
#include <iostream>
#include <thread>
#include <chrono>

auto main() -> int32_t
{
    std::cout << "Start main().." << std::endl;
    std::this_thread::sleep_until(std::chrono::high_resolution_clock::now() + std::chrono::seconds(3));

    for (int32_t i = 0; i < 5; ++i)
    {
        std::cout << "Hello? " << i << std::endl;

        std::this_thread::sleep_for(std::chrono::milliseconds(500));
    }

    return 0;
}
```

### 2) sed::this_thread::get_id()
* 현재 thread id 를 구한다.
```C++
#include <iostream>
#include <thread>

auto main() -> int32_t
{
    std::cout << "Start main().." << std::endl;

    std::cout << std::this_thread::get_id() << std::endl;

    return 0;
}
```

```
[irteam@dev-tuyy-ncl 023_thread_ex]$ g++ -std=c++17 -lpthread -o second second.cpp ;./second
Start main()..
139955340449568
```

### 3) std::this_thread::yield()
* 현재 thread에서 다른 thread를 실행 할 수 있도록한다.
```C++
#include <iostream>
#include <chrono>
#include <thread>

// "busy sleep" while suggesting that other threads run
// for a small amount of time
void little_sleep(std::chrono::microseconds us)
{
    auto start = std::chrono::high_resolution_clock::now();
    auto end = start + us;
    do {
        std::this_thread::yield();
    } while (std::chrono::high_resolution_clock::now() < end);
}

int main()
{
    auto start = std::chrono::high_resolution_clock::now();

    little_sleep(std::chrono::microseconds(100));

    auto elapsed = std::chrono::high_resolution_clock::now() - start;
    std::cout << "waited for "
              << std::chrono::duration_cast<std::chrono::microseconds>(elapsed).count()
              << " microseconds\n";
}
```

```
[irteam@dev-tuyy-ncl 023_thread_ex]$ g++ -std=c++17 -lpthread -o yield yield.cpp ;./yield
waited for 102 microseconds
```
