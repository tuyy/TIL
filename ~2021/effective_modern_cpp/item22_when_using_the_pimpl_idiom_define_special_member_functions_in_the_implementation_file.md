## Item 22: When using the Pimpl Idiom, define special member functions in the implementation file.
### Pimpl 관용구를 사용할 때에는 special member function들을 구현파일에서 정의하라
* 컴파일 시간을 줄이는 방법 : Pimpl idiom("pointer to implementation" idiom)
* std::unique_ptr 형식의 pImpl 포인터를 사용할 때에는 특수 멤버 함수들을 클래스 헤더에 선언하고 구현 파일에서 구현해야한다.
    * 컴파일러가 default로 작성하는 함수 구현들이 사용하기에 적합한 경우에도 마찬가지이다.
    * 삭제자 함수는 delete를 적용하기 전에, 혹시 raw pointer가 불완전한 형식을 가리키지는 않는지 static_assert를 이용해서 점검하기 때문이다.
    ```C++
    // Widget.h
    class Widget
    {
    public:
        Widget();
        ~Widget();

        Widget(Widget &&aRhs);
        Widget& operator=(Widget &&aRhs);
    private:
        struct Impl;
        std::unique_ptr<Imple> pImpl;
    };

    // Widget.cpp
    struct Widget::Impl
    {
        std::string name;
        std::vector<double> data;
        ...
    };

    Widget::Widget()
        : pImpl(std::make_unique<Impl>)
    {
    }

    ~Widget::Widget() = default;
    Widget::Widget(Widget &&aRhs) = default;
    Widget& Widget::operator=(Widget &&aRhs) = default;
    ```
* std::shared_ptr은 상관없다.

### 참고
* Effective Modern C++ by Scott Meyers
