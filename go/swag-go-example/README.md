# swag-go example
 swag를 사용하면 코드와 API 가이드 문서를 함께 관리할 수 있다.

## start
 ```swag init``` 을 빌드시 Makefile 이용해서 실행하자
 
* ```swag init```를 실행한 후 생성된 doc을 handler 주석이 포함된 파일에 import 해야한다.
     * ```import _ "oss.navercorp.com/taeun-ju/go-swag-test/docs"```
 
```
$ swag init
$ go run main.go
```

* 실행 결과

<img width="890" alt="스크린샷 2021-0을4-25 오후 12 11 25" src="https://user-images.githubusercontent.com/13283116/115979229-ad507580-a5bf-11eb-9707-248adaa5444c.png">


## 참고
* https://github.com/swaggo/swag
* https://linggar.asia/?p=594