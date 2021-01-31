## Item 17: Understand special member function generation.
### special member function은 C++이 스스로 작성하는 default member function을 가리킨다.
* 기본 생성자, 소멸자, 복사 생성자, 복사 배정 연산자 + C++11 이후 이동 생성자, 이동 배정 연산자
    * 기본적으로 이 함수들을 사용하는 클라이언트 코드가 존재할 때에만 작성된다.
    * 암묵적 public, inline이며 가상 소멸자가 있는 부모 클래스를 상속하는 자식 클래스의 소멸자를 제외하고 모두 nonvirtual이다.
* 이동 연산은 사실 요청에 가깝다. 실제로 일어난다는 보장은 없다.
    * 만일 이동 연산을 지원하지 않으면 복사 연산이 수행된다.
* 복잡해보이는 special member function 생성 규칙
    * 두 복사 (배정)생성자 중 하나를 직접 작성하면, 나머지 하나는 필요시 컴파일러가 생성한다.
    * 그러나 두 이동 (배정)생성자 중 하나를 직접 작성하면, 나머지 하나는 컴파일러가 생성하지 않는다.
    * 또한 복사 (배정)생성자 중 하나를 직접 작성하면, 이동 (배정)생성자는 컴파일러가 생성하지 않는다.
    * 반대로 이동 (배정)생성자 중 하나를 직접 작성하면, 복사 (배정)생성자는 컴파일러가 생성하지 않는다.
    * 결과적으로, 이동 연산들은 다음 세 조건이 모두 만족될 때에만, 그리고 필요할 때에만, 자동으로 생성된다.
        * 클래스에 그 어떤 복사 연산도 선언되어 있지 않다.
        * 클래스에 그 어떤 이동 연산도 선언되어 있지 않다.
        * 클래스에 소멸자가 선언되어 있지 않다.
* default 키워드의 사용
```C++
// default 키워드를 사용함으로써 작성자의 의도가 분명하게 나타난다.
class Base
{
public:
    virtual ~Base() = default;
    
    Base(Base&&) = default;
    Base& operator=(Base&&) = default;
    
    Base(const Base&) = default;
    Base& operator=(const Base&) = default;
};
```
* 멤버 함수 템플릿이 존재하면 special member function의 자동 작성이 비활성화되지 않는다.
```C++
class Widget
{
    ...
    template <typename T>
    Widget(const T &rhs);
    
    template<typename T>
    Widget& operator=(const T &rhs);
    
    // 컴파일러는 기본 조건만 만족한다면 복사와 이동 연산들을 작성한다.
};
```     
    
### 참고
* Effective Modern C++ by Scott Meyers
