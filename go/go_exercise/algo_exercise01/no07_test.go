package book

import (
	"reflect"
	"testing"
)

func TestNo07(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	want := []int{0, 1}

	rz := solve07n2(nums, target)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve07n2(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if v, ok := m[target-num]; ok && i != v {
			return []int{v, i}
		}
		m[num] = i
	}

	return nil
}

func solve07n1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
