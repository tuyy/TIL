# go modules
 Go 1.11과 Go 1.12 부터 ```go modules```를 지원한다.
 
* ```go.mod``` 파일이 루트 디렉터리에 있는 파일 트리에 저장된 Go 패키지의 모음
* ```go.sum``` 파일을 참고하여 처음 다운로드받은 모듈과 나중에 다운로드받는 모듈이 같은 비트를 갖는지를 검사한다.
    * 이를 이용하여 악의적인 이유 등의 복잡한 이유로 프로젝트가 의존하는 모듈이 예상치 못하게 변형되는 일을 방지함.
    * 즉, go.mod와 go.sum 모두 버전 관리 도구에 체크되어야함.
* ```go mod init```은 새로운 모듈을 생성하고, 모듈을 설명하는 go.mod 파일을 초기화한다.
* ```go build```와 ```go test```, 그 외의 패키지 빌드 커맨드들은 필요에 따라 go.mod에 새로운 디펜던시를 추가한다.
* ```go list -m all```은 현재 모듈의 디펜던시 목록을 보여준다.
* ```go get ..```은 필요한 버전의 디펜던시를 변경하거나 새로운 디펜던시를 추가한다.
* ```go mod tidy```는 사용하지 않는 디펜던시를 제거한다.

## go module 생성
```
$ go mod init navercorp.com/tuhello
go: creating new go.mod: module navercorp.com/tuhello
go: to add module requirements and sums:
        go mod tidy
```

## go.mod 파일 확인
```
$ cat go.mod 
module navercorp.com/tuhello

go 1.16
```

## 디펜던시 추가하기
```
$ go get rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go get: added rsc.io/quote v1.5.2

$ cat go.mod 
module navercorp.com/tuhello

go 1.16

require rsc.io/quote v1.5.2 // indirect
```

## 의존성 확인하기
```
$ go list -m all
navercorp.com/tuhello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```

## 버전 업그레이드
```
$ go list -m all
navercorp.com/tuhello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0

$ go get golang.org/x/text
go: downloading golang.org/x/text v0.3.6
go get: upgraded golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c => v0.3.6

$ go list -m all
navercorp.com/tuhello
golang.org/x/text v0.3.6
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```

## 업그레이드 가능 버전 확인
```
$ go list -m -versions rsc.io/sampler
rsc.io/sampler v1.0.0 v1.2.0 v1.2.1 v1.3.0 v1.3.1 v1.99.99
```

## 명시적으로 버전 변경
```
$ cat go.mod
module navercorp.com/tuhello

go 1.16

require (
	golang.org/x/text v0.3.6 // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.99.99 // indirect
)

$ go get rsc.io/sampler@v1.3.1
go: downloading rsc.io/sampler v1.3.1
go get: downgraded rsc.io/sampler v1.99.99 => v1.3.1

$ cat go.mod
module navercorp.com/tuhello

go 1.16

require (
	golang.org/x/text v0.3.6 // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.3.1 // indirect
)
```

## 사용하지 않는 디펜던시 제거
 * $ go mod tidy
```
$ go get rsc.io/quote/v3
go: downloading rsc.io/quote/v3 v3.1.0
go get: added rsc.io/quote/v3 v3.1.0

$ go list -m all
navercorp.com/tuhello
golang.org/x/text v0.3.6
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
rsc.io/quote v1.5.2
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.1

$ go mod tidy
$ go list -m all
navercorp.com/tuhello
golang.org/x/text v0.3.6
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.1
```
