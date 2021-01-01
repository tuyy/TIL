std::thread 객체는 joinable 상태이거나 unjoinable 상태이다.

unjoinable 상태인 std::thread 객체로는 다음과 같은 것들이 있다.

1) 기본 생성된 std::thread
2) 다른 std::thread 객체로 이동된 후의 std::thread 객체
3) join이 호출된 std::thread
4) detach가 호출된 std::thread

std::thread의 joinable 가능성이 중요한 이유는, 만일 joinable 한 스레드의 소멸자가 호출되면 프로그램 실행이 종료된다는 것이다.

그렇다면, 왜? joinable 한 스레드의 소멸자가 호출되면 프로그램 실행이 종료될까?
그 이유는, 다른 두 옵션이 명백히 더 나쁘기 때문이다.

1) 암묵적 join - 추적하기 어려운 성능이상
2) 암묵적 detach - 프로그램이 종료될 수도 있음


그러므로, joinable한 std::thread는 모든 경로에서 unjoinable한 상태로 만들어야한다!

이럴때 쓰는게 RAII 객체이다.

class ThreadRAII {
public:
    enum class DtorAction {join, detach};

    ThreadRAII(std::thread&& t, DtorAction a)
        : action(a), t(std::move(t)) {}

    ~ThreadRAII() {
        if (t.joinable() == false) {
            if (action == DtorAction::join) {
                t.join();
            }
            else {
                t.detach();
            }
        }
    }
private:
    DtorAction action;
    std::thread t;
};


기억해 둘 사항들
1) 모든 경로에서 std::thread를 합류 불가능으로 만들어라
2) 소멸 시 join 방식은 디버깅하기 어려운 성능 이상으로 이어질 수 있다.
3) 소멸 시 detach 방식은 디버깅하기 어려운 미정의 행동으로 이어질 수 있다.
4) 자료 멤버 목록에서 std::thread 객체를 마지막에 선언하라
