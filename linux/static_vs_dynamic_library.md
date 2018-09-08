### 정적 vs 동적 라이브러리

#### 정적 라이브러리
* 정적 라이브러리는 최종 사용자의 APP과 연결된 객체 코드를 포함한 후에 실행 될 수 있는 형태로 변경된다.
* 기본적으로 컴파일된 객체 파일의 패키지 형태를 띈다. 예를들면, libjpeg.a나 jpeg.lib 같은 파일을 의미(unix:.a, windows:.lib)
* 10MB 크기의 라이브러리를 추가해도 실질적으로 사용되는 객체만 APP에 복사된다.
* 배포 시점에 라이브러리 충돌 같은 부분은 걱정 할 필요가 없다.
* 라이브러리의 버전을 업데이트 해야한다면, 해당 정적 라이브러리 뿐만 아니라 APP도 다시 배포해야한다.
* http://www.yolinux.com/TUTORIALS/LibraryArchives-StaticAndDynamic.html

##### 리눅스에서 정적 라이브러리 만들기
```C++
g++ -c file1.cpp
g++ -c file2.cpp
g++ -c file3.cpp

// ar명령 -c:저장소생성, -r:제공된 .o파일을 저장소에 추가, -s:저장소에 인덱스생성
ar -crs libmyapi.a file1.o file2.o file3.o

// 컴파일
g++ usercode.cpp -o userapp -L. -lmyapi -static

// 실행
./userapp
```

#### 동적 라이브러리 (= 공유 라이브러리)
* 컴파일 타임에 연결되는 파일들로 최종 사용자의 APP에서 런타임 시에 동적 라이브러리들을 로드 할 수 있도록 최종 사용자의 APP에 배포된다.
* 런타임시 의존관계를 결정하고 필요하면 의존된 리소스를 로드하기 위해 동적 링커를 필요로 한다.
* 리눅스 운영체제의 동적 링커는 ld.so라고 부르고, 맥은 dyld라고 부른다.

##### 리눅스에서 동적 라이브러리 만들기
```C++
g++ -c -fPIC file1.cpp
g++ -c -fPIC file2.cpp
g++ -c -fPIC file3.cpp
g++ -shared -o libmyapi.so -fPIC file1.o file2.o file3.o

// 컴파일 (-static 옵션을 사용해서 정적 라이브러리만을 사용하도록 강제하지 않았다.)
g++ usercode.cpp -o userapp -L. -lmyapi

// 실행
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:.
./userapp
```

### 유용한 리눅스 유틸리티
* libtool, autotools
