package book

import (
	"reflect"
	"testing"
)

func TestNo22(t *testing.T) {
	input := []int{73, 74, 75, 71, 69, 72, 76, 73}
	want := []int{1, 1, 4, 2, 1, 1, 0, 0}

	rz := solve22n1(input)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

// TODO 풀이가 이해가 안간다.
func solve22n1(input []int) []int {
	answer := make([]int, len(input))
	var stack []int
	for i, v := range input {
		for len(stack) > 0 && v > input[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[last] = i - last
		}
		stack = append(stack, i)
	}
	return answer
}

type Stack2 struct {
	arr []int
}

func (s *Stack2) Push(v int) {
	s.arr = append(s.arr, v)
}

func (s *Stack2) Last() int {
	if s.Size() == 0 {
		return 0
	} else {
		return s.arr[len(s.arr)-1]
	}
}

func (s *Stack2) Pop() int {
	if s.Size() == 0 {
		return 0
	}

	item := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return item
}

func (s *Stack2) Size() int {
	return len(s.arr)
}

func NewStack2() *Stack2 {
	return &Stack2{}
}

func solve22(input []int) []int {
	rz := make([]int, len(input))
	k := 0
	for i, v := range input {
		step := 0
		for j := i; j < len(input); j++ {
			if v < input[j] {
				rz[k] = step
				k++
				break
			}
			step++
		}
	}
	return rz
}
