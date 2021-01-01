### ssh 비밀번호입력 과정을 bash script에 넣기
1. exp.sh 스크립트 작성
```
#!/usr/bin/expect

set timeout 20

set cmd [lrange $argv 1 end]
set password [lindex $argv 0]

eval spawn $cmd
expect "assword:"
send "$password\r";
interact
```

2. 스크립트 실행
    * exp.sh 을 alias 하거나 PATH에 지정하여 쉽게 접근할 수 있음
```
./exp.sh <password> ssh <anything>
```
