# 아커스 C 클라이언트 빌드 및 예제코드
 memcached 클러스터링 버전인 아커스를 빌드하고 예제를 실행한다.
 
* Centos6.xx, Centos7.xx
* https://github.com/naver/arcus-zookeeper
  * 설치버전: arcus-zookeeper-3.5.9-p1
* https://github.com/naver/arcus-c-client
  * 설치버전: arcus-c-client-1.13.0
  
## 빌드

### arcus-zookeeper
```
$ sudo yum install cppunit-dev

$ cd {HOME_ARCUS_ZOOKEEPER}
$ ant compile_jute
$ cd zookeeper-client/zookeeper-client-c
$ autoreconf -if

### prefix에 설치를 원하는 위치를 추가
$ ./configure --prefix=/home1/irteam/tuyy/libs/libzookeeper
$ make -j8
$ make install
```

### arcus-c-client
```
$ cd {HOME_ARCUS_C_CLIENT}
$ ./config/autorun.sh

### prefix에 client 설치 원하는 위치, --enable-zk-integration 은 필수, --with-zookeeper에 주키퍼 설치 위치 추가
### ./configure --help

$ ./configure --prefix=/home1/irteam/tuyy/libs/libarcus --enable-zk-integration --with-zookeeper=/home1/irteam/tuyy/libs/libzookeeper

!!주의!! make 실패시 CPPFLAGS 에 -DTHREADED 추가필요;
$ make -j8
$ make install
```

## C/C++ 클라이언트 예시코드

```c++
#include <iostream>
#include <cstdio>
#include <string>

#include <libmemcached/memcached.h>

void closeAll(memcached_pool_st *pool, memcached_st *mc, FILE *logFile)
{
    if (pool)
    {
        arcus_pool_close(pool);
        memcached_pool_destroy(pool);
    }
    if (mc)
    {
        memcached_free(mc);
    }
    fclose(logFile);
}


bool addKeyValue(memcached_st *mc, const std::string &key, const std::string &value)
{
    memcached_return rc = MEMCACHED_SUCCESS;
    if ((rc = memcached_add(mc, key.c_str(), key.size(), value.c_str(), value.size(), 10, 0)) != MEMCACHED_SUCCESS)
    {
        std::cout << "failed to add memcached. rc:" << memcached_strerror(mc, rc) << std::endl;
        return false;
    }

    return true;
}

memcached_st* fetchNewMcInPool(memcached_pool_st *pool)
{
    memcached_return rc;
    struct timespec waitTimeWhenFull; // TODO 추가

    memcached_st *newMc = memcached_pool_fetch(pool, &waitTimeWhenFull, &rc);
    if (newMc == nullptr)
    {
        std::cout << "failed to fetch memcached in pool. rc:" << rc << std::endl;
    }

    return newMc;
}

bool getValueByKey(memcached_st *mc, const std::string &key, std::string *result)
{
    char *value = nullptr;
    size_t valueSize = 0;
    uint32_t flags = 0;
    memcached_return rc = MEMCACHED_SUCCESS;
    *result = "";
    if ((value = memcached_get(mc, key.c_str(), key.size(), &valueSize, &flags, &rc)) == nullptr)
    {
        if (rc == MEMCACHED_NOTFOUND)
        {
            std::cout << "NOT FOUND" << std::endl;
        }
        else
        {
            std::cout << "failed to get value by key in arcus. rc:" << memcached_strerror(mc, rc) << std::endl;
        }

        return false;
    }

    *result = std::string(value, valueSize);
    free(value);
    return true;
}

bool deleteByKey(memcached_st *mc, const std::string &key)
{
    memcached_return rc = MEMCACHED_SUCCESS;
    if ((rc = memcached_delete(mc, key.c_str(), key.size(), 0)) != MEMCACHED_SUCCESS)
    {
        std::cout << "failed to delete by key in arcus" << std::endl;
        return false;
    }

    return true;
}

int main(void)
{
    FILE *logFile = fopen("arcus.log", "a");

    memcached_st *mc = memcached_create(nullptr);
    arcus_set_log_stream(mc, logFile);

    int maxPoolSize = 3;
    memcached_pool_st *pool = memcached_pool_create(mc, 1, maxPoolSize);

    const std::string serverAddr = "ncloud.arcuscloud.navercorp.com:17288";
    const std::string svcCode = "fc33e6ac6fb84f51bbb1412806a8024c";
    arcus_return_t rc = arcus_pool_connect(pool, serverAddr.c_str(), svcCode.c_str());
    if (rc != ARCUS_SUCCESS)
    {
        closeAll(pool, mc, logFile);
        std::cout << "failed to connect pool. rc:" << rc << std::endl;
        return -1;
    }

    memcached_st *newMc = fetchNewMcInPool(pool);
    if (newMc == nullptr)
    {
        closeAll(pool, mc, logFile);
        return -1;
    }

    const std::string key("HELLO_TEST_KEY");
    const std::string value("HELLO_TEST_VALUE");

    if (!addKeyValue(newMc, key, value))
    {
        closeAll(pool, mc, logFile);
        return -1;
    }
    std::cout << "ADD: SUCCESS" << std::endl;

    std::string result;
    if (!getValueByKey(newMc, key, &result))
    {
        closeAll(pool, mc, logFile);
        return -1;
    }
    std::cout << "GET:" << result << std::endl;

    if (!deleteByKey(newMc, key))
    {
        closeAll(pool, mc, logFile);
        return -1;
    }
    std::cout << "DEL: SUCCESS" << std::endl;

    getValueByKey(newMc, key, &result);
    std::cout << "GET:" << result << std::endl;

    // close All
    closeAll(pool, mc, logFile);
    std::cout << "FINISHED!" << std::endl;

    return 0;
}

## 실행
### !항상 main 함수를 먼저 쓰자
$ g++ -std=c++11 -Wall -g -Ilibarcus/include main.cpp libarcus/lib/libmemcached.a libarcus/lib/libmemcachedutil.a libarcus/lib/libzookeeper_mt.a -pthread -lsasl2 -o main;./main

ADD: SUCCESS
GET:HELLO_TEST_VALUE
DEL: SUCCESS
NOT FOUND
GET:
FINISHED!
```