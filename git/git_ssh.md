### 원격 접근시 매번 로그인하기 번거로우므로 git에 ssh 공개키를 등록하자
1. SSH 공개키 만들기
```
[irteam@dev-tuyy1-ncl .ssh]$ cd ~/.ssh
[irteam@dev-tuyy1-ncl .ssh]$ ssh-keygen // 여기에 비밀번호 안넣으면 로그인시 비밀번호 필요x

[irteam@dev-tuyy1-ncl .ssh]$ cat ~/.ssh/id_rsa.pub
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC1QhuhkO9JD02edAwkhUoSqxorauEoInOrtla00LHi6rK0ME83GANodkgjBR4Zh+RF03C3nUrnp5QDqEvLBDmVz61M0tWdx1Cq2n7Q65EGZ/BJIrB5I2YNZqJC6ByiC5YphvSCUtG/Ja2zrFm90EC69lGRNfQtI92koyn3vRipJuF+nUEs50i8aFprphxAZ9E3nIzwdWf8rUgXDH10FI94tM4bwkU361m9LWvviCpMc1BP/LxrODeK3L9jRR8rwtBlsMArqRFomH62nRe1uZ0pKG3kpAPneeS/1miHEFnCVwDbjGEfHB+4+KjWfrIFtQV9RZA2UMLJIoZcQYVl9u4L irteam@dev-tuyy1-ncl

```
2. Settings > SSH and GPG keys > New SSH Key 에 id_rsa.pub 키를 
넣자
