# bower 란?
* 웹 프론트엔드 패키지 관리자
# bower.json 파일을 가지고 패키지를 관리한다.

### 설치 및 사용법
```
npm install -g bower

bower --version

bower install jquery
bower uninstall jquery
bower list
bower search jquery

bower install qunit --save-dev
bower install jquery --save

bower uninstall qunit --save-dev
bower uninstall jquery --save

// bower_components 하위에 bower.json파일에 명시된 의존성 패키지가 생성된다.
{current_project}/bower_components
``
