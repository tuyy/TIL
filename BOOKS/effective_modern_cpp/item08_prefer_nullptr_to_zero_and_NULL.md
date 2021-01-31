## Item 8: Prefer nullptr to 0 and NULL.
### 빈 포인터를 나타내기 위해 NULL, 0을 사용하기 보단 C++11 부터 제공하는 nullptr을 사용하자.
* NULL, 0 은 int타입이며 포인터가 아니다.
```C++
void f(int);
void f(bool);
void f(void*);

f(0); // f(void*)이 아니라 f(int)를 호출
f(NULL); // 컴파일되지 않을 수도 있지만 보통 f(int)를 호출
f(nullptr) // f(void*) 호출
```
* nullptr은 코드의 명확성을 높여준다.
```C++
auto result = findRecord(..);
if (result == 0)  // result가 pointer인지 int인지 확신할 수 없다.
{
    ..
}

auto result = findRecord(..);
if (result == nullptr)  // result가 pointer인 것은 명백하다.
{
    ..
}
```
* 템플릿 형식 deduction에선 0과 NULL을 틀린 형식으로 인지한다.
```C++
template <typename FuncType, typename MuxType, typename PtrType>
decltype(auto) lockAndCall(FuncType func, MuxType &mutex, PtrType ptr)
{
    using MuxGuard = std::lock_guard<MuxType>;
    MxGuard g(mutex);
    return func(ptr);
}

auto result1 = lockAndCall(f1, f1m, 0); // error!
auto result2 = lockAndCall(f2, f2m, NULL); // error!
auto result3 = lockAndCall(f3, f3m, nullptr); // Ok
```

### 참고
* Effective Modern C++ by Scott Meyers
