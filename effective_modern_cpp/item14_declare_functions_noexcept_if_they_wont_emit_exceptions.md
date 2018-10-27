## Item 14: Declare functions noexcept if they won’t emit exceptions.
### 예외를 방출하지 않는 함수에 noexcept를 선언하라
* C++98에서 예외 명세가 있지만 비권장 기능으로 분류된다.
```C++
int f(int x) throw();  // f는 예외를 방출하지 않음: C++98 방식
int f(int x) noexcept;  // f는 예외를 방출하지 않음: C++11 방식
```
* 호출 스택이 풀리지 않는다고 컴파일러에게 알려주므로 최적화 여지가 가장 크다.
```C++
int f(int x) noexcept;  // 최적화 여지가 가장 크다. (C++11에선 스택이 풀릴 수도 있고 풀리지 않을 수도 있다.)
int f(int x) throw();  // 최적화 여지가 더 작다. (C++98에선 스택이 f를 호출한 곳 까지 풀린다.)
int f(int x);  // 최적화 여지가 더 작다.
```
* 강한 예외 안전성을 보장할 수 있다.
    * 표준 라이브러리에선 이동 연산이 예외를 방출하지 않음(noexcept)이 알려진 경우에만 C++98의 복사 연산을 C++11의 이동 연산으로 대체한다.
    * 표준 라이브러리의 여러 함수는 "가능하면 이동하되 필요하면 복사한다." 전략을 사용하기 때문이다.
* swap 함수들에 유용하게 사용할 수 있다.
```C++
template <typename T, size_t N>
void swap(T (&a)[N], T(&b)[N]) noexcept(noexcept(swap(*a, *b))); // 조건부 noexcept

template <class T1, class T2>
struct pair
{
    ...
    void swap(pair& p) noexcept(noexcept(swap(first, p.first)) && noexcept(swap(second, p.second)));
};
```
* 모든 메모리 해제 함수(delete)와 모든 소멸자는 암묵적으로 noexcept 선언되었다.
### 최적화가 중요하긴 하지만, 그래도 가장 중요한 것은 coreectness(정확성)이다.

### 참고
* Effective Modern C++ by Scott Meyers
