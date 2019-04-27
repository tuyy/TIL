C++11의 주요 기능으로 강조되는 것 중 하나가 완벽 전달이다.

완벽전달(perfect forwarding)은 말 그대로 한 함수가 자신의 인수들을 다른 함수에 넘겨주는 것을 뜻한다.
이때 목표는 둘째 함수가 첫 함수가 받았던 것과 동일한 객체들을 받게하는 것이다.

 물론 값 전달 방식의 매개변수로는 이런 목표를 달성할 수 없다. 그런 매개변수는 원래의 호출자가 넘겨준 인수의 복사본이기 때문이다.
마찬가지의 이유로 포인터 매개변수 역시 탈락이다.

 뿐만 아니라, 완벽 전달은 단순히 객체들을 전달하는 것만이 아니라, 그 객체들의 주요 특징, 즉 그 형식, 왼값 또는 오른값 여부, const나 volatile 여부까지도 전달하는 것을 말한다. 이때 보편참조가 쓰인다. 보편 참조 매개변수만이 전달된 인수의 왼값, 오른값에 대한 정보를 부호화하는 것이 가능하기 때문이다.

template <typename T>
void fwd(T&& param)  // 임의의 인수를 받는다.
{
    f(std::forward<T>(param));  // 그 인수를 f에 전달한다.
}

위 전달함수는 필연적으로 일반적(generic) 함수이다.
이를 연장하면, 그냥 템플릿 함수가 아니라 아래와 같이 가변인수 템플릿이어야 할 것이다.

template <typename… Ts>
void fwd(Ts&&… params)  // 임의의 인수들을 받는다.
{
    f(std::forward<Ts>(params)…);  // 그 인수들을 f에 전달한다.
}

이것이 전달 함수의 공통적인 형태이다. 특히 표준 컨테이너의 emplace류 함수들과 똑똑한 포인터 팩터리 함수 std::make_shared, std::make_uniquedptj 이런 형태를 볼 수 있다.

이러한 대상 함수 f와 전달 함수 fwd가 있다고 할 때, 만일 어떤 인수로 f를 호출했을 때 일어나는 일과 같은 인수로 fwd를 호출했을 때 일어나는 일이 다르다면 완벽 전달은 실패한 것이다.

f(표현식);  // 이 호출이 하는 일과
fwd(표현식);  // 이 호출이 하는 일이 다르다면, fwd는 표현식을 f에 완벽하게 전달하지 못한 것이다.


이러한 실패로 이어지는 종류의 인수들이 존재한다. 이러한 실패를 알고, 우회책에 대해 미리 아는 것이 매우 중요하다.
그럼 완벽하게 전달할 수 없는 인수들의 각 종류를 차례로 살펴보자.

1. 중괄호 초기치

void f(const std::vector<int> &v);

f({1, 2, 3});  // 암묵적으로 std::vector<int>로 변환된다.
fwd({1, 2, 3);  // 오류! 컴파일되지 않는다!

fwd에 전달된 인수들의 형식을 연역하고, 연역된 형식들을 f의 매개변수 선언들과 비교한다.
이때 다음 두 조건 중 하나라도 만족되면 완벽 전달이 실패한다.

1) fwd의 매개변수들 중 하나 이상에 대해 컴파일러가 형식을 연역하지 못한다.
2) fwd의 매개변수들 중 하나 이상에 대해 컴파일러가 형식을 잘못 연역한다.

앞에 나온 fwd({1, 2,3}) 호출에서 문제는 std::initializer_list가 될 수 없는 형태로 선언된 함수 템플릿 매개변수에 중괄호 초기치를 넘겨준다는 것이다.

우회책으로, 따로 선언해서 넘긴다.
auto il = {1, 2, 3};
fwd(il);  // 정상적으로 컴파일된다.


2. 널 포인터를 뜻하는 0 또는 NULL
 0이나 NULL을 널 포인터로서 템플릿에 넘기려하면 템플릿은 정수형식으로 연역하기 때문에 문제가 생긴다.
결과적으로 0, NULL은 널 포인터로서 완벽하게 전달되지 못하기 때문에, nullptr를 사용하면 된다.


3. 선언만 된 정수 static const 및 constexpr 자료 멤버
 일반적인 규칙으로, 정수 static const 자료 멤버와 정수 static constexpr 자료멤버는 클래스 안에서 정의할 필요가 없고, 선언만 하면 된다. 그런 멤버의 값에 대해 컴파일러가 const 전파를 적용해서 그런 멤버의 값을 위한 메모리를 따로 마련할 필요가 없어지기 때문이다.

