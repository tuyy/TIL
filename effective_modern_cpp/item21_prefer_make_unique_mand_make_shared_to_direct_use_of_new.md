## Item 21: Prefer std::make_unique and std::make_shared to direct use of new.
### 직접 new로 생성하기 보단 std::make_unique, std::make_shared를 선호하라
* 버전 별 지원사항
    * C++11 부터 std::make_shared 지원
    * C++14 부터 std::make_unique 지원
    ```C++
    // C++11에서 make_unique 사용하기
    template <typename T, typename ...Ts>
    std::unique_ptr<T> make_unique(Ts&&... aArgs)
    {
        return new unique_ptr<T>(std::forward<Ts>(aArgs)...);
    }
    ```
* make 류를 사용하면 형식 반복이 되지 않는다.
```C++
auto upw1(std::make_unique<Widget>());
std::unique_ptr<Widget> upw2(new Widget);  // 형식 중복!

auto spw1(std::make_shared<Widget>());
std::shared_ptr<Widget> spw2(new Widget);  // 형식 중복!
```
* 예외 안전성이 생긴다.
```C++
void processWidget(std::shared_ptr<Widget> spw, int priority);

// 호출 순서가 보장되지 않으므로 메모리 누수 가능성을 내포한다. 1) new Widget 2) computePriority()시 에러 발생
processWidget(new Widget, computePriority());

// 자원 누수의 위험이 없다.
processWidget(std::make_shared<Widget>(), computePriority());
```
* std::make_shared 단 한번으로, managed pointer와 control block을 할당한다. (원래는 두 번)
    * 그렇기 때문에 shared_ptr 참조횟수와 weak_ptr 참조 횟수가 모두 0이되어야 해당 managed pointer가 delete 된다.
    * 그러나 new로 생성한 경우 shared_ptr 참조횟수가 0이되는 경우 즉시 managed pointer가 삭제된다.
* make 함수들 중에는 커스텀 삭제자를 지정할 수 있는 것이 없다.
* std::initializer_list를 받는 생성자와 받지 않는 생성자를 모두 가진 형식의 객체를 생성할때 문제가 존재할 수 있다.
    * 중괄호를 감싸면 항상 std::initializer_list를 받는 생성자가 호출되고 괄호를 감싸면 받지 않는 생성자가 호출된다.
    ``` C++
    // 이런 경우 직접 std::initializer_list 객체를 생성해서 넘겨라
    auto initList = {10, 20};
    auto spv = std::make_shared<std::vector<int>>(initList);
    ```
* std::make_shared 함수에서 커스텀 operator new와 operator delete를 가진 객체를 생성할 때 make 함수를 쓰지 말자
    * 보통 커스텀 new, delete를 사용하는 경우 객체의 사이즈가 sizeof(Widget)과 다를 수 있다.
    * 하지만 make류 함수는 sizeof(Widget) + control block을 합친 크기로 동적할당하기 때문에 문제가 있을 수 있다.
* 커스텀 삭제자를 지정이 필요한 경우 반드시 new를 써야하는데, 이때 예외 안전성 확보하기
```C++
void processWidget(std::shared_ptr<Widget> spw, int priority);

void cusDel();

// 자원 누수의 위험이 있다.
processWidget(std::shared_ptr<Widget>(new Widget, cusDel), computePriority());

// 자원 누수의 위험이 없다. (shared_ptr의 생성을 분리하자)
std::shared_ptr<Widget> spw(new Widget, cusDel);
processWidget(std::move(spw), computePriority());
```

### 참고
* Effective Modern C++ by Scott Meyers
