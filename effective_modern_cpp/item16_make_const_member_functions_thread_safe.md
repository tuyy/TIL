## Item 16: Make const member functions thread safe.
### const 멤버함수를 thread에 안전하게 작성하라
* thread에 안전하지 않은 const 멤버함수 예시
```C++
class Polynomial
{
public:
    using RootsType = std::vector<double>;
    
    RootsType roots() const
    {
        if (!rootsAreValid)
        {
            ...
        }
        rootsAreValid = true;
        return rootVals;
    }

private:
    mutable bool rootsAreValid{false}; // const 멤버함수에서 변경가능하도록하기 위해 mutable 키워드 추가
    mutable RootsType rootVals{};
};

Polynomial p;
...
auto rootsOfP = p.roots(); // thread1
auto valsGivingZero = p.roots(); // thread2
... // data race 발생
```
* std::mutex로 해결
```C++
class Polynomial
{
public:
    using RootsType = std::vector<double>;
    
    RootsType roots() const
    {
        std::lock_guard<std::mutex> g(m);
        
        if (!rootsAreValid) {
            ...
            rootsAreValid = true;
        }
        return rootVals;
    }
private:
    mutable std::mutex m; // std::mutex를 사용하면 복사하거나 이동할 수 없다.
    mutable bool rootsAreValid{false}; // const 멤버함수에서 변경가능하도록하기 위해 mutable 키워드 추가
    mutable RootsType rootVals{};
};
```
* std::mutex의 비용이 너무 과할 땐, std::atomic 카운터를 사용해서 비용을 줄일 수 있는 경우가 많다.
    * 다른 스레드들이 보기에 해당 연산들이 중간에 가로채이지 않고 한 덩어리로 실행됨이 보장되는 카운터
    * 실제 비용이 절감되는지는 프로그램이 실행되는 하드웨어에 따라 다를 수 있다.
    ```C++
    class Point
    {
    public:
        ...
        double distanceFromOrigin() const noexcept
        {
            ++callCount;
            return std:hypot(x, y);
        }
    private:
        mutable std::atomic<unsigned> callCount{0};
        double x, y;
    };
    ```
* 동기화가 필요한 변수 하나 또는 메모리 장소 하나에 대해서는 std::atomic을 사용하는 것은 적합하다.
    * 둘 이상의 변수나 메모리 장소를 하나의 단위로서 조작해야 할 때에는 뮤텍스를 꺼내는 것이 바람직하다.

### 참고
* Effective Modern C++ by Scott Meyers
