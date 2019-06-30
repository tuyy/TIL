## go dep
* 매번 하나씩 go get 으로 의존성을 다루는건 협업할때 매우 번거로워서 ```go dep```을 이용하자
* ```Gopkg.lock  Gopkg.toml vendor``` 파일&디렉토리가 생긴다. vendor는 .gitignore 하자
* ```Gopkg.toml``` 에 ```go dep``` 설정 정보를 넣는다.
    * https://golang.github.io/dep/docs/Gopkg.toml.html

### 기본 예시
```
$ cd $GOPATH/github.com/tuyy/example # 작업할 디렉토리
$ dep init
$ dep ensure -add gopkg.in/confluentinc/confluent-kafka-go.v1/kafka
$ dep ensure
```
