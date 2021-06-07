# Zookeeper로 Kafka 정보 확인하기
 Zookeeper client를 이용하여 Kafka 정보를 확인한다.

* dms2 kafka 주키퍼 버전 및 접속 정보 확인
    * Zookeeper v3.4.14
    * zookeeper-dms2.ndrive.svc.ar1.io.navercorp.com:2181/kafka

<img width="720" alt="스크린샷 2021-06-07 오후 8 32 07" src="https://media.oss.navercorp.com/user/1698/files/85644c80-c7d0-11eb-989c-a7c3923bd636">

* https://ndrive.pasta.navercorp.com/kafka/dms2/instance?id=602dc39ed24babae64dc249d

## Zookeeper client 다운로드 및 실행

* Zookeeper 다운로드
    * https://archive.apache.org/dist/zookeeper/

```
$ curl -O https://archive.apache.org/dist/zookeeper/zookeeper-3.4.14/zookeeper-3.4.14.tar.gz
$ tar xvzf zookeeper-3.4.14.tar.gz
$ cd zookeeper-3.4.14

$ ./bin/zkCli.sh -server zookeeper-dms2.ndrive.svc.ar1.io.navercorp.com:2181/kafka

// ls, get 명령어를 활용하여 저장된 값을 확인한다.
```

## kafka 관련 Zookeeper Tree
* /kafka/brokers
* /kafka/consumers


## Zookeeper 특성 및 기본 명령어 
* znode 기본 특성
    * Persistent znodes
      * default로 생성되는 znode이며, 영구적으로 저장된다.
    * Ephemeral znodes
      * 임시 znode이며, 클라이언트가 Zookeeper 서버와 연결이 종료되자마자 제거된다.
    * Sequential znodes
      * 연속 znode에는 이름 끝에 숫자 순서로 10자리 숫자가 지정된다.


* 기본 명령어    
    * znode 목록보기
      * ```ls /[znode-name]```
    * znode 읽기
      * ```get /[znode-name]```
    * znode 생성
      * ```create -[options] /[znode-name] [znode-data]```
    * znode 쓰기
      * ```set /[znode-name] [new-znode-data]```
    * znode 삭제
      * ```delete /[znode-name]```