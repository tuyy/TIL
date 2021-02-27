package book

import (
	"fmt"
	"testing"
)

type CircularQueue struct {
	front, rear int
	arr         []int
	maxsize     int
}

func (q *CircularQueue) EnQueue(v int) bool {
	next := (q.rear + 1) % q.maxsize
	if next == q.front {
		return false
	} else {
		q.arr[next] = v
		q.rear = next
		return true
	}
}

func (q CircularQueue) Rear() int {
	if q.rear == -1 {
		return -1
	} else {
		return q.arr[q.rear]
	}
}

func (q CircularQueue) IsFull() bool {
	if q.front == q.rear {
		return true
	} else {
		return false
	}
}

func (q *CircularQueue) DeQueue() bool {
	if q.front == q.rear {
		return false
	}

	q.front++
	q.arr[q.front] = -1
	return true
}

func NewCircularQueue(maxsize int) *CircularQueue {
	return &CircularQueue{
		maxsize: maxsize,
		arr:     make([]int, maxsize),
	}
}

func TestNo25(t *testing.T) {
	que := NewCircularQueue(5)
	fmt.Println(que.EnQueue(10))
	fmt.Println(que.EnQueue(20))
	fmt.Println(que.EnQueue(30))
	fmt.Println(que.EnQueue(40))
	fmt.Println(que.Rear())
	fmt.Println(que.IsFull())
	fmt.Println(que.DeQueue())
	fmt.Println(que.DeQueue())
	fmt.Println(que.DeQueue())
	fmt.Println(que.DeQueue())
	fmt.Println(que.DeQueue())
}
