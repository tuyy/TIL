## Item 3: Understand decltype.

### decltype
* C++11 부터 가능
* template & auto의 type deduction과 달리, decltype은 주어진 이름이나 표현식의 구체적인 형식을 그대로 알려준다.
* auto 처럼 변수 선언시 사용할 수 있지만, 좀 더 동일한 형식으로 쓰고 싶을때 decltype(..)을 auto 대신 쓸 수 있다.
```C++
const int32_t age = 100;
auto newAge = age; // C++11, int32_t
decltype(age) newAge = age; // C++11, const int32_t
decltype(auto) newAge = age; // C++14, const int32_t 
```

* 추가로 아래와 같이 함수 반환형식에도 사용할 수 있다.
``` C++
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>

// V1 in C++11
template <typename Container, typename Index>
auto foo_version1(Container &aContainer, const Index aIdx)
    -> decltype(aContainer[aIdx])
{
    return aContainer[aIdx];
}

// V2 in C++14
template <typename Container, typename Index>
auto foo_version2(Container &aContainer, const Index aIdx)
{
    // Container&를 auto형으로 deduction하므로 참조형이 날라간다.
    return aContainer[aIdx];
}

// V3 in C++14
template <typename Container, typename Index>
decltype(auto) foo_version3(Container &aContainer, const Index aIdx)
{
    return aContainer[aIdx];
}

// final v4 in C++14
template <typename Container, typename Index>
decltype(auto) foo_version4(Container &&aContainer, const Index aIdx)
{
    // 보편참조로 rvalue, lvalue 모두 받을 수 있다.
    return std::forward<Container>(aContainer)[aIdx];
}

// Final v4 in C++11
template <typename Container, typename Index>
auto foo_version4(Container &&aContainer, const Index aIdx)
    -> decltype(std::forward<Container>(aContainer)[aIdx])
{
    // 보편참조로 rvalue, lvalue 모두 받을 수 있다.
    return std::forward<Container>(aContainer)[aIdx];
}

int32_t main()
{
    std::vector<std::string> sVec = {"A", "AB", "ABC"};
    foo_version4(sVec, 0) = "1";
    foo_version4(sVec, 1) = "2";
    foo_version4(sVec, 2) = "3";

    auto sPrint = [](const std::string &s){ std::cout << s << std::endl;};
    std::for_each(sVec.cbegin(), sVec.cend(), sPrint);
    return 0;
}
```

* 주의할 점
```C++
decltype(auto) f1()
{
    int32_t i = 100;
    return i;  // decltype(i) is int, so f1 returns int
}

// compile warning!
decltype(auto) f2()
{
    int32_t i = 100;
    return (i);  // decltype(i) is int&, so f2 returns int&
}
```

### 참고
* Effective Modern C++ by Scott Meyers
