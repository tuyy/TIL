# media query 란?
* 디바이스별 출력폭에 맞게 css를 처리할 수 있다.
```css
@media (max-width: 400px) { /* 논리 연산이 가능하다. and or not ..*/
  .nav {
    float: none;
    width: 100%;
  }

  .contents {
    float: none;
    width: 100%;
  }
}
```
* 참고
    * https://developer.mozilla.org/ko/docs/Web/Guide/CSS/Media_queries
