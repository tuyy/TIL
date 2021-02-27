package book

import (
	"math"
	"testing"
)

func TestNo12(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	want := 5

	rz := solve12n2(input)
	if rz != want {
		t.Fatalf("invalid result. rz:%d want:%d", rz, want)
	}
}

func solve12n2(input []int) int {
	min := math.MaxInt64
	profit := 0

	for _, v := range input {
		min = int(math.Min(float64(min), float64(v)))
		profit = int(math.Max(float64(profit), float64(v-min)))
	}
	return profit
}

func solve12n1(input []int) int {
	max := 0
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			diff := input[j] - input[i]
			if max < diff {
				max = diff
			}
		}
	}
	return max
}
