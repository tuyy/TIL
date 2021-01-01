## Item 7: Distinguish between () and {} when creating objects.
### uniform initialization (균일 초기화, 중괄호 초기화)
* 컨테이너의 초기 값을 지정하기
```C++
std::vector<std::string> sVector = {"one", "two", "three"}; // Ok
std::vector<std::string> sVector{"one", "two", "three"}; // Ok
```
* 암묵적으로 narrowing conversion을 방지해준다.
```C++
double x, y, z;
...
int sum1{x + y + z}; // error! double들의 합을 int로 표현하지 못한다.
int sum2(x + y + z); // ok
int sum2 = x + y + z; // ok
```
* most vexing parse(가장 성가신 구문)에서 자유롭다.
```C++
Widget w(); // 함수 선언과 비슷하여 혼동을 줄 수 있다.
Widget w{}; // 명확한 widget 초기화
```
### 단점
* std::initializer_list, 생성자 overloading 사이의 이상한 호출결과가 발생한다.
```C++
class Widget
{
public:
    Widget(int i, bool b) {}
    Widget(int i, double b) {}
    Widget(std::initializer_list<long double> il) {}
};

Widget w1(10, true);   // 1번째 생성자 호출
Widget w2(10, 5.0);  // 2번째 생성자 호출

// std::initializer_list 생성자가 있으면 대부분 std::initializer_list가 포함된 생성자를 호출한다.
Widget w3{10, true};   // 3번째 생성자 호출
Widget w4{10, 5.0};  // 3번째 생성자 호출
```
* 괄호와 중괄호 사이의 vector 생성자 호출 예시
```C++
std::vector<int> v1(10, 20); // 값이 20인 요소 10개짜리 vector가 생성
std::vector<int> v2{10, 20}; // 값이 10, 20인 두 요소를 가진 vector가 생성
```

### 결과적으로
* 객체 생성시 괄호와 중괄호를 세심하게 선택해야하며, 특히 일관된 형태로 사용하자


### 참고
* Effective Modern C++ by Scott Meyers
