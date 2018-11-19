## 카산드라 version upgrade
* 1.2 to 2.0

### 작업 순서
1. cassandra v2.0 설치 및 configure 파일 정보 migration
    * 꼭 확인해야하는 config : cluster name, seeds, listen_address, rpc_address, logs_path ...
2. cassandra v1.2 종료
```
$ ./bin/nodetool disablegossip
$ ./bin/nodetool drain
$ kill {cassandra_pid}
```
3. cassandra v2.0 실행 (로그 확인: warning or error 체크)
4. start 완료 후 확인사항
    * ```./bin/nodetool status```
    * opscenter 대쉬보드에 정상적으로 version upgrade 되었는지 확인
5. ```./bin/nodetool upgradesstables``` 실행
    * 해당 커맨드를 입력 안하더라도 정상동작하지만 온전히 신규 버전의 장점을 취하려면 사용하는게 좋다고함.
6. 모든 노드에도 동일하게 작업
