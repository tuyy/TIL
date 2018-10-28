## Item 23: Understand std::move and std::forward.
### std::move와 std::forward를 숙지하자
* std::move 는 무조건 rvalue 캐스팅이다.
    * 주의할 점은 const 가 포함된 값을 넘기면 복사생성자가 호출된다.
```C++
// C++11
template <typename T>
typename remove_reference<T>::type&& move(T&& param)
{
    using ReturnType = typename remove_reference<T>::type&&;
    return static_cast<ReturnType>(param);
}

// C++14
template <typename T>
decltype(auto) move(T&& param)
{
  using ReturnType = remove_reference_t<T>&&;
  return static_cast<ReturnType>(param);
}
```
* std::forward는 조건부 캐스팅이다.
    * 오른값에 묶인 참조만 오른값으로 캐스팅한다는 의미이다.
```C++
void process(const Widget& lvalArg); // 왼값 호출
void process(Widget&& rvalArg); // 오른값 호출

template <typename T>
void logAndProcess(T&& param) // univesal reference
{
    auto now = std::chrono::system_clock::now();
    
    makeLogEntry("hello", now);
    process(std::forward<T>(param));
}
```
* std::move, std::forward 둘 다, 실행시점에서는 아무 일도 하지 않는다.

### 참고
* Effective Modern C++ by Scott Meyers
