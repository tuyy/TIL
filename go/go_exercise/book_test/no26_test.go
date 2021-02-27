package book

import (
	"fmt"
	"testing"
)

type Node26 struct {
	Value      int
	Prev, Next *Node26
}

type MyCircularDeque struct {
	head, tail     *Node26
	maxsize, count int
}

func (d *MyCircularDeque) InsertFront(v int) bool {
	if d.count == d.maxsize {
		return false
	}

	node := &Node26{Value: v}
	if d.head == nil && d.tail == nil {
		d.head, d.tail = node, node
	} else {
		d.head.Prev = node
		d.tail.Next = node

		node.Next = d.head
		node.Prev = d.tail
		d.head = node
	}

	d.count++
	return true
}

func (d *MyCircularDeque) InsertLast(v int) bool {
	if d.count == d.maxsize {
		return false
	}

	node := &Node26{Value: v}
	if d.head == nil && d.tail == nil {
		d.head = node
		d.tail = node
	} else {
		node.Prev = d.tail
		node.Next = d.head
		d.head.Prev = node
		d.tail.Next = node
		d.tail = node
	}

	d.count++
	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.count == 0 {
		return false
	}

	if d.head == d.tail {
		d.head, d.tail = nil, nil
	} else {
		next := d.head.Next
		next.Prev = d.head.Prev
		d.tail.Next = next
		d.head = next
	}

	d.count--
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.count == 0 {
		return false
	}

	if d.head == d.tail {
		d.head, d.tail = nil, nil
	} else {
		prev := d.tail.Prev
		prev.Next = d.head
		d.head.Prev = prev
		d.tail = prev
	}

	d.count--
	return true

}

func (d *MyCircularDeque) GetFront() int {
	if d.head == nil {
		return -1
	} else {
		return d.head.Value
	}
}

func (d *MyCircularDeque) GetRear() int {
	if d.tail == nil {
		return -1
	} else {
		return d.tail.Value
	}
}

func (d MyCircularDeque) IsEmpty() bool {
	return d.count == 0
}

func (d MyCircularDeque) IsFull() bool {
	return d.maxsize == d.count
}

func NewCircularDeque(k int) *MyCircularDeque {
	return &MyCircularDeque{
		maxsize: k,
	}
}

func TestNo26(t *testing.T) {
	deque := NewCircularDeque(7)
	deque.InsertLast(5)
	deque.InsertLast(6)
	deque.InsertLast(7)
	deque.InsertLast(8)
	deque.InsertLast(9)
	deque.InsertLast(10)
	deque.DeleteLast()
	deque.DeleteLast()
	deque.DeleteFront()
	deque.DeleteLast()

	deque.InsertLast(10)
	deque.InsertLast(11)
	deque.InsertLast(12)
	deque.InsertLast(13)
	deque.InsertLast(14)
	deque.InsertLast(15)
	fmt.Println(deque.GetFront())
	fmt.Println(deque.GetRear())
	fmt.Println(deque.IsEmpty())
	fmt.Println(deque.IsFull())
}
