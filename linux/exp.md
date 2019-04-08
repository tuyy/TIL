#### bash shell ssh 암호 넣기 자동화

```
# ~/.bashrc
alias dgw='exp {pwd} ssh yabi1004@dgw01.nnsystem.com'


# exp 실행파일 pass 지정해야함
!/usr/bin/expect

set timeout 20

set cmd [lrange $argv 1 end]
set password [lindex $argv 0]

eval spawn $cmd
expect "password:"
send "$password\r";
interact

```

* https://stackoverflow.com/questions/12202587/automatically-enter-ssh-password-with-script
