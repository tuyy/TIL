## Item 10: Prefer scoped enums to unscoped enums.
### enum 보단 enum class를 선호하라
* enum class는 namespace 오염을 줄여준다.
```C++
// unscoped enum
enum Color {black, white, red};
auto white = false; // error! 이미 white가 선언되어 있다.

// scoped enum
enum class Color {black, white, red};
auto white = false; // Ok
Color c = white; // error! 이 범위에 white라는 enum이 없음.
Color c = Color::white; // Ok
auto c = Color::white; // Ok
```

* enum class type이 enum보다 훨씬 강력하게 적용된다.
```C++
// unscoped enum
enum Color {black, white, red};
std::vector<std::size_t> primeFactors(std::size_t x); // x의 소인수를 돌려주는 함수
Color c = red;
...
if (c < 14.5)
{
    auto factors = primeFactors(c); // Ok
}

// scoped enum
enum class Color {black, white, red};
Color c = Color::red;
...
...
if (c < 14.5) // error! 컴파일하려면 static_cast<double>(c) 처럼 캐스팅해야한다.
{
    auto factors = primeFactors(c); // error! c는 size_t가 아니라 Color type이다.
}
```
* scoped enum은 전방선언(forward declaration)이 가능하다.
    * 전방선언을하면 전체 컴파일을 막을 수 있다.
    * unscoped enum에 type만 지정하면 전방선언이 가능하다.
    ```C++
    enum Color: std::uint8_t;
    ```
* scoped enum의 default type은 int이고 unscoped enum은 default type이 없다.
* unscoped enum의 유용한 경우가 있다.
```C++
using UserInfo = std::tuple<std::string, std::string, std::size_t>;

enum UserInfoFields {uiName, uiEmail, uiReputation};
UserInfo uInfo;
...
auto val = std::get<uiEmail>(uInfo); // 더 명시적이다! enum class를 사용하면 캐스팅을 해주어야한다.
```

### 참고
* Effective Modern C++ by Scott Meyers
