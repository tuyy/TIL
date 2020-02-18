### Make 기본 예제

#### 1. 가장 기본적인 Makefile 작성
```c
test : main.o read.o write.o
    g++ -o main main.o read.o write.o

main.o : io.h main.cpp
    g++ -c main.cpp

read.o : io.h read.cpp
    g++ -c read.cpp

write.o : io.h write.cpp
    g++ -c write.cpp
```

#### 2. 매크로와 레이블 추가
```c
TARGET=test
OBJS=main.o read.o write.o

$(TARGET) : $(OBJS)
    g++ -o $(TARGET) $(OBJS)

main.o : io.h main.cpp
    g++ -c main.cpp

read.o : io.h read.cpp
    g++ -c read.cpp

write.o : io.h write.cpp
    g++ -c write.cpp

clean :
    rm $(OBJS)
````

#### 3. Pre defined macro (정의된 매크로)
* 기본적으로 make는 쉘 상에 정의된 환경변수를 그대로 쓴다.
* ```make -p```
* https://www.gnu.org/software/make/manual/html_node/Automatic-Variables.html
* CC, CXX, CPPFLAGS, RM, CFLAGS ...
```c
TARGET=test
OBJS=main.o read.o write.o

CC=g++ # pre-defined
CPPFLAGS=-Wall -std=c++14 # pre-defined

$(TARGET) : $(OBJS)
    $(CC) -o $@ $(OBJS)

main.o : io.h main.cpp
read.o : io.h read.cpp
write.o : io.h write.cpp

clean :
    $(RM) $(OBJS) $(TARGET)
```

#### 4. Suffix rule (확장자 규칙)
* 확장자를 보고 적절한 규칙을 수행시키는 규칙이다.
* ```.SUFFIXES : .cpp .o``` : .cpp to .o
```c
.SUFFIXES : .cpp .o

TARGET=test
OBJS=main.o read.o write.o

CC=g++ # pre-defined
CPPFLAGS=-Wall -std=c++14 # pre-defined

$(TARGET) : $(OBJS)
    $(CC) -o $@ $(OBJS)

main.o : io.h main.cpp
read.o : io.h read.cpp
write.o : io.h write.cpp

clean :
    $(RM) $(OBJS) $(TARGET)
```

#### 5. Internal macro (내부 매크로)
```
$* : 확장자가 없는 현재의 목표 파일(Target)
$@ : 현재의 목표 파일(Target)
$< : 현재의 목표 파일(Target)보다 더 최근에 갱신된 파일 이름
$? : 현재의 목표 파일(Target)보다 더 최근에 갱신된 파일이름
```

#### 6. TIPS
* 긴 문자 개행하기
```
OBJS = shape.o \
rectangle.o \
circle.o \
line.o \
bezier.o 
```

#### 7. 기본 Makefile form
```
.SUFFIXES : .cpp .o

CUR_DIR:=$(shell pwd)

CXX=g++

INC_DIR=$(CUR_DIR)/include
INC_DIRS=$(realpath $(INC_DIR))
INC_DIR_OPTION=$(addprefix -isystem, $(INC_DIRS))

SRC_LIST:=$(sort $(shell find -L src -name '*.cpp'))
SRCS:= $(SRC_LIST:./%=%)
OBJS=$(SRCS:.cpp=.o)

CPPFLAGS=-O2 -Wall -Werror -std=c++14 $(INC_DIR_OPTION)

TARGET=test

$(TARGET) : $(OBJS)
    $(CXX) -o $@ $(CPPFLAGS) $(OBJS)

clean :
    $(RM) $(OBJS) $(TARGET)
```

#### 8. 임시로 쓰자
```
.PHONY: all clean check

CXX=g++

CXXFLAGS=-std=c++14 -Wall -Werror -O0

TARGET=main

SRCS_CPP=$(wildcard *.cpp)
OBJS:=$(SRCS_CPP:.cpp=.o)

$(TARGET) : $(OBJS)
    $(CXX) $(CXXFLAGS) -o $@ $(OBJS)

all : $(TARGET)

check : $(TARGET)
    ./$(TARGET)

clean :
    $(RM) $(OBJS) $(TARGET)

```

#### 9. example
```
.SUFFIXES : .cpp .o

CXX=g++
CUR_DIR:=$(shell pwd)
EXT_DIR      := $(CUR_DIR)/externals
CARBON_DIR   := $(EXT_DIR)/carbon
BOOST_DIR    := $(EXT_DIR)/boost

INC_DIR=$(CUR_DIR)/include
INC_DIRS=$(realpath $(INC_DIR)) $(realpath $(CARBON_DIR)/include) $(realpath $(BOOST_DIR)/include)
INC_DIR_OPTION=$(addprefix -isystem, $(INC_DIRS)) $(addprefix -I, $(INC_DIRS))

SRC_LIST:=$(sort $(shell find -L src -name '*.cpp'))
SRCS:= $(SRC_LIST:./%=%)
OBJS=$(SRCS:.cpp=.o)

LIB_DIR_OPTION := -L$(BOOST_DIR)/lib

DLIBS := -lpthread \
         -lboost_system \
         -lcrypto \
         -lssl

SLIBS := $(CARBON_DIR)/lib/libcarbon.a

CPPFLAGS=-O2 -Wall -Werror -std=c++17 -g $(INC_DIR_OPTION) $(LIB_DIR_OPTION) $(DLIBS) $(SLIBS)

TARGET=main

$(TARGET) : $(OBJS)
    $(CXX) -o $@ $(CPPFLAGS) $(OBJS)

clean :
    $(RM) $(OBJS) $(TARGET)
```
