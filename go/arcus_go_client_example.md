# Arcus Client Wrapper Connection Test
 arcus_c_client 설치 후 CGO를 이용하여 빌드하여 연결 테스트


* arcus_c_client 설치
    * https://github.com/tuyy/TIL/blob/master/cpp/arcus_c_client_example.md


```go
package arcus

import (
    "log"
    "unsafe"
)

// #cgo CFLAGS: -isystem ../../externals/libarcus/include
// #cgo LDFLAGS: -pthread -lsasl2 -lm -lstdc++ /home1/irteam/tuyy/libarcus/lib/libmemcached.a /home1/irteam/tuyy/libarcus/lib/libmemcachedutil.a /home1/irteam/tuyy/libarcus/lib/libzookeeper_mt.a
// #include <libmemcached/memcached.h>
import "C"

func Connect() {
    logPath := C.CString("arcus.log")
    defer C.free(unsafe.Pointer(logPath))

    flag := C.CString("a")
    defer C.free(unsafe.Pointer(flag))

    logFd := C.fopen(logPath, flag)
    defer C.fclose(logFd)

    mc := C.memcached_create(nil)
    C.arcus_set_log_stream(mc, logFd)

    pool := C.memcached_pool_create(mc, 1, 3)

    arcusHost := C.CString("ncloud.arcuscloud.navercorp.com:17288")
    defer C.free(unsafe.Pointer(arcusHost))

    serviceCode := C.CString("fc33e6ac6fb84f51bbb1412806a8024c")
    defer C.free(unsafe.Pointer(serviceCode))

    rc := C.arcus_pool_connect(pool, arcusHost, serviceCode)
    if rc != C.ARCUS_SUCCESS {
        log.Printf("failed to connect arcus pool. rz:%d\n", rc)
        return
    }

    // TODO close!
    log.Println("SUCCESS")
}


$ go run main.go
2021/06/07 11:20:39 SUCCESS
```
