package book

import (
	"container/list"
	"fmt"
	"strings"
	"testing"
)

type Node struct {
	Pre, Next *Node
	Value     string
}

type Deque struct {
	Head, Last *Node
	Size       int
}

func (d *Deque) Append(v string) {
	node := &Node{nil, nil, v}
	if d.Head == nil && d.Last == nil {
		d.Head = node
		d.Last = node
	} else {
		tmp := d.Head
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		node.Pre = tmp
		tmp.Next = node
		d.Last = node
	}
	d.Size++
}

func (d *Deque) PrintAll() {
	if d.Head == nil {
		return
	}

	tmp := d.Head
	for tmp.Next != nil {
		fmt.Printf("%s ", tmp.Value)
		tmp = tmp.Next
	}
	fmt.Printf("%s\n", tmp.Value)
}

func (d *Deque) PopFront() string {
	if d.Head == nil {
		return ""
	}

	oldHead := d.Head
	if oldHead == d.Last {
		d.Last = oldHead.Next
	}

	v := oldHead.Value
	d.Head = oldHead.Next
	if d.Head != nil {
		d.Head.Pre = nil
	}
	d.Size--
	return v
}

func (d *Deque) PopBack() string {
	if d.Last == nil {
		return ""
	}

	oldLast := d.Last
	if oldLast == d.Head {
		d.Head = oldLast.Pre
	}

	v := oldLast.Value
	d.Last = oldLast.Pre
	if d.Last != nil {
		d.Last.Next = nil
	}
	d.Size--
	return v
}

func TestDeque(t *testing.T) {
	d := Deque{}
	d.Append("1")
	if d.Size != 1 {
		t.Fatal("invalid maxsize")
	}
	d.Append("2")
	d.Append("3")
	d.Append("4")
	d.Append("5")
	d.PrintAll()

	d.PopFront()
	rz := d.PopFront()
	want := "2"
	if rz != want {
		t.Fatalf("invalid result. rz:%s want:%s", rz, want)
	}

	if d.Size != 3 {
		t.Fatalf("invalid maxsize. rz:%d\n", d.Size)
	}
	d.PopBack()
	rz = d.PopBack()
	want = "4"
	if rz != want {
		t.Fatalf("invalid result. rz:%s want:%s", rz, want)
	}

	d.PopFront()
	d.PopFront()
	d.PrintAll()
}

func TestNo13(t *testing.T) {
	input := "1->2"
	want := false
	//input := "1->2->1"
	//want := true
	//input := "1->2->2->1"
	//want := true

	rz := solver13n2(input)
	if rz != want {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solver13n2(input string) bool {
	l := list.New()
	for _, v := range strings.Split(input, "->") {
		l.PushBack(v)
	}

	fast, slow := l.Front(), l.Front()
	for fast != nil && fast.Next() != nil {
		fast = fast.Next().Next()
		slow = slow.Next()
	}
	cmp := slow
	if fast == nil {
		cmp = slow.Prev()
	}

	for cmp != nil && slow != nil && cmp.Value == slow.Value {
		cmp = cmp.Prev()
		slow = slow.Next()
	}

	return cmp == nil
}

func solver13n1(input string) bool {
	d := Deque{}
	for _, v := range strings.Split(input, "->") {
		d.Append(v)
	}

	for d.Size > 1 {
		if d.PopFront() != d.PopBack() {
			return false
		}
	}

	return true
}
