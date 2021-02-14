### container/heap
* heap.Interface 를 구현하면 쉽게 사용할 수 있다.
* https://golang.org/pkg/container/heap/

```go
package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	elem := old[n-1]
	*h = old[0 : n-1]
	return elem
}

func main() {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < 10; i++ {
		n := rand.Intn(100)
		heap.Push(h, n)
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	fmt.Println(h)

	heap.Remove(h, 2) // 정렬된 순서 기준으로 지정한 idx 노드가 삭제된다.

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}

### Result
81 87 47 59 81 18 25 40 56 0 
&[0 18 25 56 40 81 47 87 59 81]
0 18 40 47 56 59 81 81 87 

```
