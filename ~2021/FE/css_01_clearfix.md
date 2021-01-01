# clearfix 란?

```html
<!DOCTYPE html>
<html lang="kr">
  <head>
    <meta charset="utf-8">
    <title>Hello Test</title>
    <link rel="stylesheet" href="index.css">
  </head>
  <body>
    <div class="container">
        <div class="box box1">
          first box
        </div>

        <div class="box inner-box">
          <div class="box2">
            inner box1
          </div>

          <div class="box3">
            inner box2
          </div>
        </div>

        <div class="box box4">
          <p style="margin:0;">위 block이 두번째 row에 포함되는 버그를 해결하는 것이 clearfix이다.</p>
        </div>
    </div>
  </body>
</html>
```

* 해결방법
    * float을 사용하는 태그의 부모요소에 ```..::after {clear: both; content: ""; display: block;}``` 을 추가한다.
```css
.container {
  background-color: #dedede;
  height: 100px;
}

.box {
  height: 100px;
}

.box1 {
  background-color: #f29f21;
}

.box2 {
  float:left;
  width: 20%;
  height: 100%;
  background-color: yellow;
}

.box3 {
  float: right;
  width: 20%;
  height: 100%;
  background-color: red;
}

.box4 {
  background-color: green;
}

.inner-box {
  background-color: blue;
}

.inner-box::after {  /* 이것이다. */
  content: "";
  display: block;
  clear: both;
}
```
