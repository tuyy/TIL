# Modernizr 라이브러리란?
* 사용자가 사용중인 브라우저가 HTML5의 특정 기능을 지원하는지 감지해 주는 역할을 한다.

### 사용법
```javascript
// 만약 브라우저가 placeholder 기능을 사용하지 않는 경우 처리
if(Modernizr.input.placeholder == false){
  /*브라우저가 HTML5 플레이스 홀더 기능을
  지원하지 않을 때 사용할 코드를 넣는다*/
}
```

```CSS
.no-cssreflections .target {
  /* cssreflections를 지원하지 않는 경우 스타일 */
}
.cssreflections .target {
  /* cssreflections를 지원하는 경우 스타일 */
}
```
