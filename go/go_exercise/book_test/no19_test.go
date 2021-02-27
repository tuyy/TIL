package book

import (
	"fmt"
	"testing"
)

func TestNo19(t *testing.T) {
	n5 := &Node{nil, nil, "5"}
	n4 := &Node{nil, n5, "4"}
	n3 := &Node{nil, n4, "3"}
	n2 := &Node{nil, n3, "2"}
	head := &Node{nil, n2, "1"}

	result := solve19n1(head, 2, 5)
	for n := result; n != nil; n = n.Next {
		fmt.Println(n.Value)
	}
}

func solve19n1(head *Node, m int, n int) *Node {
	// TODO p.239
	return nil
}

func solve19(head *Node, m, n int) *Node {
	var prev, reversePre, reverseLast *Node
	root := head

	no := 1
	for head != nil {
		if no >= m && no <= n {
			next := head.Next
			if no == m {
				reversePre = prev
				reverseLast = head
				head.Next = nil
			} else if no == n {
				head.Next = prev
				reversePre.Next = head
				reverseLast.Next = next
				break
			} else {
				head.Next = prev
			}
			prev = head
			head = next
		} else {
			prev = head
			head = head.Next
		}
		no++
	}

	return root
}
