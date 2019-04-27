* evernote -> github 이전

아래는 보편 참조(universal ref)의 예이다.
(T&& 대신 const string& 을 받을 때와 비교해보자)
std::multiset<string> names;

template<typename T>
void logAndAdd(T&& name)
{
    auto now = std::chromo::system_clock::now();
    log(now, “log”);
    names.emplace(std::forward<T>(name));
}

string petName = “hello”;
logAndAdd(petName);            // (1) lvalue

logAndAdd(string(“zone"));     // (2) rvalue
logAndAdd(“Patty dog”);        // (3) const char*
올바르게 보편 참조로써 동작한다.
(1) lvalue으로 names에 복사
(2) rvalue으로 names에 이동
(3) 임시 string 객체에 대한 복사가 아니라 multiset안에 string을 생성


위 코드에서 index를 받아서 이름을 조회하여 넣는 경우 구현은 어떻게 해야할까?
일단 직관적으로 아래와 같이 중복적재(overloading)를 생각할 수 있다.

string nameFromIdx(int idx); // idx에 해당하는 이름을 반환

void logAndAdd(int idx)
{
    auto now = std::chromo::system_clock::now();
    log(now, “log”);
    names.emplace(nameFromIdx(idx));
}

그러나 클라이언트 쪽에서 idx 의 값이 short 타입이라면,
보편참조 함수가 호출될까? int형 함수가 호출될까?

분명한 것은 구현자는 int형 보통함수가 호출 되길 기대할 것이다.

(1) 보편 참조 함수 -> 주어진 인수와 정확히 부합함
(2) int형 함수 -> short 인수가 int로 승격되어 호출에 부합함

-> 하지만 기대와 달리 보편 참조 함수가 호출되어 에러가 발생할 것이다.
-> 왜냐하면 보편 참조는 개발자가 생각하는 것보다 일반적으로 훨씬 많은 인수를 빨아들이기 때문이다.


마찬가지 이유로 클래스 내부에서 보편 참조 중복적재(오버로딩)에도 문제가 많다.

class Person
{
public:
    template<typename T>
    explicit Person(T&& n)
        : name(std::forward<T>(n))
    {};  // (1)

    Person(const Person& rhs); //(2)
};

Person p = “Jone”;
auto cloneOfP(p); // (1) 보편참조 생성자 호출 -> p는 string이기 때문에 컴파일 에러

const Person p = “Jone”;
auto cloneOfP(p); // (2) 복사 생성자 호출

C++의 중복적재 해소 규칙 중에는 어떤 함수 호출이 템플릿 인스턴스와 비템플릿 함수(보통 함수)에 똑같이 부합한다면,
보통 함수를 우선시한다는 규칙이 있다.

결과적으로, 보편 참조 매개변수에 대한 중복적재는 가능한 한 피해야 한다!
