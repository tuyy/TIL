## babel 이란?
* 자바스크립트 코드 ES6+ to ES5로 변환시켜주는 트랜스파일러이다.

### 설치
```
npm install --save-dev @babel/core @babel/cli @babel/preset-env
```

### 실행
```
package.json

"scripts": {
  "build": "babel src --out-dir lib",
  "test": "echo \"Error: no test specified\" && exit 1"
},

src 에 ES6 문법으로 코드 작성
lib 에 결과가 나옴


// .babelrc 바벨 설정파일
{
    "presets": ["@babel/preset-env"]
}
```
