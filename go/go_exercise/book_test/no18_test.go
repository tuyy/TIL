package book

import (
	"fmt"
	"testing"
)

func TestNo18(t *testing.T) {
	//n5 := &Node{nil, nil, "5"}
	n4 := &Node{nil, nil, "4"}
	n3 := &Node{nil, n4, "3"}
	n2 := &Node{nil, n3, "2"}
	head := &Node{nil, n2, "1"}

	result := solve18n2(head)
	for n := result; n != nil; n = n.Next {
		fmt.Println(n.Value)
	}
}

func solve18n2(head *Node) *Node {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next, even.Next = odd.Next.Next, even.Next.Next
		odd, even = odd.Next, evenHead.Next
	}

	odd.Next = evenHead
	return head
}

func solve18n1(head *Node) *Node {
	odd := &Node{}
	even := &Node{}

	root1, root2 := odd, even

	for head != nil && head.Next != nil {
		odd.Next = head
		even.Next = head.Next
		odd = odd.Next
		even = even.Next
		head = head.Next.Next
	}
	if head != nil {
		odd.Next = head
		even.Next = nil
	}

	odd.Next.Next = root2.Next
	return root1.Next
}
