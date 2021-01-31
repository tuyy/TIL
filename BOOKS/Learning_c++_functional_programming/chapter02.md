## 함수형 프로그래밍에서 함수 다루기
### 일급 함수
* 언어에서 제공하는 것. 
* 다른 함수의 매개변수로 함수 전달 가능, 변수에 함수 대입 가능, 컨테이너에 함수 저장 가능, 런타임에 새로운 함수 생성 가능

### 고차 함수
* 일급 함수와 의미가 비슷. 하나 이상의 함수를 인수로 사용할 수 있으며 반환 값으로 함수 사용이 가능

### 고차함수 map, filter, reduce 와 STL

#### map은 std::transform 함수에 매핑된다.

#### filter는 std::copy_if 함수에 매핑된다. (reject은 std::remove_copy_if)
```C++
std::vector<std::string> sNewVec;
std::copy_if(sVec.begin(),
             sVec.end(),
             std::back_inserter(sNewVec),
             [](const auto &v) { return v.size() > 4; });
```

#### reduce는 std::accumulate 함수에 매핑된다.
```C++
#include <numeric>

std::vector<int32_t> sVec = {1, 2, 3, 4, 5};

auto foldl = std::accumulate(sVec.begin(), sVec.end(), 0, [](auto a, auto b){return a+b*2;});
std::cout << foldl << std::endl; // 15

auto foldl2 = std::accumulate(sVec.begin(), sVec.end(), 0);
std::cout << foldl2 << std::endl; // 15
```

### 순수함수를 작성하라.
* 모든 함수는 외부 변수의 값을 읽거나 변경하지 않도록 만들 수 있다.
* 이를 통해 의도치 않은 동작을 하지 않게 할 수 있다.

### 커링으로 함수 분리하기
```C++
template <typename Func, typename... Args>
auto curry(Func func, Args... args)
{
    return [=](auto... lastParam)
    {
        return func(args..., lastParam...);
    };
}

auto add = [](auto a, auto b){return a+b;};
std::cout << add(1, 2) << std::endl;

auto add10 = curry(add, 10);
std::cout << add10(2) << std::endl;
```
