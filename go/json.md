### json
 TODO 작성중

* https://golang.org/pkg/encoding/json/

```go
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "os"
    "time"
)

type Person struct {
    Name     string    `json:"name"`
    Age      int       `json:",omitempty"` // 0이면 무시, `json:"-"` <- 무시
    Id       int       `json:",string"`
    BirthDay time.Time `json:"birthday"`
}

func main() {
    // 직렬화. struct to json([]byte)
    j, err := json.Marshal(Person{Name: "tu", Age: 0, Id: 777, BirthDay: time.Now()})
    if err != nil {
        panic(err)
    }
    
    buf := bytes.NewBuffer(j)
    fmt.Fprintln(os.Stdout, buf)
    fmt.Println(string(j))
    
    // 역직렬화. json([]byte) to struct
    p := Person{}
    err = json.Unmarshal(j, &p)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("%#v\n", p)
}

# Result
{"name":"tu","Id":"777","birthday":"2021-02-04T00:34:17.61746+09:00"}
{"name":"tu","Id":"777","birthday":"2021-02-04T00:34:17.61746+09:00"}
main.Person{Name:"tu", Age:0, Id:777, BirthDay:time.Time{wall:0x24cdb120, ext:63747963257, loc:(*time.Location)(0x11bd460)}}
```
