package book

import (
	"reflect"
	"sort"
	"testing"
)

func TestNo36(t *testing.T) {
	cases := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		//{[]int{2, 3, 6, 7}, 7, [][]int{{7}, {2, 2, 3}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}

	for _, c := range cases {
		rz := solve36n3(c.candidates, c.target)
		if !reflect.DeepEqual(rz, c.want) {
			t.Fatalf("invalid result. rz:%v want:%v", rz, c.want)
		}
	}
}

func solve36n3(candidates []int, target int) [][]int {
	var result [][]int

	var dfs func(int, []int)

	dfs = func(total int, nums []int) {
		if total < 0 {
			return
		}
		if total == 0 {
			result = append(result, nums)
			return
		}

		for _, v := range candidates {
			if nums != nil && nums[len(nums)-1] > v {
				continue
			}

			copied := make([]int, len(nums))
			copy(copied, nums)
			copied = append(copied, v)

			dfs(total-v, copied)
		}
	}
	dfs(target, nil)

	return result
}

type Stack36n2 struct {
	sum, idx int
	path     []int
}

func solve36n2(candidates []int, target int) [][]int {
	var result [][]int

	s := []Stack36n2{{target, 0, nil}}

	for len(s) > 0 {
		v := s[len(s)-1]
		s = s[:len(s)-1]

		if v.sum < 0 {
			continue
		}
		if v.sum == 0 {
			result = append(result, v.path)
			continue
		}

		for i := v.idx; i < len(candidates); i++ {
			copied := make([]int, len(v.path))
			copy(copied, v.path)
			copied = append(copied, candidates[i])

			s = append(s, Stack36n2{v.sum - candidates[i], i, copied})
		}
	}

	return result
}

func solve36n1(candidates []int, target int) [][]int {
	var result [][]int
	var dfs func(int, int, []int)
	dfs = func(csum int, index int, path []int) {
		if csum < 0 {
			return
		}
		if csum == 0 {
			result = append(result, path)
			return
		}

		for i := index; i < len(candidates); i++ {
			copied := make([]int, len(path))
			copy(copied, path)
			copied = append(copied, candidates[i])

			dfs(csum-candidates[i], i, copied)
		}
	}
	dfs(target, 0, nil)
	return result
}

type StackVal36 struct {
	k      int
	result []int
}

func solve36(candidates []int, target int) [][]int {
	var result [][]int

	for i := 0; i < len(candidates); i++ {
		s := []StackVal36{{i, nil}}

		for len(s) > 0 {
			v := s[len(s)-1]
			s = s[:len(s)-1]

			if v.k > len(candidates) {
				sort.Ints(v.result)

				sum := 0
				for _, i := range v.result {
					sum += i
				}
				if sum == target {
					isExists := false
					for _, ints := range result {
						if reflect.DeepEqual(ints, v.result) {
							isExists = true
							break
						}
					}
					if !isExists {
						result = append(result, v.result)
					}
				}
				continue
			}

			for _, c := range candidates {
				copied := make([]int, len(v.result))
				copy(copied, v.result)
				copied = append(copied, c)

				s = append(s, StackVal36{v.k + 1, copied})
			}
		}
	}

	return result
}
