인수가 템플릿에 전달되었을 때 템플릿 매개변수에 대해 연역된 형식에는 그 인수가 왼값인지 아니면 오른값인지에 대한 정보가 부호화되어 있다. 이는 오직 인수가 보편 참조 매개변수를 초기화하는 데 쓰일 때에만 부호화 된다.

template <typename T>
void func(T&& param);

이러한 부호화 매커니즘은 간단하다.

왼값 인수가 전달되면, T는 왼값 참조
오른값 인수가 전달되면, T는 비참조(non-reference)

C++에선 참조에 대한 참조는 위법이지만, 보편 참조시엔 적법하다. 
하지만, 새로운 규칙이 있다.

만일 두 참조 중 하나라도 왼값 참조이면 결과는 왼값 참조이다. 그렇지 않으면(즉, 다 오른값 참조이면) 결과는 오른값 참조이다.

이는 참조 축약(reference collapsing)이라 한다.

template <tpyename T>
void f(T&& fParam) {
    …
    someFunc(std::forward<T>(fParam);
}
// std::forward<T>(…) 는 전달된 인수가 오른값이라는 점(T가 비참조 형식)이 부호화되어 있으면 fParam(왼값)을 오른값으로 캐스팅한다.

// 왼값
Widget& &&forard(typename remove_reference<Widget&>::type& param)
{
    return static_cast<Widget& &&>(param);  // 두 참조중 하나라도 왼값이면(참조 축약 발생), 왼값 참조!
}

// 오른값
Widget&& &&forard(typename remove_reference<Widget>::type& param)
{
    return static_cast<Widget&&>(param);  // 오른값 참조! (참조축약 발생하지 않음)
}


참조 축약이 일어나는 문맥은 네 가지이다.

1) 템플릿 인스턴스 화

2) auto 변수에 대한 형식 연역

3) typedefdhk quf칭 선언의 지정 및 사용
template <typename T>
class Widget {
public:
    typedef T&& RvalueRefToT;
    ...
};

4) decltype의 사용


기억해 둘 사항들
1) 참조 축약은 템플릿 인스턴스화, auto 형식 연역, typedef와 별칭 선언의 지정 및 사용, decltype의 지정 및 사용이라는 네 가지 문맥에서 일어난다.
2) 컴파일러가 참조 축약 문맥에서 참조에 대한 참조를 만들어 내면, 그 결과는 하나의 참조가 된다. 원래의 두 참조 중 하나라도 왼값 참조이면 결과는 왼값 참조이고, 그렇지 않으면 오른값 참조이다.
3) 형식 연역이 왼값과 오른값을 구분하는 문맥과 참조 축약이 일어나는 문맥에서 보편 참조는 오른값 참조이다.
