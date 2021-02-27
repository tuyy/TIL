package book

import (
	"reflect"
	"testing"
)

func TestNo34(t *testing.T) {
	input := []int{1, 2, 3}
	want := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	rz := solve34n2(input)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve34n2(input []int) [][]int {

	var dfs func(int, map[int]struct{})

	var result [][]int

	dfs = func(idx int, path map[int]struct{}) {
		if len(input) == len(path) {
			l := make([]int, len(path))
			i := 0
			for k, _ := range path {
				l[i] = k
				i++
			}
			result = append(result, l)
		}
		if idx == len(input) {
			return
		}

		for _, num := range input {
			_, ok := path[num]
			if ok {
				continue
			}

			newPath := map[int]struct{}{}
			for k, v := range path {
				newPath[k] = v
			}
			newPath[num] = struct{}{}

			dfs(idx+1, newPath)
		}
	}

	dfs(0, make(map[int]struct{}))
	return result
}

func solve34n1(input []int) [][]int {
	return nil
}

type StackVal2 struct {
	idx    int
	result []int
}

func solve34(input []int) [][]int {
	var result [][]int

	s := []StackVal2{{0, nil}}

	for len(s) > 0 {
		v := s[len(s)-1]
		s = s[:len(s)-1]

		if len(v.result) == len(input) {
			result = append(result, v.result)
			continue
		}

		for i := len(input) - 1; i >= 0; i-- {
			isExists := false
			for _, j := range v.result {
				if j == input[i] {
					isExists = true
					break
				}
			}

			if !isExists {
				copied := make([]int, len(v.result))
				copy(copied, v.result)
				copied = append(copied, input[i])

				s = append(s, StackVal2{v.idx + 1, copied})
			}
		}
	}

	return result
}
