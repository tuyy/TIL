package book

import (
	"fmt"
	"testing"
)

type MyNode struct {
	Next  *MyNode
	Value int
}

// TODO 나중에 병합정렬 이해하면서 다시 풀어보기
func TestNo14(t *testing.T) {
	n1 := &MyNode{nil, 4}
	n2 := &MyNode{n1, 2}
	n3 := &MyNode{n2, 1}

	n4 := &MyNode{nil, 4}
	n5 := &MyNode{n4, 3}
	n6 := &MyNode{n5, 1}

	rz := mergeTwoLists(n3, n6)

	for i := rz; i != nil; i = i.Next {
		fmt.Printf("%d ", i.Value)
	}
	fmt.Println()
}

func mergeTwoLists(n1 *MyNode, n2 *MyNode) *MyNode {
	if (n1 == nil) || (n2 != nil && n1.Value > n2.Value) {
		n1, n2 = n2, n1
	}
	if n1 != nil {
		n1.Next = mergeTwoLists(n1.Next, n2)
	}
	return n1
}
