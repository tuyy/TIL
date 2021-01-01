## Item 12: Declare overriding functions override.
### overriding 함수들을 override로 선언하자
* overriding이 일어나려면  필수조건을 만족해야한다.
    * 부모 함수가 virtual function 이어야 한다.
    * 부모 함수와 자식 함수의 virtual function의 이름이 반드시 동일해야한다. (소멸자는 제외)
    * 부모 함수와 자식 함수의 매개변수 형식들이 반드시 동일해야한다.
    * 부모 함수와 자식 함수의 const성이 반드시 동일해야한다.
    * 부모 함수와 자식 함수의 반환 형식과 예외 명세가 반드시 호환되어야 한다.
    * 멤버 함수들의 참조 한정사들이 반드시 동일해야한다. (C++11 이후에 추가된 조건)
    ```C++
    class Widget
    {
    public:
        ...
        void doWork() &; // *this가 왼값일때만 적용된다.
        void doWork() &&; // *this가 오른값일때만 호출된다.
        ...
    };
    
    w.doWork(); // 왼값용 Widget::doWork 호출
    makeWidget().doWork(); // 오른값용 Widget::doWork 호출
    ```
* 위와 같은 조건이 일치하지 않으면 overriding 되지 않고 완전 새로운 함수가 생성된다.
    * 하지만 override 키워드를 사용하면, 부모 함수 중에 해당 조건의 함수가 존재하는지 확인한다!
    
* 추가로 참조 한정사에 대하여
```C++
class Widget
{
public:
    using DataType = std::vector<double>;
    ...
    
    DataType& data() &
    { return values; }
    
    DataType& data() &&
    { return std::move(values); }
    ...
private:
    DataType values;
};

auto vals1 = w.data(); // 오른값 data() 호출
auto vals2 = makeWidget().data(); // 오른값 data() 호출
```

### 참고
* Effective Modern C++ by Scott Meyers
