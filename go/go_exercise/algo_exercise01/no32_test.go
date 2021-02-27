package book

import "testing"

func TestNo32(t *testing.T) {
	input := [][]int{
		{1, 1, 0, 0, 1},
		{1, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{1, 0, 0, 1, 1},
	}
	want := 5

	rz := solve32(input)
	if rz != want {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve32(m [][]int) int {
	result := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 1 {
				mark(i, j, m)
				result++
			}
		}
	}
	return result
}

func mark(i int, j int, m [][]int) {
	if m[i][j] == 0 {
		return
	}

	m[i][j] = 0
	if i+1 < len(m) {
		mark(i+1, j, m)
	}
	if i-1 >= 0 {
		mark(i-1, j, m)
	}
	if j+1 < len(m[i]) {
		mark(i, j+1, m)
	}
	if j-1 >= 0 {
		mark(i, j-1, m)
	}
}
