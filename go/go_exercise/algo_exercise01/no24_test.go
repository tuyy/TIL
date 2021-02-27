package book

import (
	"fmt"
	"testing"
)

type Queue struct {
	arr []int
}

func (q *Queue) Push(item int)  {
	q.arr = append(q.arr, item)
}

func (q *Queue) Pop() int {
	if len(q.arr) == 0 {
		return -1
	}

	item := q.arr[0]
	q.arr = q.arr[1:]
	return item
}

func (q Queue) Peek() int {
	if len(q.arr) == 0 {
		return -1
	}
	return q.arr[0]
}

func (q Queue) Empty() bool {
	return len(q.arr) == 0
}

func TestNo24(t *testing.T) {
	queue := Queue{}
	queue.Push(10)
	queue.Push(5)
	queue.Push(1)

	fmt.Printf("%d %d %d\n", queue.Pop(), queue.Pop(), queue.Peek())
	fmt.Println(queue.Empty())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}
