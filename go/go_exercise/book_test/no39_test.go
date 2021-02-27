package book

import (
	"fmt"
	"sort"
	"testing"
)

func TestNo39n2(t *testing.T) {
	k := 4
	input := [][]int{
		{1, 2},
		{1, 3},
		{3, 4},
		{4, 2},
	}
	want := true
	
	rz := solve39(k, input)
	if rz != want {
	    t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve39(k int, input [][]int) bool {
	result := true
	graph := makeGraphNo39(input)

	var dfs func(int, []int)

	dfs = func(n int, nums []int) {
		if nums != nil {
			isExists := false
			for _, v := range nums[:len(nums)-1] {
				if v == n {
					isExists = true
					break
				}
			}
			if isExists {
				result = false
				return
			}
		}

		nums = append(nums, n)

		for _, v := range graph[n] {
			copied := make([]int, len(nums))
			copy(copied, nums)
			copied = append(copied, v)

			dfs(v, copied)
		}
	}

	for i := 1; i <= k; i++ {
		dfs(i, nil)
		if !result {
			return false
		}
	}
	return true
}

func makeGraphNo39(input [][]int) map[int][]int {
	result := make(map[int][]int)
	for _, nums := range input {
		result[nums[0]] = append(result[nums[0]], nums[1])
	}
	for k := range result {
		sort.Ints(result[k])
	}
	return result
}

type Stack39 []int

func (s *Stack39) Pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack39) Push(v int) {
	*s = append(*s, v)
}

func TestNo39(t *testing.T) {
	//c := 3
	input := [][]int{{0, 1}, {1, 2}, {2, 0}}
	//input := [][]int{{0, 1}, {0, 2}, {1, 2}}

	m := make(map[int][]int)
	for _, v := range input {
		m[v[0]] = append(m[v[0]], v[1])
	}

	traced := make(map[int]struct{})

	var dfs func(int) bool

	dfs = func(v int) bool {
		if _, ok := traced[v]; ok {
			return false
		}

		traced[v] = struct{}{}
		for _, n := range m[v] {
			if !dfs(n) {
				return false
			}
		}

		delete(traced, v)

		return true
	}

	for k, _ := range m {
		if !dfs(k) {
			fmt.Println("순환구조다.")
			break
		}
	}
}

func TestNo39n1(t *testing.T) {
	//c := 3
	input := [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 1}}

	graph := make(map[int][]int)
	for _, v := range input {
		graph[v[0]] = append(graph[v[0]], v[1])
	}

	var dfs func(int) bool

	traced := make(map[int]struct{})

	dfs = func(v int) bool {
		if _, ok := traced[v]; ok {
			return false
		}

		traced[v] = struct{}{}

		for _, i := range graph[v] {
			if !dfs(i) {
				return false
			}
		}

		delete(traced, v)
		return true
	}

	fmt.Println(dfs(0))
}