class Widget {
public:
    static constexpr std::size_t MinVals = 28;
    …
};

std::vector<int> widgetData;
widgetData.reserve(Widget::MinVals);  // MinVals를 사용.

 위와 같은 경우 정의가 없어도 잘 동작하긴 한다.(원래는 정의가 필수이다.)
하지만 만일 어떤 코드에서 MinVals의 주소를 취한다면, MinVals를 위한 저장소가 필요해지며, 그러면 이 코드는 컴파일되지만 MinVals의 정의가 없어서 링크에 실패한다.

void f(std::size_t val);  // 이 점을 염두해두고 값으로 선언

f(Widget::MinVals);
fwd(Widget::MinVals);  // 오류! 링크에 실패한다.

소스 코드에 MinVals의 주소를 취하는 부분이 없다고 해도, fwd의 매개변수는 보편 참조이며, 컴파일러가 산출한 코드에서 참조는 포인터처럼 취급되는 것이 보통이기 때문에, 문제가 발생한다.

이러한 이유로, 정수 static const 및 constexpr 자료멤버를 참조로 전달하려면 그 멤버를 정의할 필요가 있다.

constexpr std::size_t Widget::MinVals;  / Widget의 .cpp 파일에서


4. 중복적재된 함수 이름과 템플릿 이름

void f(int (*pf)(int));  // pf는 processing function

void f(int pf(int));  // 위와 동일하다.

int processVal(int value);
int processVal(int value, int priority);

f(processVal);  // OK.

컴파일러는 두 함수 중 어떤 것이 필요한지 매개변수 형식을 보고 알 수 있다. 그래서 컴파일러는 int를 받는 processVal을 선택해서 그 함수의 주소를 f에 넘겨준다.

하지만, 함수 템플릿인 fwd에는 그러한 호출에 필요한 형식에 관한 정보가 전혀 없다. 따라서 컴파일러는 어떤 중복적재를 선택해야 할지 결정하지 못한다.

fwd(processVal);  // 오류! 어떤 processVal인지?

processVal 자체에는 형식이 없으며, 형식이 없으면 형식 연역도 없다.

우회책으로, 전달하고자 하는 중복적재나 템플릿 인스턴스를 명시적으로 지정하면 된다.

using ProcessFuncType = int (*)(int);  // typedef들을 만든다.
ProcessFuncType processValPtr = processVal;  // processVal에 필요한 시그니처를 명시한다.

fwd(processValPtr);  // Ok
fwd(static_cast<ProcessFuncType>(workOnVal);  // Ok

물론 이를 위해서는 fwd가 전달하는 함수 포인터의 형식을 알고 있어야한다.
다행히, 완벽 전달 함수의 문서화에 그 형식에 대한 정보가 나와 있으리라고 가정하는 것이 비합리적인 생각은 아니다.


5. 비트필드

struct IPv4Header {
    std::uint32_t version:4,
                         IHL:4,
                         DSCP:6,
                         ECN:2,
                         totalLength:16;
};

void f(std::size_t sz);  // 호출할 함수
IPv4Header h;
…
f(h.totalLength);  // Ok
fwd(h.totalLength)  // 여기서 전달되면 에러가 발생한다!

문제는 fwd의 매개변수가 참조이고, h.totalLength는 비const 비트필드라는 점이다. 그게 뭐가 문제인가 싶겠지만,
C++ 표준은 둘의 조합에 대해 “비const 참조는 절대로 비트필드에 묶이지 않아야 한다.” 라고 명확하게 선고하기 때문이다.

이는 임의의 비트들을 가리키는 포인터를 생성하는 방법은 없으며,(C++에서 직접 가리킬 수 있는 가장 작은 것은 char이다.)
따라서 참조를 직접적으로 비트들에 묶는 방법도 없다.

우회책으로, 해당 복사본을 직접 생성해서 그 복사본으로 전달 함수를 호출하면 된다.

auto length = static_cast<std::uint16_t>(h.totalLength);
fwd(length);  // 복사본을 전달한다.


기억해 둘 사항들
1) 완벽 전달은 템플릿 형식 연역이 실패하거나 틀린 형식을 연역했을 때 실패한다.
2) 인수가 중괄호 초기치이거나 0 또는 NULL로 표현된 널 포인터, 선언만 된 정수 static const 및 constexpr 자료 멤버, 템플릿 및 중복적재된 함수 이름, 비트필드이면 완벽 전달이 실패한다.
