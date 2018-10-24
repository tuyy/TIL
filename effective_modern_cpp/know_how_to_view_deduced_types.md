## Item 4: Know how to view deduced types.
* 궁극적으로 볼 때, 형식 연역 규칙들을 제대로 숙지하는 것보다 나은게 없다.

### deduction된 형식을 확인할 수 있는 3가지 방법?
1. IDE에서 지원
2. 컴파일 시점
```C++
template <typename T>
class TD;

auto k = ..;
TD<decltype(k)> td; // compile error 발생, k의 타입을 알 수 있다.
```
3. 실행 시점
```C++
std::cout << typeid(x).name() << std::endl;  // 출력으로 확인, 그러나 정확하진 않다.
```

### 참고
* Effective Modern C++ by Scott Meyers
