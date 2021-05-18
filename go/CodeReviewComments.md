# Go Code ReviewComments
 Go 코드 작성시 참고할 코딩 컨벤션

* https://github.com/golang/go/wiki/CodeReviewComments

## Error Strings
 에러 문자열은 대문자로 시작하거나 구둣점으로 끝나면 안된다. go의 에러 핸들링 스타일 때문에 보통 여러 에러 문자열이 결합되기 때문이다.
 

## Context
 context.Context 는 함수의 제일 앞에 위치해야한다.
 
## 비어있는 Slice 선언하기
 아래 방법 중 첫 번째 방식을 선호하라
 
```go
var t []string // json 변환시 null로 인식
or
t := []string{} // json 변환시 []로 인식
```

## Indent Error Flow

```go
if err != nil {
  // error handling
} else {
  // normal code
}

// instead, write:

if err != nil {
  // error handling
}
// normal code
```

## Initialisms
 축약어(?)는 대문자로 시작해서 소문자로 쓰지 말자
 
 ```go
//  Url, Http. (X)
//  URL, url, HTTP, http (o)
 ```
 
## 변수 이름
 지역변수명은 짧게 써라. 범위가 넓어 질 수록 좀 더 설명적인 이름을 사용하라
 
## 상수 이름
 상수명은 ```maxLength``` 이렇게 써라.
 
```
// MaxLength or MAX_LENGTH (X)
```

## Getter, Setter

```go
type Object struct {
  owner Owner
}
owner := obj.Owner()
if owner != user {
    owner.SetOwner(user)
}
```
