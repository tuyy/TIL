# webpack 이란?
* javascript 모듈화 도구이다.
* 최종 bundle.js 생성

## webpack 설치
``` npm install webpack webpack-cli --save-dev```

## webpack.config.js
```
module.exports = {
    context: __dirname + '/src', // 모듈 파일 폴더
    mode: 'development',
    entry: './app.js',
    output: {
        path: __dirname + '/dist', // 번들 파일 폴더
        filename: 'app.bundle.js' // 번들 파일 이름 규칙
    }
}

실행
* webpack
* webpack --mode=development|production
```


## 참고
* https://webpack.js.org/
