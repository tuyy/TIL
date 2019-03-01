# js에서 "use stric;" 란?
* ECMAScript 5버전에 있는 새로운 기능으로써, 프로그램 함수를 엄격한 운용 컨텍스트 안에서 실행시킬 수 있게 한다. 즉, 좀 더 많은 예외를 발생시킨다.
    * 흔히 발생하는 코딩 실수를 잡아내서 예외를 발생시킨다.
    * 상대적으로 안전하지 않은 액션 수행시 예외를 발생시킨다.
    * 혼란스럽거나 제대로 고려되지 않은 기능들을 비활성화 시킨다.
    
### 예시
```javascript
// Non-strict code...

"use strict"; // 1. 전역적으로 사용 가능하다.

(function(){
    "use strict"; // 2. 함수안에서도 가능하다.

    // Define your library strictly...
})();

// Non-strict code...
```
