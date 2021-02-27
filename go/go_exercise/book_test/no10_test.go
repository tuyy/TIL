package book

import (
	"sort"
	"testing"
)

func TestNo10(t *testing.T) {
	input := []int{1, 4, 3, 2}
	want := 4

	rz := solve10(input)
	if rz != want {
		t.Fatalf("invalid result. rz:%d want:%d", rz, want)
	}
}

func solve10(nums []int) int {
	sort.Ints(nums)

	min := 0
	for i := 0; i < len(nums); i += 2 {
		min += nums[i]
	}

	return min
}
