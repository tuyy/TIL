package book

import (
	"fmt"
	"testing"
)

// 연결리스트 뒤집기
func TestNo15(t *testing.T) {
	n1 := &Node{nil, nil, "1"}
	n2 := &Node{nil, nil, "2"}
	n3 := &Node{nil, nil, "3"}
	n4 := &Node{nil, nil, "4"}
	n5 := &Node{nil, nil, "5"}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	for n := n1; n != nil; n = n.Next {
		fmt.Println(n.Value)
	}
	fmt.Println("======")
	reverseNode := reverseList(n1)
	for n := reverseNode; n != nil; n = n.Next {
		fmt.Println(n.Value)
	}
	fmt.Println("======")
	reverseNode2 := reverseList2(reverseNode)
	for n := reverseNode2; n != nil; n = n.Next {
		fmt.Println(n.Value)
	}
}

func reverseList2(head *Node) *Node {
	if head == nil {
		return nil
	}

	var prev *Node
	var next *Node
	node := head

	for node != nil {
		next, node.Next = node.Next, prev
		prev, node = node, next
	}

	return prev
}

func reverseList(head *Node) *Node{
	return reverse(head, nil)
}

func reverse(node, prev *Node) *Node {
	if node == nil {
		return prev
	}

	next := node.Next
	node.Next = prev
	return reverse(next, node)
}
