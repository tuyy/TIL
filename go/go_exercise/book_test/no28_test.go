package book

import (
	"fmt"
	"testing"
)

func TestNo28(t *testing.T) {
	m := NewHashMap()
	m.Put(10, 10) // 3
	m.Put(20, 20) // 3
	fmt.Println(m.Get(10))
	fmt.Println(m.Get(20))
	m.Remove(10)
	fmt.Println(m.Get(10))
	fmt.Println(m.Get(20))
	m.Remove(20)
	fmt.Println(m.Get(10))
	fmt.Println(m.Get(20))
}

func (m HashMap) Get(k int) int {
	result := m.arr[m.hash(k)]
	if result != nil {
		for result != nil {
			if result.Key == k {
				return result.Value
			}
			result = result.Next
		}
	}

	return -1
}

type Node28 struct {
	Key   int
	Value int
	Next  *Node28
}

type HashMap struct {
	arr        []*Node28
	bucketSize int
}

func (m *HashMap) Put(k int, v int) {
	hashKey := m.hash(k)

	n := &Node28{Key: k, Value: v}
	if m.arr[hashKey] == nil {
		m.arr[hashKey] = n
	} else {
		last := m.arr[hashKey]
		var prev *Node28
		for last != nil {
			if last.Key == k {
				last.Value = v
				return
			}
			prev = last
			last = last.Next
		}
		prev.Next = n
	}
}

func (m HashMap) hash(k int) int {
	return (k*3 + 1*3) % m.bucketSize
}

func (m *HashMap) Remove(k int) {
	node := m.arr[m.hash(k)]
	if node == nil {
		return
	}

	var prev *Node28
	for node != nil {
		if node.Key == k {
			if prev == nil {
				m.arr[m.hash(k)] = node.Next
				return
			} else {
				prev.Next = node.Next
				break
			}
		}
		prev = node
		node = node.Next
	}
}

const bucketSize = 10

func NewHashMap() *HashMap {
	return &HashMap{
		bucketSize: bucketSize,
		arr:        make([]*Node28, bucketSize),
	}
}
