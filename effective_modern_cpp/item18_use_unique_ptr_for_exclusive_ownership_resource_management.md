## Item 18: Use std::unique_ptr for exclusive-ownership resource management.
### 소유권 독점 자원 관리는 std::unique_ptr를 사용하라
* 거의 raw pointer 크기의 아주 작고 빠른 이동 전용 스마트 포인터이다.
* std::unique_ptr은 복사는 불가능, 이동시 기존 객체는 nullptr이 된다.
* 주로 팩토리 함수의 반환형식이나 pimp 구현시 사용한다.
* 템플릿 형식으로 삭제자를 지정할 수 있다.
```C++
auto delInvmt = [](Investment *pInvestment)
{
    makeLogEntry(pInvestment);
    delete pInvestment;
};

template <typename ...Ts>
std::unique_ptr<Investment, decltype(delInvmt)> makeInvestment(Ts&&... aArgs)
{
    std::unique_ptr<Investment, decltype(delInvmt)> pInv(nullptr, delInvmt);
    if (...)
    {
        pInv.reset(new Stock(std::forward<Ts>(aArgs)...));    
    }
    else if (...)
    {
        pInv.reset(new Bond(std::forward<Ts>(aArgs)...));
    }
    else
    {
        ....
    }
    return pInv;
}
```
* 삭제자는 함수 객체, 함수 포인터와 람다 표현식으로 구현할 수 있는데, 보통 람다 표현식의 크기가 더 작다.
    * 람다 표현식을 선호하자
* 배열 선언이 가능하다. 배열 포인터로 선언하는 경우 unique_ptr은 적절하게 delete[]를 호출한다.
    * 하지만 내장 배열보다는 std::array, std::vector, std::string이 거의 항상 더 나은 선택이다.
    ```C++
    std::unique_ptr<char[]> str = new char[100];
    ```
* std::shared_ptr로의 변환이 매우 쉽고 효율적이다.

### 참고
* Effective Modern C++ by Scott Meyers
