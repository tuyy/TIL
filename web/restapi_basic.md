## Restfull API
 REST는 Representational State Transfer 의 약자로, 아래와 같이 구성된다.
 
 * 자원(Resource) - URI
 * 행위(Verb) - HTTP METHOD
 * 표현(Representations)

### REST API 설계
* 리소스명은 동사보다는 명사를 사용한다.
```
GET /books/show/1 (x)
GET /books/1 (o)
```

* 행위는 HTTP METHOD로 표현한다.
```
GET /books    // book 전체 요청
GET /books/1  // books 중 id 1인 값 요청
POST /books   -> 201 CREATED
DELETE ..     -> 204 NO CONTENT
PUT ..        -> 204 NO CONTENT
```

* 슬래시 구분자는 계층 관계를 나타내는데 사용한다.
```
http://restapi.navercorp.com/books/nobels/ (x) :마지막에 슬래시를 사용하면 안된다.
http://restapi.navercorp.com/books/nobels  (o)
```

* 하이픈(-)은 URI 가독성을 높이는데 사용하고, 밑줄(_)은 사용하지 않는다.

* URI 경로는 소문자로 한다.

* 파일 확장자는 URI에 포함시키지 않는다.
```
GET / members/soccer/345/photo.jpg HTTP/1.1 Host: restapi.example.com Accept: image/jpg   (x)
GET / members/soccer/345/photo HTTP/1.1 Host: restapi.example.com Accept: image/jpg    (o)
```

* 리소스 간의 관계는 아래와 같이 표현적으로 나타낸다.
```
GET /users/{userid}/devices
GET /users/{userid}/likes/devices
```

* HTTP 응답 상태 코드
    * 적절한 응답 코드를 선택해야한다.
```
200 OK
201 CREATE

400 BAD REQUEST
401 UNAUTHORIZED
403 FORBIDDEN
405 METHOD NOT ALLOWED

301 MOVED PERMANENTLY
500 INTERNAL SERVER ERROR
```

### 참고
* https://ko.wikipedia.org/wiki/REST
* https://meetup.toast.com/posts/92
