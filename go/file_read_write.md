### File Read, Write

* Read:
    * os ```func Open(name string) (*File, error)```
    * io/ioutil ```func ReadFile(name string) ([]byte, error```
    * 라인 단위로 읽기는 ```bufio.NewScanner(..)```
    * 한 꺼번에 읽기는 ```ioutil.ReadAll(..)```
* Write:  
    * os ```func Create(name string) (*File, error)```
    * os ```func OpenFile(name string, flag int, perm FileMode) (*File, error)```
    * io/ioutil ```func WriteFile(filename string, data []byte, perm os.FileMode) error```
* https://golang.org/pkg/os/
* https://golang.org/pkg/io/ioutil/

```go
package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

const (
    fileName  = "/Users/yy/tuyy/go/src/github.com/tuyy/logsummary-go/test.txt"
    fileName2 = "/Users/yy/tuyy/go/src/github.com/tuyy/logsummary-go/test2.txt"
    fileName3 = "/Users/yy/tuyy/go/src/github.com/tuyy/logsummary-go/test3.txt"
)

func main() {
    // File open for reading
    f, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    
    // 한꺼번에 읽기
    b, err := ioutil.ReadAll(f)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))
    
    // Line 단위 읽기
    sc := bufio.NewScanner(f)
    for sc.Scan() {
        fmt.Printf("sc.Scan >> %s\n", sc.Text())
    }
    if err = sc.Err(); err != nil {
        panic(err)
    }
    
    // File open 후 바로 읽기
    content, err := ioutil.ReadFile(fileName)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(content))
    
    // File open and write (이미 존재하면 덮어쓰기)
    b2 := []byte(`hello world my name is tu!`)
    err = ioutil.WriteFile(fileName2, b2, 0644)
    if err != nil {
        panic(err)
    }
    
    // File 삭제
    err = os.Remove(fileName2)
    if err != nil {
        panic(err)
    }
    
    // write mode File open (단, 이미 존재하면 truncate)
    f2, err := os.Create(fileName3)
    if err != nil {
        panic(err)
    }
    defer f2.Close()
    fmt.Fprintf(f2, "bye1\n")
    fmt.Fprintf(f2, "bye2\n")
    
    // APPEND File open
    f3, err := os.OpenFile(fileName, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer f3.Close()
    fmt.Fprintf(f3, "\nhello world1\n")
    fmt.Fprintf(f3, "hello world2\n")
}

# Result
hello1
hello2

hello world1

hello world2

hello1
hello2

hello world1
hello world2

```
