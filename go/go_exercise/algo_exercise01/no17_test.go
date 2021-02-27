package book

import (
	"fmt"
	"testing"
)

func TestNo17(t *testing.T) {
	n4 := &Node{nil, nil, "4"}
	n3 := &Node{nil, n4, "3"}
	n2 := &Node{nil, n3, "2"}
	head := &Node{nil, n2, "1"}

	result := solve17n3(head)
	for n := result; n != nil; n = n.Next {
		fmt.Println(n.Value)
	}
}

func solve17n3(head *Node) *Node {
	prev := &Node{}
	root := prev
	prev.Next = head

	for head != nil && head.Next != nil {
		b := head.Next
		head.Next = b.Next
		b.Next = head

		prev.Next = b

		head = head.Next
		prev = prev.Next.Next
	}

	return root.Next
}

func solve17n2(head *Node) *Node {
	node := head

	for node != nil && node.Next != nil {
		node.Value, node.Next.Value = node.Next.Value, node.Value
		node = node.Next.Next
	}
	return head
}

func solve17(head *Node) *Node {
	var prev *Node
	node := head
	no := 1
	for node != nil {
		if no%2 == 0 {
			prev.Value, node.Value = node.Value, prev.Value
		}
		prev = node
		node = node.Next
		no++
	}

	return head
}
