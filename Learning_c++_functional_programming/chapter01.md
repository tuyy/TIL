## 모던 C++과 친숙해지기

### 새로운 기능 익히기
* auto, decltype, nullptr, 비멤버 begin()/end(), range-based for, lambda, smart pointer, tuple 등

#### auto 키워드로 데이터 타입을 자동으로 정의하기
```C++
int v1 = 10;
const v2 = 10;
int &v3 = v1;
const int v4 = v2;
int &&v5 = 10;
const &&v6 = 10;

auto v = .. v1,2,3,4,5,6 .. // 모두 int로 값 복사
const auto v = .. // 모두 const int로 값 복사

auto &v = .. // const 인 값은 const가 붙고 나머진 모두 & 적용
const auto &v = .. // 모두 const & 적용

// 예외사항 universal reference
int v1 = 10; // or int& or int&&
auto &&v = v1; // int&
const auto &&v = v1; // const int&

const int v1 = 10; // or const int& or const int&&
auto &&v = v1; // const int&
const auto &&v = v1; // const int&
```

* trailing return type
```C++
auto main() -> int
{
    // ..
}
```

#### decltype 키워드로 표현식 타입 질의하기
* decltype은 컴파일 타임에 객체나 표현식의 타입을 알고 싶을 때 사용한다.
```C++
decltype(func1()) f1;
decltype(i) i1;
decltype(x->d) x1;
..

// 템플릿과 auto, delcltype의 사용
template <typename I, typename J>
auto add(I i, J j) -> decltype(i + j)
{
    return i + j;
}
```

#### null 포인터 ```nullptr```
#### 비멤버 함수 begin(), end()
#### 범위 기반 for 루프
#### std::array
```C++
std::array<int, 3> sArr = {1, 2, 3};
// 왜 std::array를 쓰나? 그냥 int[3]은 안되나?
// std::array는 더 명식적이고 STL을 쓸 수 있다.

// tip! 컬렉션 요소 접근에 [idx] 보다 .at을 써라! 잘못된 인덱스를 참조하는 경우 미정의 동작이 발생한다.
```

#### 알고리즘 사용하기 ```#include <algorithm>```
#### lambda 표현식
* 일급 함수와 순수 함수를 만들 때 유용하다.
```C++
auto main() -> int
{
    std::vector<int32_t> sVec{1, 2, 3, 4};
    std::for_each(std::cbegin(sVec),
                  std::cend(sVec),
                  [&](auto v){ p(v*v); });

    auto k = []() -> int { // 값 반환
        return 10;
    };
    p(k());
    
    return 0;
}
```

#### 스마트 포인터로 메모리 관리하기
* std::unique_ptr
* std::shared_ptr
* std::weak_ptr

#### tuple을 사용해 다양한 데이터타입 저장하기
```C++
// C++17 Structured Binding Declarations
std::map<int, int> sMap = {{1, 11}, {2, 22}, {3, 33}};
for (const auto [k, v] : sMap)
{
    std::cout << "k:" << k << " v:" << v << std::endl;
}
```
