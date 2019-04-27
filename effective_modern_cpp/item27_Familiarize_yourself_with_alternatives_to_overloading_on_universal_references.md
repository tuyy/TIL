보편 참조에 대한 중복적재는 자유 함수에서도, 멤버 함수(특히 생성자)에서도 다양한 문제를 일으켰다.

이를 대비해 사용할 수 있는 기법들을 확인해보자.

1. 중복적재를 포기한다.
 중복적재를 포기하면서 포기해야되는 성능과 효율은 아깝지 않은가?

2. const T&의 매개변수를 사용한다.
 하지만, 이러한 설계의 단점은 우리가 원하는 만큼 효율적이지 않다는 것이다.

3. 값 전달 방식의 매개변수를 사용한다.
 복잡도는 높지 않고, 성능을 높일 수 있는 방법이지만, 지고간에 반할 수 있다.

4. 꼬리표 배분을 사용한다.
 사실 const 왼값 참조 전달이나 값 전달은 완벽 전달을 지원하지 않는다. 보편 참조를 사용하려는 이유가 완벽 전달이라면, 보편 참조 말고는 다른 대안이 없다. 이때 가능한 방식은 꼬리표 배분 방식이다.

template <typename T>
void logAndAdd(T&& name) {  // 배분자 함수 (dispatcher 함수)
    logAndAddImpl(
        std::forward<T>(name),
        std::is_integral<typename std::remove_reference<T>::type>()
    );
}
=====================================================
template <typename T>
void logAndAddImpl(T&& name, std::false_type) {  //  비정수 인수를 받는 버전
    auto now = std::chromo::system_clock::now();
    log(now, “logAndAdd”);
    names.emplace(std::forward<T>(name));
}

std::string nameFromIdx(int idx);
void logAndAddImpl(int idx, std::true_type) {  // 정수 인수를 받는 버전
    logAndAdd(nameFormIdx(idx));
}

결과적으로, 중복적재의 선택은 전적으로 그 꼬리표 매개변수가 결정한다. 따라서 보편참조 매개변수가 항상 주어진 인수에 대한 정확한 부합을 산출한다는 사실은 더 이상 중요하지 않다.

5. 보편 참조를 받는 템플릿을 제한한다.
 클래스에 보편 참조를 받은 생성자를 추가하면 비const 왼값을 복사할 때 그 생성자가(복사 생성자가 아니라) 호출된다. 보편 참조 매개변수를 포함하는 함수 템플릿이 중복적재 해소의 후보가 되는 조건들을 적절히 제한할 수 있는 또 다른 기법이 있다.

auto&&  <- 왼값&왼값 참조를 대입하면 최종적으로 왼값 참조로 연역된다.
auto&&  <- 오른값을 대입하면 최종적으로 오른값으로 연역된다.

class Person {
public:
    template <
        typename T,
        typename = std::enable_if_t<
            !std::is_base_of<Person, std::decay_t<T>>::value  // Person type이 아니고, (파생클래스도)
            &&
            !std::is_integral<std::remove_reference_t<T>>::value  // 정수가 아니고,
        >
     >
     explicit Person(T&& n)
         : name(std::forward<T>(n))
     { … }

     explicit Person(int idx)
         : name(nameFromIdx(idx))
     { … }

private:
     std::string name;
};

6. 절충점들
 완벽 전달이 더 효율적이라는 점은 하나의 규칙이다. 단지 선언된 매개변수 형식을 만족하기 위해 임시 객체를 생성하는 비효율성이 없기 때문이다.

물론, 완벽 전달의 단점은 완벽 전달이 불가능한 인수들이 있을 수 있고, 클라이언트가 유효하지 않은 인수를 전달했을 때 나오는 오류 메시지가 난해하다는 점이다.

그나마 괜찮은 컴파일 에러를 보는 방법!
static_assert(…); 를 사용하여 생성가능한지 미리 assert 한다.

static_assert(std::is_constructible<std::string, T)::value,
                     “Parammeter n can’t be used to construct a std::string”);


기억해 둘 사항들
1) 보편 참조와 중복적재의 조합에 대한 대안으로는 구별되는 함수 이름 사용, 매개변수를 const에 대한 왼값 참조로 전달, 매개변수를 값으로 전달, 꼬리표 배분 사용 등이 있다.
2) std::enable_if를 이용해서 템플릿의 인스턴스화를 제한함으로써 보편 참조와 중복적재를 함께 사용할 수 있다. std::enable_if는 컴파일러가 보편 참조 중복적재를 사용하는 조건을 프로그래머가 직접 제어하는 용도로 쓰인다.
3) 보편 참조 매개변수는 효율성 면에서 장점인 경우가 많지만, 대체로 사용성 면에서는 단점이 된다.
