# gulp 란?
* 자바의 메이븐과 비슷한 FE용 빌드, 테스트 자동화 도구
* gulpfile.js 에 빌드, 테스트 관련 내용을 작성해야한다.

### 설치
* 기본 예시
```
npm install init
npm install gulp

// gulpfile.js 추가
gulp --tasks
gulp
```

### Getting Started (gulpfile.js)
#### Create Taksk
```
const { series, parallel } = require('gulp');

function javascript(cb) {
  cb();
}

function css(cb) {
  cb();
}

function clean(cb) {
  // body omitted
  cb();
}

exports.default = parallel(javascript, css);  // gulp 입력시 javascript, css 동시 실행
exports.mycmd = series(css, javascript);  // gulp mycmd 입력시 css, javascript 함수 순으로 순차 실행
exports.build = series(clean, parallel(css, javascript));  // gulp build 입력시 clean 실행 후 css, javascript 동시 실행

===================================================================

taeunui-MacBook-Pro:yeoman_basic_test tuyy$ gulp --tasks
[11:02:14] Tasks for ~/tuyy/study/project/yeoman_basic_test/gulpfile.js
[11:02:14] ├─┬ mycmd
[11:02:14] │ └─┬ <series>
[11:02:14] │   ├── css
[11:02:14] │   └── javascript
[11:02:14] └─┬ default
[11:02:14]   └─┬ <parallel>
[11:02:14]     ├── javascript
[11:02:14]     └── css
```

### 참고
* https://gulpjs.com/docs/en/getting-started/quick-start
