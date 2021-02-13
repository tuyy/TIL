### kafka 토픽 삭제하기

```

### 설정에 따라 삭제가 불가능 할 수 있다.
$ ./kafka-topics.sh --delete --zookeeper localhost:2181 --topic savebo
Topic savebo is marked for deletion.
Note: This will have no impact if delete.topic.enable is not set to true.

### 토픽 상태 확인
$ ./kafka-topics.sh --describe --zookeeper localhost:2181 --topic savebo
Topic:savebo    PartitionCount:1        ReplicationFactor:1     Configs:
        Topic: savebo   Partition: 0    Leader: 0       Replicas: 0     Isr: 0

### 주키퍼를 이용하여 삭제
$ ./zookeeper-shell.sh localhost:2181
Connecting to localhost:2181
Welcome to ZooKeeper!
JLine support is disabled

WATCHER::

WatchedEvent state:SyncConnected type:None path:null
ls /brokers/topics
[savebo, Hello, index_del, __consumer_offsets, mytest2]
rmr /brokers/topics/savebo
```
