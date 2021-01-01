#### std::any
* c++17 부터 지원 ```#include <any>```
* 간단 정리
```
기존에 명확한 타입을 지정하지 않을 때, void*를 사용하였다.
하지만, void* 를 사용하는 경우 실수를 유발하기가 쉽다.
이를 보안하기 위해 std::any가 지원된다.

std::any는 타입을 알 수 있고, 복사가 가능하며(타입을 알기때문에), 스마트하게 deleter가 호출된다.
```

* 참조: https://blogs.msdn.microsoft.com/vcblog/2018/10/04/stdany-how-when-and-why/
