# 12가지 Go Best Practices
* "Best Practice란, 지속적으로 우수한 결과를 보여 주는 기법이나 방법이다."
* Simple, Readable, Maintainable 한 코드를 작성하는 방법
* 링크 : https://talks.golang.org/2013/bestpractices.slide

## 첫 번째, 예시 코드
```go
type Gopher struct {
    Name     string
    AgeYears int
}

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
    err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
    if err == nil {
        size += 4
        var n int
        n, err = w.Write([]byte(g.Name))
        size += int64(n)
        if err == nil {
            err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
            if err == nil {
                size += 4
            }
            return
        }
        return
    }
    return
}
```

### error 중첩을 피해라
 당연한 소리지만, 중첩을 줄이면 가독성이 좋아진다.

```go
func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
	err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err != nil {
		return
	}
	size += 4

	n, err := w.Write([]byte(g.Name))
	if err != nil {
		return
	}
	size += int64(n)

	err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
	if err != nil {
		return
	}
	size += 4

	return
}
```

### 중복을 제거하라 
 중복은 죄악이다. 가능한 무조건 중복을 제거하도록 하자
 
 ```go
type binWriter struct {
    w    io.Writer
    size int64
    err  error
}

func (w *binWriter) Write(v interface{}) {
    if w.err != nil {
        return
    }
    if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
        w.size += int64(binary.Size(v))
    }
}

// 중복이 제거된 코드
func (g *Gopher) WriteTo(w io.Writer) (int64, error) {
    bw := &binWriter{w: w}
    bw.Write(int32(len(g.Name)))
    bw.Write([]byte(g.Name))
    bw.Write(int64(g.AgeYears))
    return bw.size, bw.err
}
```

### Type switch를 사용하라
 interface{}로 받은 값은 Type switch를 이용하자
 
```go
func (bw *BinWriter) Write(v interface{})  {
    if bw.err != nil {
        return
    }
    
    switch x := v.(type) {
    case string:
        bw.Write(len(x))
        bw.Write([]byte(x))
    case int:
        bw.Write(int64(x))
    default:
        if bw.err = binary.Write(bw.w, binary.LittleEndian, v); bw.err == nil {
            bw.size += int64(binary.Size(v))
        }
    }
}
```

### Flush
 IO 작업은 꾀 비싼 작업이다. buffer를 이용해서 한번에 Write하도록 하자
 
```go
type Gopher struct {
    Name     string
    AgeYears int
}

func (g *Gopher) WriteTo(w io.Writer) (int64, error) {
    bw := &BinWriter{w:w}
    bw.Write(g.Name)
    bw.Write(g.AgeYears)
    return bw.Flush()
}

type BinWriter struct {
    w io.Writer
    buf bytes.Buffer
    err error
}

func (bw *BinWriter) Flush() (int64, error) {
    if bw.err != nil {
        return 0, bw.err
    }
    return bw.buf.WriteTo(bw.w)
}

func (bw *BinWriter) Write(v interface{})  {
    if bw.err != nil {
        return
    }
    
    switch x := v.(type) {
    case string:
        bw.Write(len(x))
        bw.Write([]byte(x))
    case int:
        bw.Write(int64(x))
    default:
        bw.err = binary.Write(&bw.buf, binary.LittleEndian, v)
    }
}
```

### Function Adapters
 중복되는 errorHandle 로직 함수로 분리 예제

```go
func init() {
    http.HandleFunc("/", errorHandler(betterHandler))
}

func errorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := f(w, r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            log.Printf("handling %q: %v", r.RequestURI, err)
        }
    }
}

func betterHandler(w http.ResponseWriter, r *http.Request) error {
    if err := doThis(); err != nil {
        return fmt.Errorf("doing this: %v", err)
    }

    if err := doThat(); err != nil {
        return fmt.Errorf("doing that: %v", err)
    }
    return nil
}
```

## 코드 구성하기

### import
 외부 라이브러리 등은 공백으로 구분한다.
 
```go
import (
    "fmt"
    "io"
    "log"

    "golang.org/x/net/websocket"
)
```

TODO..
