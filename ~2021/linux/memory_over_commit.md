#### memory overcommit?
* 실제 메모리 이상의 공간을 확보할 수 있는 Linux의 메모리 관련 메커니즘이다.
* memory commit은 프로세스는 메모리 할당이 되었다고 생각하지만 실제로는 물리 메모리 어느 곳에도 할당되어지지 않은 상태이다.
    * system call을 통한 메모리 요청(malloc)
    * kernel은 실제 물리메모리에 할당은 하지 않고 가상의 메모리 영역으로 주소 반환
* memory overcommit은 memory commit을 더 할 수 있게 해준다는 뜻이다.

#### 왜 필요한가?
* 프로세스는 대부분 할당 받은 메모리를 모두 사용하지 않는다.
* 하지만 많은 메모리를 필요로하는 순간이 있기도 하다. 바로 fork() 시스템콜과 같은 작업을 처리할 때 필요하다.
* fork()를 통해 부모 프로세스가 사용중인 모든 메모리를 복사 할 수 있으므로 실제 자식 프로세스가 메모리를 사용하지 않아도 많은 양의 복사가 이루어진다.
* 부모 프로세스의 메모리 영역을 복사하는 과정에서 memory commit이 일어나게 되고, 경우에 따라서는 overcommit이 필요하다.

#### OOM Killer?
* overcommit 가능한 메모리가 부족할때 커널은 OOM Kiler를 이용하여 프로세스를 kill하여 메모리를 확보한다.

#### linux overcommit 설정
* OOM Killer는 vm.overcommit_memory 커널 설정에 의존한다. 
* overcommit 사용 여부를 설정할 수 있다.
* https://www.kernel.org/doc/Documentation/vm/overcommit-accounting
