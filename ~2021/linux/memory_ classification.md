### 프로세스의 다양한 메모리와 적재되는 위치, 기능에 대해

#### 텍스트(.text)
* 일반적으로 프로그램의 실행 코드가 존재하는 영역이다.
* 소스 코드를 컴파일해서 만들어진 기계어 코드가 실행을 위해 적재되는 곳을 의미한다.

#### 데이터
* 데이터 영역은 주로 전역 변수나 정적 심볼들이 사용되는 영역이다.
* 데이터 영역의 분류
    * .rodate : 읽기 전용으로 초기화되는 영역이다. 예를들어 const로 선언되는 변수, 상수 문자열, 포맷 스트링 등이 들어간다.
    * .data : 읽기/쓰기가 가능한 영역으로 초기화되는 영역이다.
    * BSS : 초기화되지 않는 영역이다. 초기화되지 않은 전역 변수나 static으로 선언된 변수가 여기에 위치한다.
    
#### 스택(stack)
* 스택 영역은 로컬 자동 변수, 즉 임의로 메모리를 잡지 않아도 함수의 시작에서 자동으로 생성, 파괴되는 변수를 말한다.

#### 힙(heap)
* 힙은 malloc이나 calloc 같은 메모리 할당 함수를 사용하여 얻는 공간으로 프로세스 어디서든 접근 가능하다.
* 프로세스가 종료하지 않는 한 자동으로 파괴되지 않으므로 할당과 해제를 프로그래머가 엄격히 해야한다.

#### 전역 변수들의 형태에 따른 메모리 위치
```
int num; // bss에 위치
int num= 1; // data에 위치
char str[] = "hello"; // data에 위치(읽기/쓰기 가능)
char *p_str = "hello"; // 문자열 hello는 .rodata에, p_str 포인터 변수는 data에 위치
const char str[] = "hello"; // 문자열 hello는 .rodata에, str변수는 data에 위치
static int i_val; // bss에 위치
```

#### Stack Size에 대해
* 일반적으로 스택은 할당과 해제가 매우 빠르다고 알려졌다. 하지만 스택은 실제로 할당만 할 뿐 해제를 하지 않는다.
* 스택은 필요에 따라 할당을 하면서 계속 키워나가는데 사용 후 반환해도 여전히 스택 크기를 계속 유지하면서 재사용한다.
* 이러한 특징 때문에 프로세스가 매우 큰 스택 메모리를 한 번이라도 사용하게 되면 프로세스가 종료하기 전까지는 메모리가 낭비된다.

``` 스택 크기를 제한하는 설정
$ ./ulimit -a
..
stack size       (kbytes, -s) 8192
..
```

```
#include <stdio.h>

#define BUFFER 1024 * 1024

int exhaust_stack(int count)
{
    char buffer[BUFFER]; // 1MB
    if (count <= 0)
    {
        printf("reach break position.. \n");
        return 0;
    }
    sprintf(buffer, ">> exhaust 1MB stack(addr:%p), will more %d MB..\n", buffer, 10 - count + 1);
    printf("%s", buffer);
    exhaust_stack(count - 1);
    return 0;
}

int main()
{
    exhaust_stack(10); // error Segmentation fault: 11 발생, max stack size:8MB
    return 0;
}
```

* 참고로 스레드의 경우 별도 스택 공간과 크기 설정을 가지므로 위의 설정과는 다르게 동작한다.
* 리눅스의 경우 스레드는 힙 공간에 별도의 가상 스택을 만들게 되므로 프로세스의 스택과는 연관이 없다.

