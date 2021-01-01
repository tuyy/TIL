### Global Interpeter Lock
* coarse-grained 락 : 큰 단위로 락을 잡는 방식
* GIL은 단일 코어 CPU 환경에서 성능이 빠르다.
* 하지만 자원 배분을 한 쓰레드에게 할당하고 끝날때까지 다른 쓰레드가 접근하지 못한다.
* CPU Bound에는 취약하지만 I/O Bound에는 영향이 없다.
* GIL의 장점도 많지만 이로 인해 실제 쓰레드를 사용할 때 원하는 효과를 얻을 수 없다.
* 결론은 python Thread를 사용할바엔 multiprocessing 모듈을 사용하자.
