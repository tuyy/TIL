### std::shared_mutex in C++17

#### shared_mutex는 lock mode를 두 가지 지원
* exclusive mode는 lock이 걸리면 하나의 쓰레드만 접근 가능하다. (일반 mutex도 가능)
* shared mode는 lock이 걸리면 읽기 자원에 대하여 여러 쓰레드가 동시에 접근가능하다.

```C++
// exclusive mode
{
std::unique_lock<std::mutex> lock;
..
}

// shared mode
{
std::shared_lock<std::shared_mutex> lock;
..
}
```
    
#### lock 기준
* shared mode로 lock이 잡히면 exclusive mode의 lock은 block 상태가 된다.
* exclusive mode로 lock이 잡히면 shared mode의 lock은 block 상태가 된다.
* shared mode로 lock이 threshold 만큼 잡히면 이후 shared mode lock 요청은 block 된다.

#### 결론
* 고도의 성능이 필요하지 않는 이상, mutex를 쓰자. 시스템이 복잡해지기만 한다.
    
* https://en.cppreference.com/w/cpp/named_req/SharedMutex
