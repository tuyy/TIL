## buffer
* buffer는 어떻게 사용하는게 좋을까?
* C++은 대표적으로 3가지 방법이 있다.
```C++
// 1) char sBuffer[BUFFER_SIZE] = {'\0'};
// 2) std::array<char, BUFFER_SIZE> sBuffer = {'\0'};
// 3) auto sBuffer = std::make_unique<char[]>(BUFFER_SIZE);
```

### 1) ```char sBuffer[BUFFER_SIZE]```
* c 언어에서 부터 사용되어온 대표적인 형태
* 초기값을 넣지 않으면 임의의 값들이 넣어진다.
* sBuffer[idx] 잘못된 idx에 접근하면 미정의 동작을 유발한다.
* 포인터로 넘기는 경우 size도 같이 넘겨야 한다.
* 여러모로 불편하고 오래된 느낌
```
[irteam@dev-tuyy-ncl 010_array_ex]$ time ./main 1 10000000

real    0m0.036s
user    0m0.033s
sys     0m0.002s
```

#### 2) ```std::array<char, BUFFER_SIZE> sBuf```
* 1)의 C++ 버전이다.
* 초기값을 넣지 않으면 임의의 값들이 넣어진다. (char[]와 동일)
* .at(..), .fill(..) 등 지원하고 STL 알고리즘 함수 사용하기 쉽다.
* .size() 함수로 size를 구할 수 있다.
* char[]의 상위 호환 버전
```
// 특이사항으로 char[]와 초기화 성능이 거의 동일
[irteam@dev-tuyy-ncl 010_array_ex]$ time ./main 2 10000000

real    0m0.036s
user    0m0.031s
sys     0m0.005s
```

#### 3) ```auto sBuf = std::make_unique<char[]>(BUFFER_SIZE)```
* unique_ptr은 스마트 포인터인데, 배열을 지원한다.
* 기본적으로 '\0' 값이 할당된다. (다른 버퍼 형식보다 초기화 속도가 느릴 수 밖에 없다.)
* sBuf.get() <- char[] 와 동일하게 사용할 수 있다.
```
// 초기값에 '\0'을 할당하므로 초기화 속도가 느림
[irteam@dev-tuyy-ncl 010_array_ex]$ time ./main 3 10000000

real    0m1.472s
user    0m1.468s
sys     0m0.001s
```
