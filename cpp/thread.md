## std::thread
* C++11 이후 사용 가능. ```#include <thread>```
* std::thread 객체 초기화와 동시에 실행된다.
```C++
#include <thread>

void foo(int arg1, int arg2, int arg3)
{
    // TODO
}

std::thread sT1(foo, 1, 2, 3); // thread가 실행할 함수와 인자를 함께 넘길 수 있다.
sT1.join(); // thread가 끝날때까지 기다린다.
```
* std::thread에 넘겨진 함수의 return 값은 무시되고 예외가 throw되면 std::terminate가 호출된다. (프로그램 종료)
* std::mutex, std::atomic 등을 이용하여 thread 간에 communication 할 수 있다.
* std::promise를 이용하여 return 값이나 예외를 알 수 있다.
* 어떤 thread도 나타내지 않을 수 있다.
```C++
// 1) 기본 생성 (sT1)
std::thread sT1;

// 2) move 이후 남겨진 thread (sT2)
std::thread sT2(foo, 1, 2, 3);
std::thread sT3 = std::move(sT2); 

// 3) join or detach 이후
sT3.join();
sT3.detach();
```
* 두 개의 thread 객체는 같은 상태를 나타낼 수 없다. 즉 이동 생성자만 지원한다.
```C++
std::vector<std::thread> sVec;
sVec.push_back(std::thread(foo, 1, 2, 3));
sVec.push_back(std::thread(foo, 2, 3, 4));
sVec.push_back(std::thread(foo, 3, 4, 5));

std::for_each(sVec.cbegin(), sVec.cend(), [](auto &v){ v.join(); });
```

* 사용 가능한 멤버 함수가 매우 직관적이다.
    * public 멤버 함수: join(), detach(), joinable(), operator=(&&), swap(), get_id()

* example
```C++
#include <iostream>
#include <vector>
#include <thread>
#include <chrono>
#include <string>
#include <algorithm>

void boo(const size_t aSize, const std::string &sPrefix = "")
{
    std::cout << "boo() id:" << std::this_thread::get_id() << std::endl;
    throw std::runtime_error("raise exception.."); // 프로그램 종료
}

void foo(const size_t aSize, const std::string &sPrefix = "")
{
    std::cout << "foo() id:" << std::this_thread::get_id() << std::endl;

    for (size_t i = 0; i < aSize; ++i)
    {
        std::cout << sPrefix << " "  << i << std::endl;

        std::this_thread::sleep_for(std::chrono::milliseconds(100));
    }
}

auto main() -> int32_t
{
    std::cout << std::thread::hardware_concurrency() << " concurrent threads are supported." << std::endl;

    std::cout << "main() id:" << std::this_thread::get_id() << std::endl;

    std::vector<std::thread> sVec(3);

    std::for_each(std::begin(sVec), std::end(sVec), [](auto &v){ v = std::thread(foo, 30, "hello" ); });

    std::cout << "before join" << std::endl; // t.joinable() return true

    std::for_each(std::begin(sVec), std::end(sVec), [](auto &v){ v.join(); });

    std::cout << "after join" << std::endl; // t.joinable() return false

    return 0;
}

```
