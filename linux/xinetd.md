### 슈퍼데몬 xinet 설정법 (testsynap 예시)
1. sudo vi  /etc/xinetd.d/testsynap 파일 추가
    * 서버 경로가 아래 있어야한다.
```
service testsynap
{
disable = no
flags = REUSE
socket_type = stream
wait = no
user = irteam
port = 27901
protocol = tcp
server = /home1/irteam/tuyy/collect/test/home/synapfilter
}
```

2. sudo vi /etc/services 에 아래와 같이 등록
```
..
testsynap       27901/tcp               # testsynap
testsynap       27901/udp               # testsynap
..
```

3. xindetd 재시작
```
$ /etc/rc.d/init.d/xinetd restart
```
