## Item 2: Understand auto type deduction.

### auto
* C++11 부터 가능
* 한 가지 상황을 제외하고는 template deduction과 직접적인 대응관계가 존재한다.
* 결과적으로 template deduction과 auto의 deduction은 동일하다. (한 가지만 빼고!)
```C++
template <typename T>
void f(ParamType param);
f(expr);

// auto를 이용해서 변수를 선언할 때 auto는 템플릿의 T와 동일한 역할을 하며,
// 변수의 type specifier는 ParamType과 동일한 역할을 한다.
const auto k = expr;
```

* 한 가지는 바로, 중괄호 초기화를 넣는 경우이다.
```C++
auto k = {0}; // std::initializer_list<int>로 deduction

// template의 경우
template <typename T>
void f(T param);
f({1});  // compile error!!!!

template <typename T>
void f(std::initializer_list<T> param);  // compile error 아님!
```

* C++14 에서는 알아야할 것이 하나 더 남아있다.
```C++
auto createInitList()
{
    return {1, 2, 3};  // 템플릿 형식과 동일한 에러 발생!
}

// 람다 메개변수 형식에서도 마찬가지로 템플릿 형식과 동일한 에러 발생!
auto resetV = [&v](const auto& newValue) { v = newValue; };   // C++14
resetV({1, 2, 3});
```

### 참고
* Effective Modern C++ by Scott Meyers


