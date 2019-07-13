### cmake 기본
```
cmake_minimum_required(VERSION 3.14)
project(TUReader)

set(CMAKE_CXX_STANDARD 17)

set(SRC_DIR src)
set(SRCS ${SRC_DIR}/main.cpp)

set(TESTS test test/TestHelloWorld.cpp)

set(GTEST_DIR test/googletest)

# CMakeLists.txt 가 존재하는 하위 디렉토리 지정
add_subdirectory(${GTEST_DIR})

# 링크 디렉토리 목록 : {디렉토리} {디렉토리} ..
link_directories(externals/boost/lib)

# 링크 라이브러리 목록 : {boost_iostreams} {stdc++} ..
link_libraries(boost_iostreams boost_serialization)

# include 할 디렉토리 목록 : {디렉토리} {디렉토리} ..
include_directories(externals/boost/include )
include_directories(${GTEST_DIR}/googletest/include)
include_directories(${GTEST_DIR}/googlemock/include)

# 최종 실행 target명과 관련 소스들
add_executable(TUReader ${SRCS})
add_executable(runTestHelloWorld ${TESTS})

# 특정 target에 링크 라이브러리 등록
target_link_libraries(runTestHelloWorld gtest gtest_main)
target_link_libraries(runTestHelloWorld gmock gmock_main)
```

* [cmake 관련 괜찮은 한글 블로그 링크](https://www.tuwlab.com/27260)
