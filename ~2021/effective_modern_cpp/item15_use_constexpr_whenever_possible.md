## Item 15:  Use constexpr whenever possible.
### 가능하면 항상 constexpr을 사용하라
* 개념적으로 constexpr은 어떠한 값이 상수일 뿐만 아니라 컴파일 시점에서 알려진다는 점을 나타낸다.
    * 하지만 함수에 적용될때는 상황이 좀 더 미묘해진다.
    * 결과적으로 constexpr의 결과가 반드시 const는 아닐 수 있고, 반드시 컴파일 시점에 알려진다는 보장은 없다.
* 컴파일 시점에 알려지는 값들에는 특별한 권한이 있다.
    * 읽기 전용 메모리에 배치될 수 있다. (임베디드 개발자에게 중요한 특성이다.)
    * 정수 상수 표현식(integral constant expression)이 요구되는 문맥에서 사용할 수 있다.
    ```C++
    int sz;
    ...
    constexpr auto arraySize1 = sz; // error!
    std::array<int, sz> data1; // error!
    
    constexpr auto arraySize2 = 10; // Ok, this is compile time constant
    std::array<int, arraySize2> data2; // Ok, arraySize2 is contexpr.
    
    const auto arraySize = sz; // Ok
    std::array<int, arraySize> data; // error. arraySize는 컴파일 상수가 아님
    ```
* constexpr 함수의 올바른 관점
    * 컴파일 시점 상수를 요구하는 문맥에 constexpr 함수를 사용할 수 있다.
    * 컴파일 시점에서 알려지지 않는 하나 이상의 값들로 constexpr 함수를 호출하면 함수는 보통의 함수처럼 동작한다.
    ```C++
    constexpr int pow(int base, int exp) noexcept
    {
        ...
    }
    
    constexpr auto numConds = 5;
    std::array<int, pow(3, numConds> results; //Ok
    ```
    * constexpr 함수는 반드시 리터럴 형식(literal type)들을 받고 돌려주어야 한다.
        * 리터럴 형식은 컴파일 도중에 값을 결정할 수 있는 형식이기 때문이다. (C++11 이후부턴 void를 제외한 모든 내장 type)
        * 생성자와 적절한 멤버 함수들이 constexpr인 사용자 형식도 리터럴 형식이 될 수 있다.
        ```C++
        class Point
        {
        public:
            constexpr Point(int a, int b) noexcept
                : x(a), y(b)
            {
            }
            
            constexpr int getX() const noexcept {return x;}
            constexpr int getY() const noexcept {return y;}
            ... // C++11은 setX, setY 불가능(오직 반환값이 리터럴이어야 하기 때문이다.)하지만 C++14부턴 가능
        private:
            int x, y;
        };
        
        constexpr Point p1(1, 2); // Ok. constexpr 생성자가 컴파일 시점에서 실행됨
        constexpr Point p2(3, 4); // Ok. constexpr 생성자가 컴파일 시점에서 실행됨
        
        constexpr Point midpoint(const Point &p1, const Point &p2) noexcept
        {
            return {(p1.getX() + p2.getX()) / 2, (p1.getY() + p2.getY()) / 2};
        }
        
        constexpr auto mid = midpoint(p1, p2); // Ok.
        ```
        
* C++11과 C++14에서의 constexpr 함수 사용
    * C++11에서 constexpr 함수는 실행 가능 문장이 많아야 하나이어야 한다. (보통 return문)
    * C++14에선 아래와 같은 구현도 가능하다.
    ```C++
    constexpr int pow(int base, int exp) noexcept //C++14
    {
        auto result = 1;
        for (int i = 0; i < exp; ++i) result *=base;
        return result;
    }
    ```

### 참고
* Effective Modern C++ by Scott Meyers
