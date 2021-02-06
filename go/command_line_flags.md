### Command Line Flags
 단순하고 직관적이다.
 
* 타입별로 3가지 방식으로 지정 가능 ```ex) Int(..), IntVal(..), Val(..)```
* https://golang.org/pkg/flag/

```go
package main

import (
    "flag"
    "fmt"
)

type CmdArgs struct {
    num1, num2 int
    str1 string
    isTest1 bool
}

var cmdArgs CmdArgs

func init() {
    // num0 := flag.Int("num0", 0, "test num0")
    
    flag.IntVar(&cmdArgs.num1, "num1", 0, "test num1")
    flag.IntVar(&cmdArgs.num2, "num2", 0, "test num2")
    flag.StringVar(&cmdArgs.str1, "str1", "", "test string")
    flag.BoolVar(&cmdArgs.isTest1, "istest1", false, "test bool")
    
    flag.Parse()
}

func main() {
    fmt.Println("num1:", cmdArgs.num1)
    fmt.Println("num2:", cmdArgs.num2)
    fmt.Println("str1:", cmdArgs.str1)
    fmt.Println("istest1:", cmdArgs.isTest1)
    fmt.Println("flag.Args():", flag.Args()) // 파싱 이후 나머지
}

# Result
yyui-MacBookPro:kafkaReader-go yy$ go run main.go --str1="hello world" -istest1 -num1=33 asd asd ee w
num1: 33
num2: 0
str1: hello world
istest1: true
flag.Args(): [asd asd ee w]

yyui-MacBookPro:kafkaReader-go yy$ go run main.go -h
Usage of /var/folders/x0/5bg51mqs4_v447sjk2vq59680000gn/T/go-build130070483/b001/exe/main:
  -istest1
        test bool
  -num1 int
        test num1
  -num2 int
        test num2
  -str1 string
        test string
```
