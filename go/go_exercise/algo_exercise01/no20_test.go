package book

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	stack.Push("11")
	stack.Push("22")
	stack.Push("33")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}

func TestNo20(t *testing.T) {
	input := "()[]{}"
	want := true

	rz := solve20n2(input)
	if rz != want {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve20n2(input string) bool {
	table := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	stack := NewStack()
	for _, c := range input {
		if c == '{' || c == '(' || c == '[' {
			stack.Push(string(c))
		} else if c == '}' || c == ')' || c == ']' {
			if table[string(c)] != stack.Pop() {
				return false
			}
		}
	}

	return stack.Len() == 0
}

func solve20n1(input string) bool {
	stack := NewStack()
	for _, c := range input {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(string(c))
			continue
		}

		item := stack.Pop()
		var comp string
		switch c {
		case ')':
			comp = "("
		case ']':
			comp = "["
		case '}':
			comp = "{"
		default:
			continue
		}

		if item != comp {
			return false
		}
	}
	return true
}

type Stack struct {
	arr []string
}

func (s Stack) Len() int{
	return len(s.arr)
}

func (s *Stack) Push(item string) {
	s.arr = append(s.arr, item)
}

func (s *Stack) Pop() string {
	item := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return item
}

func NewStack() *Stack {
	return &Stack{}
}
