package book

import (
	"fmt"
	"testing"
)

// push, pop, top, empty

type QStack struct {
	que []int
}

func (qs *QStack) Push(item int) {
	qs.que = append(qs.que, item)
}

func (qs *QStack) Pop() int {
	if len(qs.que) == 0 {
		return -1
	}

	item := qs.que[len(qs.que)-1]
	qs.que = qs.que[:len(qs.que)-1]
	return item
}

func (qs *QStack) Top() int {
	if len(qs.que) == 0 {
		return -1
	}

	return qs.que[len(qs.que)-1]
}

func (qs *QStack) Empty() bool {
	return len(qs.que) > 0
}

func TestNo23(t *testing.T) {
	stack := QStack{}
	stack.Push(10)
	stack.Push(5)
	fmt.Printf("%d %d\n", stack.Pop(), stack.Pop())
	fmt.Println(stack.Empty())
}
