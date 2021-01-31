### 전통적인 프로세스 실행 방법 fork-exec
* fork는 유닉스 계열의 프로세스를 복제하는 전통적인 방법이다.
* 프로세스를 복제하는 이유는 멀티태스킹을 하기 위해서다.
* 복수개의 CPU가 설치된 경우에 뛰어난 응답성과 성능을 보여줄 가능성이 크다.
* 반면 프로세스 간 통신 비용이 큰 경우 기대한 만큼 성능이 나오지 않을 것이다.
* 그래서 복제된 프로세스들이 공통으로 사용하는 파일이나 I/O는 mmap을 이용해서 비용을 최소화하는 것도 고려해야한다.
* 프로세스 복제가 가장 많이 쓰이는 예로는 shell이 있다.
* fork-exec 순서의 호출을 더 간단하게 코딩하기 위해 popen이나 system과 같은 함수를 사용할 수 있다.
* popen은 부모 프로세스와 자식 프로세스 간에 주고받을 데이터가 있을 때 파이프를 이용하여 통신할 수 있는 기능이 있다.
* system은 어떤 피드백도 없이 간단히 프로그램을 실행할 때 사용한다.

```fork-exec CODE
...
char *argv_exec[] = {"forkexec_child", (char*)NULL};
switch(pid_child = fork())
{
  case 0: /* child process */
    execv(argv_exec[0], argv_exec);
    break;
  case -1: /* error */
    perror("FAIL: FORK");
    break;
  default: /* parent process */
    wait(NULL);
    break;
}
...
```

#### vfork
* vfork는 fork-exec를 좀 더 가볍게 하려고 지원하는 기능이다.
* fork-exec는 exec가 호출되는 순간 fork로 복제되었던 페이지 테이블은 모두 해제되는 단점이 있다.
* 즉, fork-exec에는 쓰지도 않는 자원들의 복제 때문에 오버헤드가 존재한다.
* 이러한 오버헤드를 줄이기 위해 copy on write를 지원하는 vfork가 제안되었다.
* 하지만 요즘 fork에도 지원하는 기능이다. 결과적으로 vfork를 쓸 이유가 없다.

#### system
* fork-exec를 간단하게 구현한 형태이다.
* system은 실행 명령어가 작동되는 동안에 부모 프로세스가 잠시 정지된다.
* 또한 SIGCHLD 시그널과 SIGINT, SIGQUIT 시그널도 무시된다.
* 정말 간단한 경우가 아니라면 fork-exec로 구현하는 것을 권장한다.
* 왠만하면 정말 쓰지 말자

### 확장된 프로세스 실행 방법 posix_spawn
* 왜? posix_spawn은 fork-exec를 대체하려 하는가?
* fork는 부모 프로세스를 복제할 때 모든 정적 정보를 복제하기 때문에 쓰지도 않는 자원을 복제하는 오버헤드가 존재한다.
* posix_spawn은 부모 프로세스의 6가지 자원(열린 파일, 프로세스 그룹ID, 유저/그룹ID, 시그널 마스크, 스케쥴링)을 선택적으로 복제 및 관리할 수 있도록 디자인 되었다.
