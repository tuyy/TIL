# godepgraph
 graphviz 기반 GO 의존성 그래프 이미지 생성하는 프로그램
 
* ```$ go get github.com/kisielk/godepgraph```
* ```$ godepgraph github.com/kisielk/godepgraph | dot -Tpng -o godepgraph.png```
    * ```$ dot``` 이 실행되지 않는 경우 ```$ brew install graphviz```
  

## 예시
* -o 옵션은 포함할 패키지 prefix, -i 옵션은 제외할 패키지 prefix
  * 콤마로 구분함
```
$ ls
Makefile	conf		dist		go.mod		go.sum		hellobo.png	pkg		png.png
$ pwd
/Users/jutaeun/tuyy/go-workspaces/src/oss.navercorp.com/taeun-ju/hellobo-go
$ godepgraph -s -i oss.navercorp.com/taeun-ju/hellobo/pkg/log -o oss.navercorp.com/taeun-ju/hellobo/pkg oss.navercorp.com/taeun-ju/hellobo/pkg/cmd/hellobo | dot -Tpng -o hellobo.png
```

## 출력 결과
![hellobo](https://user-images.githubusercontent.com/13283116/118144942-fd488c80-b447-11eb-9256-afc1625028cd.png)
