## Item 9: Prefer alias declarations to typedefs.
### typedef 보다 alias declarations(using)을 사용하자
* 더 직관적이다.
```C++
// C++98 typedef
typedef std::unique_ptr<std::unordered_map<std::string, std::string>> UPtrMapSS;

// C++11 alias declarations
using UPtrMappSS = std::unique_ptr<std::unordered_map<std::string, std::string>>;
```
* 함수 포인터를 사용할 때, 보다 더 명확하다
```C++
typedef void (*FP)(int, const std::string&); // typedef
using FP = void(*)(int, const std::string&); // alias declarations
```
* alias templates (별칭 템플릿)이 가능하다.
```C++
template <typename T>
using MyAllocList = std::list<T, MyAlloc<T>>;

// alias templates
MyAllocList<Widget> lw;

template <typename T>
struct MyAllocList
{
    typedef std::list<T, MyAlloc<T>> type;
};

// typedef
MyAllocList<Widget>::type ltw;
```
* 템플릿 매개변수를 사용하는 클래스 내부에서 typedef된 list를 생성하면 더 복잡해진다.
```C++

template <typename T>
struct MyAllocList
{
    typedef std::list<T, MyAlloc<T>> type;
};

template <typename T>
class Widget
{
public:
    typename MyAllocList<T>::type list; // typedef 앞에 typename을 반드시 붙여야한다.
    ...
};

template <typename T>
using MyAllocList = std::list<T, MyAlloc<T>>; // alias templates

template <typename T>
class Widget
{
public:
    MyAllocList<T> list; // typename과 ::type이 없음
};
```

* C++98 이전의 type_traits와 C++11,14 이후 type_traits 비교
```C++
// C++98 typedef를 사용하는 경우 형식은 항상 std::변환<T>::type 이다.
std::remove_const<T>::type;
std::remove_reference<T>::type;
std::add_lvalue_reference<T>::type;

// C++14 version
std::remove_const_t<T>;
std::remove_reference_t<T>;
std::add_lvalue_reference_t<T>;

// C++11에서 C++14를 흉내내기도 쉽다.
using remove_const_t = typename remove_const<T>::type;
...
```

### 참고
* Effective Modern C++ by Scott Meyers
