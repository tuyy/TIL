package book

import (
	"math"
	"testing"
)

// TODO 다시 풀어봐야한다.
func TestNo08(t *testing.T) {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	want := 8
	//input := []int{0, 0, 0, 0, 2}
	//want := 0

	rz := solve08(input)
	if rz != want {
	    t.Fatalf("invalid result. rz:%d want:%d", rz, want)
	}
}

func solve08(input []int) int {
	total := 0
	left, right := 0, len(input)-1
	leftMax, rightMax := input[left], input[right]

	for left < right {
		leftMax = int(math.Max(float64(leftMax), float64(input[left])))
		rightMax = int(math.Max(float64(rightMax), float64(input[right])))

		if leftMax <= rightMax {
			total += leftMax - input[left]
			left++
		} else {
			total += rightMax - input[right]
			right--
		}
	}

	return total
}

