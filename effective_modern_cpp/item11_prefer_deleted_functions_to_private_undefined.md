## Item 11: Prefer deleted functions to private undefined ones.
### 정의되지 않은 private 함수보다 delete keyword를 선호하라
* 간결하고 명확하다.
```C++
// C++98 정의되지 않은 private 함수 선언
class Widget
{
public:
    ...
private:
    Widget(const Widget&); // not defined
    Widget& operator=(const Widget&); // not defined
    ...
};

// C++11 delete 키워드 사용
class Widget
{
public:
    Widget(const Widget&) = delete;
    Widget& operator=(const Widget&) = delete;
    ...
};
```
* delete 키워드는 public 범위에서 사용하는 것이 관례이다.
    * 컴파일러는 접근성을 점검한 후 삭제 여부를 확인하기 때문이다. (명확한 컴파일 오류 메시지를 위하여)
* delete 키워드는 모든 함수에 적용할 수 있다. (private 접근 제어는 클래스 내에서만 사용 가능하다.)
```C++
bool isLucky(int number);
if (isLucky('a')){...}
if (isLucky(true)){...}
if (isLucky(1.5)){...}

// 행운의 번호가 반드시 정수여야 한다면, 위와 같은 경우 컴파일을 막아야한다.
bool isLucky(char) = delete;
bool isLucky(bool) = delete;
bool isLucky(double) = delete;
```
* 템플릿 인스턴스에 대한 삭제도 가능하다.
```C++
template <typename T>
void processPointer(T *ptr);

// 두 가지 특별한 포인터가 호출되지 않게 하려면,
template<>
void processPointer<void>(void*) = delete;

template<>
void processPointer<char>(char*) = delete;

template<>
void processPointer<const void>(const void*) = delete;

template<>
void processPointer<const char>(const char*) = delete;
```

### 참고
* Effective Modern C++ by Scott Meyers
