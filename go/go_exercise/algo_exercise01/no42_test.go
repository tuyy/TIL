package book

import (
	"fmt"
	"math"
	"testing"
)

func TestNo42(t *testing.T) {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	want := 3

	rz := solve42n1(nums)
	if rz != want {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve42n1(trees []int) int {
	depth := 0
	q := []int{0}

	for len(q) > 0 {
		pre := len(q)
		depth++

		for i := 0; i < pre; i++ {
			i := q[0]
			q = q[1:]

			left := i*2+1
			if len(trees) > left && trees[left] != -1 {
				q = append(q, left)
			}

			right := i*2+2
			if len(trees) > right && trees[right] != -1 {
				q = append(q, right)
			}
		}
	}
	fmt.Println(depth)
	return depth
}

func solve42(nums []int) int {
	idx := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != -1 {
			idx = i
			break
		}
	}

	x, i := 0, 0
	for ; idx >= x; i++ {
		x = x + int(math.Pow(2.0, float64(i)))
	}
	return i
}
