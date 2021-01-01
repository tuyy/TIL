일반적으로, std::async를 호출해서 함수를 실행하는 것에는 그 함수를 비동기적으로 실행하겠다는 의도가 깔려있다.
그러나, std::async 호출은 함수를 어떤 시동방침(launch policy)에 따라 실행한다는 좀 더 일반적인 의미를 가진다.

std::launch::async 시동 방침을 지정하면 f는 반드시 비동기적으로, 다시 말해서 다른 스레드에서 실행된다.
std::launch::deferred 시동 방침을 지정하면 f는 std::future 객체에 대해 get or wait가 호출될 때에만 동기적으로 실행된다. 즉 f가 실행될 때까지 호출자는 차단된다. get이나 wait가 호출되지 않으면 f는 결코 실행되지 않는다.

하지만, default 시동 방침은 아래와 같다.

auto fut1 = std::async(f);
auto fut2 = std::async(std::launch::async | std::launch::deferred, f);

 결과적으로, default 시동 방침은 f가 비동기적일 수도, 동기적으로 실행될 수 있음을 의미한다.
이러한 유연성 덕분에 스레드의 생성과 파괴, 과다구독 회피, 부하 균형화의 책임을 떠맡을 수 있는 것이다.


그래서, 어떤 과제에 대해 기본 시동 방침과 함께 std::async를 사용하는 것은 다음 조건들이 모두 성립할 때에만 적합하다.

1) 과제가 get or wait를 호출하는 스레드와 반드시 동시적으로 실행될 필요 없는 경우
2) 여러 스레드 중 어떤 스레드의 thread_local 변수들을 읽고 쓰는지가 중요하지 않는 경우
3) 과제가 지연된 상태일 수도 있다는 점이 wait_for or wait_until을 사용하는 코드에 반영되어 있는 경우

이 조건 중 하나라도 성립하지 않는다면, 시동 방침을 std::launch::async로 강제해야한다.

auto fut = std::async(std::launch::async, f);  // f를 비동기적으로 실행

template <typename F, typename… Ts>
auto reallyAsync(F&& f, Ts&&… params) {
    return std::async(std::launch::async,
                                std::forward<F>(f),
                                std::forward<Ts>(params)...);
}


기억해 둘 사항들
1) std::async의 기본 시동 방침은 과제의 비동기적 실행과 동기적 실행을 모두 허용한다.
2) 그러나 이러한 유연성 때문에 thread_local 접근의 불확실성이 발생하고, 과제가 절대로 실행되지 않을 수도 있고, 시간 만료 기반 wait 호출에 대한 프로그램 논리에도 영향을 미친다.
3) 과제를 반드시 비동기적으로 실행해야 한다면 std::launch::async를 지정하라
