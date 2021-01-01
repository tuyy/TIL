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

* 실제 사용중인 예시
    * .a, .so 상관없이 link_libraries & link_directories 에 등록해야한다.
```
cmake_minimum_required(VERSION 3.14)
project(collect)

set(CMAKE_CXX_STANDARD 17)

set(EXT_DIR externals)
set(CARBON_DIR   ${EXT_DIR}/carbon)
set(NMQUEUE_DIR  ${EXT_DIR}/nmqueue)
set(BOOST_DIR    ${EXT_DIR}/boost)
set(LIBMIME_DIR  ${EXT_DIR}/libmime)
set(LIBARCUS_DIR ${EXT_DIR}/libarcus)
set(OWFS_DIR     ${EXT_DIR}/owfs)
set(LIBQEX_DIR   ${EXT_DIR}/libqueryexecutor)
set(SYNAP_DIR    ${EXT_DIR}/synapfilter)
set(GTEST_DIR    ${EXT_DIR}/googletest-master)

set(MODEL_DIR model)

include_directories(SYSTEM
include
${MODEL_DIR}/include
${CARBON_DIR}/include
${NMQUEUE_DIR}/include
${BOOST_DIR}/include
${LIBMIME_DIR}/include
${LIBARCUS_DIR}/include
${LIBARCUS_DIR}/include/c-client-src
${OWFS_DIR}/include
${LIBQEX_DIR}/include
${LIBQEX_DIR}/include/nbase)

set(CMAKE_CXX_FLAGS "-g -O0 -Wall -Werror -fpic -Wno-non-virtual-dtor -pthread")

link_libraries(stdc++
boost_serialization
boost_iostreams
boost_filesystem
mime
icuuc
icudata
icui18n
nsUniversalCharDet
nmqueue
owfsclient
memcached
memcachedutil
zookeeper_mt
queryexecutor
nbaseclient)

link_directories(${EXT_DIR}/boost/lib
${LIBMIME_DIR}/lib
${LIBMIME_DIR}/externals/chardet/lib
${LIBMIME_DIR}/externals/icu/lib
${LIBMIME_DIR}/externals/opendkim/lib
${NMQUEUE_DIR}/lib
${OWFS_DIR}/lib
${LIBARCUS_DIR}/lib
${LIBQEX_DIR}/lib)

file(GLOB ALL_SRCS
        model/src/model/*.cpp
        model/src/model/index/*.cpp
        src/*.cpp
        src/action/*.cpp
        src/arcus/*.cpp
        src/http/*.cpp
        src/indexing/*.cpp
        src/mon/*.cpp
        src/npc/*.cpp
        src/owfs/*.cpp
        src/queue/*.cpp
        src/synap/*.cpp
        src/util/*.cpp
        src/workers/*.cpp)

ADD_LIBRARY(collect SHARED ${ALL_SRCS})

# unittest

set(GTEST_DIR ${EXT_DIR}/googletest-master)

set(GTEST_SRCS
        unittest/src/main.cpp
        unittest/src/DataMover.cpp)

add_subdirectory(${GTEST_DIR})

ADD_EXECUTABLE(unittest ${GTEST_SRCS})

target_include_directories(unittest PUBLIC ${GTEST_DIR}/googletest/include)

target_link_libraries(unittest gtest)

SET_TARGET_PROPERTIES(unittest PROPERTIES COMPILE_FLAGS "-g -Wall -Werror -fpic -Wno-non-virtual-dtor -lstdc++ -pthread")

```
* 예시2
```
cmake_minimum_required(VERSION 3.14)
project(messenger_search)

set(CMAKE_CXX_STANDARD 17)

set(EXT_DIR externals)
set(CARBON_DIR       ${EXT_DIR}/carbon)
set(ICU_DIR          ${EXT_DIR}/icu)
set(UCHARDET_DIR     ${EXT_DIR}/uchardet)
set(NMQUEUE_DIR      ${EXT_DIR}/nmqueue)
set(BOOST_DIR        ${EXT_DIR}/boost)
set(JSON_DIR         ${EXT_DIR}/json)
set(SYNAP_DIR        ${EXT_DIR}/synapfilter)
set(OWFS_DIR         ${EXT_DIR}/owfs)
set(LIBZIPPP_DIR     ${EXT_DIR}/libzippp)
set(JEMALLOC_DIR     ${EXT_DIR}/libjemalloc)
set(GOOGLE_TEST_DIR  ${EXT_DIR}/googletest-master)

set(MODEL_DIR model)

include_directories(SYSTEM
    include
    ${CARBON_DIR}/include
    ${MODEL_DIR}/include
    ${BOOST_DIR}/include
    ${ICU_DIR}/include
    ${UCHARDET_DIR}/include
    ${SYNAP_DIR}/include
    ${OWFS_DIR}/include
    ${LIBZIPPP_DIR}/lib/libzip-1.0.1/lib
    ${LIBZIPPP_DIR}/src
    ${JSON_DIR}/include
    ${NMQUEUE_DIR}/include)

set(CMAKE_CXX_FLAGS "-g -O0 -Wall -Werror -fpic -Wformat-security -Wformat-nonliteral  -Wformat=2 -fno-omit-frame-pointer -Wno-non-virtual-dtor -DGNU_SOURCE -D_REENTRANT -D_DEBUG -D__STDC_LIMIT_MACROS -Wl,--wrap=__cxa_throw")

link_libraries(nmqueue
    boost_system
    boost_serialization
    boost_thread
    boost_chrono
    boost_filesystem
    boost_iostreams
    boost_locale
    boost_regex
    pthread
    owfsclient
    crypto
    ssl
    snf
    uchardet
    icuuc
    icudata
    icui18n
    zippp
    zip)

link_directories(${BOOST_DIR}/lib
    ${NMQUEUE_DIR}/lib
    ${OWFS_DIR}/lib
    ${LIBZIPPP_DIR}
    ${LIBZIPPP_DIR}/lib/libzip-1.0.1/lib/.libs
    ${UCHARDET_DIR}/lib64
    ${SYNAP_DIR}/lib
    ${ICU_DIR}/lib)

file(GLOB ALL_SRCS
    model/src/model/*.cpp
    model/src/model/archiver/*.cpp
    src/*.cpp
    src/action/*.cpp
    src/connector/*.cpp
    src/connector/npc/*.cpp
    src/owfs/*.cpp
    src/queue/*.cpp
    src/util/*.cpp
    src/zip/*.cpp)

ADD_LIBRARY(messenger_search SHARED ${ALL_SRCS})

# unittest

set(GTEST_DIR ${EXT_DIR}/googletest-master)

set(GTEST_SRCS
    test/BotNameTest.cpp
    test/CacheTest.cpp
    test/CharDetTest.cpp
    test/ChronoTest.cpp
    test/CoordinatorTest.cpp
    test/ExtractTest.cpp
    test/FileAppenderTest.cpp
    test/FileTextFilterTest.cpp
    test/GeneratorTest.cpp
    test/GetMessageDirection.cpp
    test/GetTenantId.cpp
    test/GetUserDetail.cpp
    test/GzipTest.cpp
    test/JsonTest.cpp
    test/MessageParserTest.cpp
    test/MutexTest.cpp
    test/NPCIOTest.cpp
    test/NormalizeTest.cpp
    test/OwfsTest.cpp
    test/QueueTest.cpp
    test/TestMain.cpp
    test/UtilTest.cpp
    test/UuidInfoTest.cpp)

add_subdirectory(${GTEST_DIR})

ADD_EXECUTABLE(unittest ${GTEST_SRCS})

target_include_directories(unittest PUBLIC ${GTEST_DIR}/googletest/include ${GTEST_DIR}/googlemock/include)

target_link_directories(unittest PUBLIC ${CARBON_DIR}/lib)

target_link_libraries(unittest gtest gmock messenger_search carbon aprutil-1 apr-1)

SET_TARGET_PROPERTIES(unittest PROPERTIES COMPILE_FLAGS "-g -Wall -Werror -fpic -Wno-non-virtual-dtor -lstdc++ -pthread -Wl,--wrap=__cxa_throw")
```

* [cmake 관련 괜찮은 한글 블로그 링크](https://www.tuwlab.com/27260)
