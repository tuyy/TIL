## gdb
* GNU linux 디버깅툴, breakpoint를 걸어놓고 쉽게 프로그램 흐름을 따라갈 수 있다.
* 변수 값 확인, 변수 값 임의로 수정, 스택 추적 등 지원
```
// -g 옵션으로 디버깅 심볼 생성해야함
g++ -g -o main main.cpp
gdb main
```

```
set listsize 20
set variable [변수명] = [대입값]
bt
f [..frame number]
i b // info breakpoint
i locals // 로컬 변수 확인
d [1,2,3 .. Num] // breakpoint 삭제 (여러개 입력 가능)
p [변수명]
r [..args] // 실행, 재실행
q // 종료
c // 다음 breakpoint로
s // 다음 스탭
n // 함수 건너띄고 다음 스탭
b [line]
b [filepath]:[line]
```
