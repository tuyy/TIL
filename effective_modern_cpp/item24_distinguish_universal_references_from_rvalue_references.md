## Item 24: Distinguish universal references from rvalue references.
### universal references(forwarding reference)와 rvalue references를 구분하라
* forwarding reference를 나타내는 두 가지 문맥
    * 둘 다 형식 연역이 일어나는 공통점을 가진다.
```C++
template <typename T>
void foo1(T&& arg);  // forwarding reference

auto&& boo = arg;  // forwarding reference

Widget&& foo2 = Widget();  // 오른값 참조

template <typename T>
void foo3(std::vector<T>&& param); // 오른값 참조

template <typename T>
void foo4(const T&& param); // 오른값 참조

template <class T>
class MyVector
{
public:
    void push_back(T&& x);  // 오른값 참조
    ...
};

// 이 시점에 T의 형식이 정해진다. 즉, push_back을 호출하는 시점에 형식연역이 일어나지 않는다.
MyVector<Widget> T;

// auto의 forwarding reference example in C++14
auto timeFuncInvocation = [](auto&& func, auto&&... params)
{
    std::forward<decltype(func)>(func)(std::forward<decltype(params)>(params)...);
};
```

### 참고
* Effective Modern C++ by Scott Meyers
