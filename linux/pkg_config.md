### pkg-config
* linux 환경에서 링킹 정보를 손쉽게 관리
* xxxx.pc 파일에 해당 링킹 정보를 넣고 PKG_CONFIG_PATH 환경변수에 경로를 넣어야한다. (아래 참고)
* 환경변수는 ```~/.bash_profile```에 넣는다.
```
[irteam@dev-tuyy-ncl ~]$ pkg-config --libs rdkafka
Package rdkafka was not found in the pkg-config search path.
Perhaps you should add the directory containing `rdkafka.pc'
to the PKG_CONFIG_PATH environment variable
No package 'rdkafka' found

[irteam@dev-tuyy-ncl src]$ PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/lib/pkgconfig pkg-config --libs rdkafka
-L/usr/local/lib -lrdkafka
```
