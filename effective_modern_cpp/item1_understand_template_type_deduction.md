## Item 1: Understand template type deduction.

```C++
// template <typename T>
// void f(ParamType param)
// f(expr);

template <typename T>
void f(const T &param);

int x = 0;
f(x); // T:int ParamType:const int&
```

### 1. ParamType이 참조&포인터이고 보편참조가 아닌 경우.
* expr의 const성과 ParamType의 const은 그대로 유지된다.
* ParamType은 const까지 포함해서 함께 deduction된다.

### 2. ParamType이 보편참조(universal reference)인 경우.
* expr의 const성은 그대로 유지된다. (보편참조는 ParamType이 const을 가질 수 없다.)
* lvalue인 경우 T&, rvalue인 경우 T&&로 deduction된다.

### 3. ParamType이 참조&포인터&보편참조가 아닌 경우.
* expr의 const은 무시된다. 반면 ParamType의 const성은 유지된다.
* expr의 참조성은 제거되고(복사가 발생하므로), 포인터는 그대로 유지된다.

### 4. expr이 배열인 경우.
* 일반적으로 배열은 포인터 형식으로 deduction된다.
* 하지만 ParamType에 &를 추가하면, 배열참조형식으로 deduction된다.
* 이를 이용하면 아래와 같이 배열의 크기를 구하는 템플릿 함수를 작성할 수 있다.
```
template <typename T, std::size_t N>
constexpr std::size_t arraySize(const T (&)[N]) noexcept
{
    return N;
}
```

### 5. 함수 인수
* 일반적으로 배열인수와 마찬가지로 포인터로 deduction된다.
* ParamType에 &를 추가하면 참조함수형식으로 deduction된다.

### 참고
* Effective Modern C++ by Scott Meyers
