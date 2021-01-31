C++14에서 가장 고무적인 기능은 일반적 람다, 즉 매개변수 명세에 auto를 사용하는 람다이다.

auto f= [](auto x){ return normalize(x); }

이 람다가 산출하는 클로저 클래스의 함수 호출 연산자는 다음과 같은 모습이다.

class 컴파일러가_만든_어떤_클래스_이름 {
public:
    template <typename T>
    auto operator()(T x) const  // auto 반환 형식을 가진다.
    { return normalize(x); }
};

만약 왼값, 오른값 여부가 필요하다면? auto만으로는 불가능하다.

auto f = [](auto x){ return normalize(std::forward<???>(x)); };

이런 경우, decltype을 사용할 수 있다.

참조축약이 발생해도 (&&, &&)는 &&로 연역된다.

Widget&& && forward(Widget& param) {
    return static_cast<Widget&& &&>(param);  // 참조 축약 적용 전,
}

Widget&& forward(Widget& param) {
    return static_cast<Widget&&>(param);  // 참조 축약 적용 후
}

즉, 오른값 참조 형식으로 std::forward를 인스턴스화한 결과는 비참조 형식으로 인스턴스화한 결과와 같다.

결과적으로, 아래와 같은 코드를 작성하면, 매개변수를 완벽하게 넘길 수 있다.

auto f = [](auto&& x){ return normalize(std::forward<decltype(x)>(x)); };

auto f = [](auto&&... x){ return normalize(std::forward<decltype(x)>(x)...); };


기억해 둘 사항들
1) std::forward를 통해서 전달할 auto&& 매개변수에는 decltype을 사용하라.
