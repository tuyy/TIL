## Item 5: Prefer auto to explicit type declarations.
* auto의 장점
    * 비교적 타이핑이 아주 짧아진다.
    * 잘못된 타입을 선언할 실수를 막아준다.
    * 대체로 std::function 객체는 auto로 선언된 객체보다 메모리를 더 소비한다.
    
* auto의 단점
    * 명시적이지 않다. (어떤 타입인지 한 눈에 알 수 없어 추적이 필요할 수 있다.)
    * auto의 잘못된 타입 추론 관련 문제가 있을 수 있다. (std::initialization_list..)
    
* 그래서?
    * auto의 사용은 선택의 문제이다. 
    * 명확히 타입을 알 수 있는 경우가 아니라면 auto를 쓰지 않도록 할 것 이다.
    
### 참고
* Effective Modern C++ by Scott Meyers
