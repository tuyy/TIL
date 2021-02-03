### 정적 링킹
* 일반적인 라이브러리 include해서 사용하는 방법
* 컴파일 타임에 라이브러리를 링킹하며 실행파일 크기가 커진다.

### 동적 링킹
* .so 파일을 LD_LIBRARY_PATH 에 등록해두는 방법
* 매번 컴파일을 할 필요가 없다.