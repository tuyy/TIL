### 간단 카프카 클러스터 구성
* kafka, zookeeper 3대 이상 필요함
* 설정 파일 두개 수정, 파일 하나 추가 필요
```
########## config/zookeeper.properties 수정 ##########
..
dataDir=/home1/irteam/data/zookeeper
# the port at which the clients will connect
clientPort=2181
# disable the per-ip limit on the number of connections since this is a non-production config
maxClientCnxns=0

initLimit=5
syncLimit=2
server.1=dev-tuyy0-cassandra001-ncl.nfra.io:2888:3888 # 클러스터 구성 장비들 추가
server.2=dev-tuyy0-cassandra002-ncl.nfra.io:2888:3888
server.3=dev-tuyy0-cassandra003-ncl.nfra.io:2888:3888
..

########## dataDir에 주키퍼 식별id 파일 추가 (각 주키퍼 서버당 값을 다르게 줘야함) ##########
[irteam@dev-tuyy0-cassandra001-ncl kafka_2.10-0.10.2.1]$ cat /home1/irteam/data/zookeeper/myid
1

########## config/server.properties 수정 ##########
..
broker.id=0 # 브로커 서버 당 uniq하게 지정해야한다.
zookeeper.connect=dev-tuyy0-cassandra001-ncl.nfra.io:2181,dev-tuyy0-cassandra002-ncl.nfra.io:2181,dev-tuyy0-cassandra003-ncl.nfra.io:2181
..
```

### 실행
```
실행
$ cd /home1/irteam/apps/kafka_2.10-0.10.2.1
$ ./bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
$ ./bin/kafka-server-start.sh -daemon config/server.properties

# 토픽 생성
$ ./bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 3 --partitions 3 --topic mytest1

# producer
$ bin/kafka-console-producer.sh --broker-list localhost:9092 --topic mytest1
This is a message
This is another message

# consumer
$ bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic mytest1 --from-beginning
This is a message
This is another message

```
