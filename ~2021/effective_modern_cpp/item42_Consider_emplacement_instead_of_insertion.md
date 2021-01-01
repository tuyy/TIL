std::vector<std::string> vs;    // std::string 컨테이너
vs.push_back(“xyzzy”);         // 문자열 리터럴을 추가

위 코드를 수행하면 어떤일이 벌어질까?
컴파일러는 std::string과 const char[6] 형식이 일치하지 않음을 인식하여 문자열 리터럴로 부터 std::string의 임시 객체를 생성한다.
이 임시 std::string 객체를 push_back에 전달하는 코드를 산출한다. 코드는 아래와 같다.

vs.push_back(std::string(“xyzzy”)); // 임시 std::string 객체를 생성해서 push_back에 전달

실행시점에서 위의 push_back 호출에 의해 일어나는 일은 다음과 같다.

1. std::string(“xyzzy”) 임시 std::string 객체를 생성한다. (임시 객체이므로 오른값이다.)
2. 이 임시객체가 오른값 오버로딩 push_back 함수로 전달된다. 이 버전의 임시객체는 오른값 참조 매개변수 x에 묶인다.
    그런 다음 std::vector를 위한 메모리 안에서 x의 복사본이 생성된다. 이에 의해 std::vector 안에 실제 새 객체가 만들어진다.
3. push_back이 반환된 즉시 임시 std::string 객체가 파괴되어 소멸자가 실행된다.

문제는 임시 객체의 생성과 소멸이다. 만약 std::vector 안에 실제 새객체에 직접 문자열 리터럴을 추가할 수 있다면?! 성능업!
바로 emplace_back을 사용하면 가능하다.

vs.emplace_back(“xyzzy”);  // “xyzzy”를 이용해서 std::string을 vs 안에 직접 생성

push_back 류 함수들은 삽입할 객체를 받지만, emplace_back 류 함수들은 삽입할 객체의 생성자를 위한 인수들을 받는다.

std::string queenOfDisco = “hello”;
vs.push_back(queenOfDisco);   // vs 끝에서 queenOfDisco를 복사 생성
vs.emplace_back(queenOfDisco);  // push_back과 동일

push_back 류와 emplace_back 류 중 어느 것이 더 빠른지 알고 싶다면 직접 측정해봐야한다.
(무조건 emplace_back 류가 빠르지 않다.)

emplace_back 류를 사용하는게 성능에 바람직한지 식별하는 발견법은 다음과 같다.

1. 추가할 값이 컨테이너에 배정되는 것이 아니라 컨테이너 안에서 생성된다. (배정은 임시객체가 필요하다..)

vs.emplace(vs.begin(), “xyzzy”);  // "xyzzy”를 vs의 시작에 추가, 생성이 아니라 이동배정이 됨.

노드 기반이 아닌 컨테이너는 emplace_back이 항상 배정 대신 생성을 이용해서 새 값을 컨테이너에 넣는다.
(std::vector, std::deque, std::string)

2. 추가할 인수 형식들이 컨테이너가 담는 형식과 다르다.
직접 컨테이너에 넣으니까 임시객체 생성이 불필요하다.

3. 컨테이너가 기존 값과의 중복 때문에 새 값을 거부할 우려가 별로 없다.
컨테이너가 중복을 허용하거나 또는 추가할 값들이 대부분 고유한 경우에 해당한다. 이 조건이 중요한 이유는, 중복 제한이 있는 경우 일반적으로 emplace 구현은 노드를 생성해서 그것을 기존 컨테이너 노드들과 비교 한다는 것이다.


정리하자면, 아래와 같은 경우이다.
vs.emplace_back(“xyzzy”);   // 새 값을 컨테이너의 끝에 생성, 인수의 형식이 컨테이너 요소의 형식과 다름, 중복요소 거부하지 않음
vs.emplace_back(50, ‘x’);


emplace_back 류를 피할지 고민해야하는 두 가지 경우는 아래와 같다.

1. 임시객체가 필요한 경우

std::list<std::shared_ptr<Widget>> ptrs; // 커스텀 삭제자가 필요하여 std::make_shared 사용하지 못함

// 임시객체가 생성되어 컨테이너에 넣는 도중에 메모리 부족이 발생해도 임시 Widget 객체의 killWidget 커스텀 삭제자가 호출됨
ptrs.push_back(std::shared_ptr<Widget>(new Widget, killWidget));
ptrs.push_back({ new Widget, killWidget });

반면,

// 컨테이너에 넣는 도중에 메모리 부족이 발생하면 커스텀 삭제자를 호출할 방법이 없으므로 메모리 누수가 발생한다.
ptrs.emplace_back(new Widget, killWidget);

사실 이런 경우엔, 자원 관리 객체를 생성하는 작업을 별도의 문장으로 수행해야한다.

std::shared_ptr<Widget> spw(new Widget, killWidget);
ptrs.push_back(std::move(spw));


2. explicit 생성자들과의 상호 작용 방식에 대해

std::vector<std::regex> regexes;
regexes.emplace_back(nullptr); // 정규식들의 컨테이너에 nullptr을 추가?

std::regex r = nullptr;  // 컴파일 안됨, 복사 초기화 방식
regexes.push_back(nullptr);  // 컴파일 안됨

std::regex r(nullptr);  // 컴파일됨, 직접 초기화 방식

emplace 류 함수는 직접 초기화 방식을 사용한다.
regexes.emplace_back(nullptr);  // 컴파일됨. 직접 초기화에서는 포인터를 받는 explicit로 선언된 std::regex 생성자가 있음.
regexes.push_back(nullptr);  // 복사 초기화에서는 이런 생성자를 사용할 수 없음.


기억해 둘 사항들
1) 이론적으로, 생성 삽입 함수들은 종종 해당 삽입 버전보다 더 효율적이어야 하며, 덜 효율적인 경우는 절대로 없어야 한다.
2) 실질적으로, 만일 (1) 추가하는 값이 컨테이너로 배정되는 것이 아니라 생성되고, (2) 인수 형식들이 컨테이너가 담는 형식과 다르고, 그 값이 중복된 값이어도 컨테이너가 거부하지 않는다면, 생성 삽입함수가 삽입 함수보다 빠를 가능성이 아주 크다.
3) 생성 삽입 함수는 삽입 함수라면 거부 당했을 형식 변환들을 수행할 수 있다.
