package book

import (
	"container/heap"
	"fmt"
	"testing"
)

// 힙으로 풀어봐야한다.
func TestNo27(t *testing.T) {
	arr := []int{1, 2, 3, 4, 1, 9, 3, 2, 5}
	SortInsert(arr)

	a := []int{1, 4, 5}
	b := []int{1, 3, 4}
	c := []int{2, 6}

	var result []int
	result = append(result, a...)
	result = append(result, b...)
	result = append(result, c...)

	SortInsert(result)

	fmt.Println(result)
}

func SortInsert(arr []int) {
	for i := 1; i < len(arr); i++ {
		k := arr[i]
		for j := i - 1; j >= 0; j-- {
			if k >= arr[j] {
				break
			}
			arr[j+1], arr[j] = arr[j], arr[j+1]
		}
	}
}

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
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func TestNo27_1(t *testing.T) {
	h := &IntHeap{}
	heap.Init(h)

	input := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	for _, ints := range input {
		for _, v := range ints {
			fmt.Printf("%d ", v)
			heap.Push(h, v)
		}
	}
	fmt.Println()
	fmt.Println(h.Len())
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
