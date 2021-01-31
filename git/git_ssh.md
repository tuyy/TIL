### 원격 접근시 매번 로그인하기 번거로우므로 git에 ssh 공개키를 등록하자
1. SSH 공개키 만들기
```
[irteam@dev-tuyy1-ncl .ssh]$ cd ~/.ssh
[irteam@dev-tuyy1-ncl .ssh]$ ssh-keygen // 여기에 비밀번호 안넣으면 로그인시 비밀번호 필요x

[irteam@dev-tuyy1-ncl .ssh]$ cat ~/.ssh/id_rsa.pub
ssh-rsa ....

```
2. Settings > SSH and GPG keys > New SSH Key 에 id_rsa.pub 키를 
넣자
