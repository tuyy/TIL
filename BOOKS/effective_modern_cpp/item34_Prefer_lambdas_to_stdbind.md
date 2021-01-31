C++11에서는 람다가 거의 항상 std::bind보다 나은 선택이다.
C++14에서는 람다가 확고하게 std::bind보다 나은 선택이다.

std::bind보다 람다를 선호할 가장 중요한 이유는 람다가 가독성이 훨씬 더 좋다는 것이다.

// C++14
auto betweenL = [lowVal, highVal](const auto &val) { return lowVal <= val && highVal >= val; };

// std::bind
using namespace std::placeholders;
auto betweenB = std::bind(
    std::logical_and<>(),
    std::bind(std::less_equal<>(), lowVal, _1),
    std::bind(std::less_equal<>(), _1, highVal)
);

어떤 경우든 이제는 람다 버전이 더 짧을 뿐 아니라 이해하기도 쉽고 유지보수 하기도 쉽다.

추가적으로, 명시적이지 않은 문제도 있다.

Widget w;
using namespace std::placeholders;
auto compressRateB = std::bind(compress, w, _1);  // w는 값으로 저장될까? 참조로서 저장될까?

답은 값으로 전달된다. 이 답을 얻으려면 std::bind의 작동 방식을 알고 있어야 한다.

// w는 값으로 갈무리되고 lev는 값으로 전달된다. 명백하다!
auto compressRateL = [w](CompLevel lev) { return compress(w, lev); };


그러나, C++11에서는 다음과 같이 제한된 두 경우라면 std::bind 사용을 정당화할 수 있다.
1) 이동 갈무리 지원
2) 다형적 함수 객체

class PolyWidget {
public:
    template <typename T>
    void operator()(const T& param) const;
};

std::bind로 표현하면 다음과 같이 하면 된다.

PolyWidget pw;
auto boundPW = std::bind(pw, _1);

하지만 C++11 람다로는 이런 일이 불가능하다. (어쩔 수 없이 std::bind를 사용해야한다.)
그러나 C++14는 auto 매개변수를 가진 람다로 간단히 구현이 가능하다.

auto boundPW = [pw](const auto &param){ pw(param); };

기억해 둘 사항들
1) std::bind를 사용하는 것보다 람다가 더 읽기 쉽고 표현력이 좋다. 그리고 더 효율적일 수 있다.
2) C++14가 아닌 C++11에서는 이동 갈무리를 구현하거나 객체를 템플릿화된 함수 호출 연산자에 묶으려 할 때 std::bind가 유용할 수 있다.
