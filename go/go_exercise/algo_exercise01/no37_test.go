package book

import (
	"reflect"
	"testing"
)

func TestNo37(t *testing.T) {
	nums := []int{1,2,3}
	want := [][]int{
		{3},
		{1},
		{2},
		{1,2,3},
		{1,3},
		{2,3},
		{1,2},
		{},
	}

	rz := solve37n2(nums)
	if !reflect.DeepEqual(rz, want) {
	    t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}

}

func solve37n2(nums []int) [][]int {
	var result [][]int

	var dfs func([]int)

	dfs = func(path []int) {
		for _, num := range nums {
			if path != nil && path[len(path)-1] >= num {
				continue
			}

			copied := make([]int, len(path))
			copy(copied, path)
			copied = append(copied, num)

			dfs(copied)
		}
		result = append(result, path)
	}

	dfs(nil)
	return result
}

func solve37n1(nums []int) [][]int {
	var result [][]int

	s := []StackVal37{{0, nil}}

	for len(s) > 0 {
		v := s[len(s)-1]
		s = s[:len(s)-1]

		result = append(result, v.path)

		for i := v.idx; i < len(nums); i++ {
			copied := make([]int, len(v.path))
			copy(copied, v.path)
			copied = append(copied, nums[i])

			s = append(s, StackVal37{i+1, copied})
		}
	}

	return result
}

type StackVal37 struct {
	idx int
	path []int
}

func solve37(nums []int) [][]int {
	var result [][]int

	s := []StackVal37{{0, nil}}

	for len(s) > 0 {
		v := s[len(s)-1]
		s = s[:len(s)-1]

		if len(nums) < v.idx {
			continue
		}

		isExists := false
		for _, ints := range result {
			if reflect.DeepEqual(ints, v.path)	 {
				isExists = true
				break
			}
		}
		if !isExists {
			result = append(result, v.path)
		}

		for i := v.idx; i < len(nums); i++ {
			copied := make([]int, len(v.path))
			copy(copied, v.path)
			copied = append(copied, nums[i])

			s = append(s, StackVal37{i+1, copied})
		}
	}

	return result
}
