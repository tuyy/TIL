## Item 25: Use std::move on rvalue references, std::forward on universal references.
### 오른값 참조일땐 std::move, 보편참조일땐 std::forward를 사용하라
#### std::move
* std::move에 std::forward를 적용해도 동작엔 특이사항이 없다. 하지만 소스 코드가 장황해지고 관용구에서 벗어난 문법이 되어 사용하지 않는다.
* 가장 주의해야할 점은, 이동 후 해당 값은 미정의 값이 되므로 참조의 마지막에 사용하여야 한다.
```C++
class Widget
{
public:
    Widget(Widget&& rhs)
        : name(std::move(rhs.name)),
          p(std::move(rhs.p))
        {…}
    ….
private:
    std::string name;
    std::shared_ptr<SomeDataStructure> p;
};
```

#### std::forward
```C++
class Widget
{
public:
    template<class T>
    void setName(T&& newName)
    { name = std::forward<T>(newName); }
};
```

#### 반환값 최적화에서 (RVO: 불필요한 복사를 하지 않고 컴파일러가 최적화하여 return하는 최적화 기법)
* 값을 반환하는 함수의 지역변수에 std::move, std::forward를 사용하면 안된다.
* 왜냐하면 RVO가 일어나지 않을 수 있기 때문이다.

### 참고
* Effective Modern C++ by Scott Meyers
