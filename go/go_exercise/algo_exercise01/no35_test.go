package book

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

type No35Result [][]int

func (r No35Result) Len() int {
	return len(r)
}

func (r No35Result) Less(i, j int) bool {
	a := ""
	for _, v := range r[i] {
		a += strconv.Itoa(v)
	}
	b := ""
	for _, v := range r[j] {
		b += strconv.Itoa(v)
	}
	return a < b
}

func (r No35Result) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func TestNo35(t *testing.T) {
	want := [][]int{
		{2, 4},
		{3, 4},
		{2, 3},
		{1, 2},
		{1, 3},
		{1, 4},
	}

	rz := solve35()
	sort.Sort(No35Result(rz))

	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve35() [][]int {
	var result [][]int

	var dfs func(int, []int)

	dfs = func(idx int, path []int) {
		if k == idx {
			result = append(result, path)
			return
		}

		for i := idx+1; i <= n; i++ {
			if path != nil && path[len(path)-1] >= i {
				continue
			}

			copied := make([]int, len(path))
			copy(copied, path)
			copied = append(copied, i)

			dfs(idx+1, copied)
		}
	}

	dfs(0, nil)
	return result
}

var result [][]int

const (
	n = 4
	k = 2
)

func dfs33(idx int, values []int) {
	if idx == k {
		sort.Ints(values)
		isExists := false
		for _, ints := range result {
			if reflect.DeepEqual(ints, values) {
				isExists = true
				break
			}
		}
		if !isExists {
			result = append(result, values)
		}
		return
	}

	for v := 1; v <= n; v++ {
		isExists := false
		for _, i := range values {
			if i == v {
				isExists = true
			}
		}

		if !isExists {
			copied := make([]int, len(values))
			copy(copied, values)
			copied = append(copied, v)

			dfs33(idx+1, copied)
		}
	}
}

type StackVal44 struct {
	k      int
	result []int
}

func dfsWithStack() [][]int{
	s := []StackVal44{{0, nil}}

	var result [][]int

	for len(s) > 0 {
		v := s[len(s)-1]
		s = s[:len(s)-1]

		if v.k == k {
			sort.Ints(v.result)
			isExists := false
			for _, nums := range result {
				if reflect.DeepEqual(nums, v.result) {
					isExists = true
					break
				}
			}
			if !isExists {
				result = append(result, v.result)
			}
			continue
		}

		for i := 1; i <= n; i++ {
			isExists := false
			for _, j := range v.result {
				if j == i {
					isExists = true
					break
				}
			}

			if !isExists {
				copied := make([]int, len(v.result))
				copy(copied, v.result)
				copied = append(copied, i)

				s = append(s, StackVal44{v.k+1, copied})
			}
		}
	}
	return result
}
