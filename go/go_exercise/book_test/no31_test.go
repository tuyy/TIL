package book

import (
	"reflect"
	"testing"
)

func TestNo31(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	want := []int{1, 2}

	rz := solve31(nums, k)
	if !reflect.DeepEqual(rz, want) {
	    t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve31(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] = m[num] + 1
	}

	var result []int
	for key, val := range m {
		if val >= k {
			result = append(result, key)
		}
	}

	return result
}
