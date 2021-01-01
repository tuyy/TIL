#### autotools?
* GNU 빌드 시스템(GNU build system), 또는 간단히 Autotools는 소스 코드 패키지를 수많은 유닉스 계열 운영 체제로 포팅할 수 있게 도와 주는 프로그래밍 도구이다.
* C컴파일러는 시스템마다 다르기 때문에, 여러 운영체제 포팅을 위해 도움이 필요하다.
```
./configure // 현재 시스템 분석(필요한 라이브러리가 있으면 사용자에게 알림)하여 Makefile 생성
make  // 실제 소스 컴파일
sudo make install  // make 명령어로 빌드된 바이너리와 라이브러리를 prefix(confiture 설정으로 지정가능)로 지정된 디렉토리에 설치
```

#### Autotools를 적용하려면?
* 아래 두 파일 작성(or 복사) 필요
  * Makefile.am  // automake의 인자로 필요
  * configure.ac  // autoconf의 인자로 필요(autoscan으로 생성하여 약간의 수정으로 사용할 수 있다.)
  
#### 과정
```
// 대략적인 과정
1. Makefile.am 작성
2. autoscan 명령어 입력 후 configure.ac 로 mv, 내용 약간 수정
3. aclocal, autoheader, autoconf 입력
4. automake --foreign --add-missing --copy 입력  // 최종 configure 파일 생성됨
5. 이후..
    1) ./configure
    2) make
    3) sudo make install
```

#### 참고
* https://ko.wikipedia.org/wiki/GNU_%EB%B9%8C%EB%93%9C_%EC%8B%9C%EC%8A%A4%ED%85%9C
