람다(lambda)란? 익명 함수!

capture mode는 오직 람다가 생성된 범위 안에서 보이는,
매개변수를 포함한 static이 아닌 지역 변수에만 적용된다.

std::find_if, std::remove_if, std::count_if 등 알고리즘 함수들,
std::sort, std::nth_element, std::lower_bound 등 알고리즘 함수들,
std::unique_ptr, std::shared_ptr을 위한 커스텀 삭제자 등에 유용하게 사용할 수 있다.

람다 기본 용어
1) 람다 표현식(lambda expression)은 이름 그대로 하나의 표현식

std::find_if(container.begin(), container.end(),
                  [](int val) { return 0 < val && val < 10; });

2) 클로저(closure)는 람다에 의해 만들어진 실행시점 객체이다.
3) 클로저 클래스는 클로저를 만드는데 쓰인 클래스를 말한다.


C++11의 기본 갈무리 모드(default capture mode)는 두 가지로,
하나는 참조에 의한 갈무리 모드이고 또 하나는 값에 의한 갈무리 모드이다.

기본 참조 갈무리 모드의 위험성 (참조가 대상을 잃을 위험이 있다.)
기본 값 갈무리 모드의 위험성 (참조가 대상을 잃는 문제가 있고, 자기 완결적이지 않다.)

filters.emplace_back(
    [=](int value) { return value % divisor == 0; }  // 기본 값 갈무리 모드, 대상을 잃지 않는다.
);

여기선 이러한 방식이 충분할 수 있지만, 기본 값 갈무리가 대상을 잃는 문제에는 특효약이 아니다.
왜냐하면 포인터를 값으로 기본 값 갈무리하면 그 포인터는 람다에 의해 생성된 클로저 안으로 복사되기 때문이다. (대상을 잃을 수 있음)

클래스 멤버 함수에 람다를 사용할 때, 멤버 변수는 this를 통해 갈무리되기 때문에 위험하다!
우회책으로,

void Widget::addFilter() const {
    auto divisorCopy = divisor;  // 멤버 변수를 지역변수에 복사한다.
    filters.emplace_back(
        [divisorCopy](int value) { return value % divisorCopy == 0; }  // 지역변수의 복사본을 갈무리한다.
    );
}

C++14에는 더 나은 방법이 있다. 바로 일반화된 람다 갈무리이다.

void Widget::addFilter() const {
    filters.emplace_back(
        [divisor = divisor](int value)  // divior를 클로저에 복사한다.
        { return value % divisorCopy == 0; }
    );
}

값에 의한 기본 갈무리 모드의 또 다른 단점은 해당 클로저가 자기 완결적이고 클로저 바깥에서 일어나는 자료의 변화로부터 격리되어 있다는 오해를 부를 수 있다는 점이다.

왜냐하면,
람다는 매개변수, 지역 변수 뿐만 아니라 정적 저장소 수명 기간을 가진 객체에도 의존할 수 있기 때문이다.
(전역 범위나 이름 공간 범위에서 정의된 객체와 클래스, 함수, 파일 안에서 static으로 선언된 객체)

void addDivisorFilter() {
    static auto calc1 = computeSomeValue1();
    static auto calc2 = computeSomeValue2();

    static auto divisor = computeDivisor(calc1, calc2);

    filters.emplace_back(
        [=](int value)  // 아무것도 갈무리하지 않는다.
        { return value % divisor == 0; }   // 위의 정적 변수를 지칭한다.
    );

    ++divisor;  // divisor를 수정한다.
}

lambda는 무엇일까

auto mylmbda = [obj](){
    // 로직
}

위에가 아래처럼 변환한다! 결국 그냥 std::function() .. 

class lambda
{
public:
    int32_t operator(void)()
    {
        // .. 입력한 로직
    }
    T obj;  // 캡쳐로 받은것
};


기억해 둘 사항들
1) 기본 참조 갈무리는 참조가 대상을 잃을 위험이 있다.
2) 기본 값 갈무리는 포인터(특히 this)가 대상을 잃을 수 있으며, 람다가 자기 완결적이라는 오해를 부를 수 있다.
