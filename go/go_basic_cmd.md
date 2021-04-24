# go 기본 명령어
 go run, go build, go install
 
## go run
* ```$ go run .```
    * 현재 디렉토리(패키지) 기준으로 main 함수를 찾아서 실행한다.

## go build
* ```$ go build .```
    * 현재 디렉토리(패키지) 기준으로 main 함수를 찾아서 binary를 생성한다.
    * 파일이 존재하는 디렉토리(패키지)명 기준으로 명령어 실행 위치 기준으로 생성된다.
    
## go install
* ```$ go install hello.go```
    * $GOPATH 하위에 bin 디렉토리에 binary를 생성한다. ($GOBIN 하위)
