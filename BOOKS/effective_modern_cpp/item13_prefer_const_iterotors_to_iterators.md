## Item 13: Prefer const_iterators to iterators.
### 가능한 const를 들이대라는 말은 iterator에도 적용된다.
* C++98 스타일 const_iterator 적용
```C++
typedef std::vector<int>::iterator IterT;
typedef std::vector<int>::const_iterator ConstIterT;

std::vector<int> values;
...
ConstIterT ci = std::find(static_cast<ConstIterT>(values.begin()), static_cast<ConstIterT>(values.end()), 1983);
values.insert(static_cast<IterT>(ci), 1988); // ConstIterT로 컴파일 안될 수 도 있음
```

* C++11 스타일 const_iterator 적용
```C++
std::vector<int> values;
auto ci = std::find(values.cbegin(), values.cend(), 1983);
values.insert(ci, 1988);
```

* 컨테이너 뿐 아니라 일반 배열도 사용할 수 있는 일반화된 iter 사용 예시
```C++
// std::cbegin 류는 C++14부터 제공
template <typename C, typename V>
void findAndInsert(C &container, const V &target, const V&insertVal)
{
    auto it = std::find(std::cbegin(container), std::cend(container), targetVal);
    container.insert(it, insertVal);
}

// C++11 에서 사용 가능한 cbegin 함수 구현
template <typename C>
auto cbegin(const C &container) -> decltype(container.begin())
{
    // container는 const C&이므로 begin호출시 const_iterator 반환
    return container.begin();
}
```

### 참고
* Effective Modern C++ by Scott Meyers
