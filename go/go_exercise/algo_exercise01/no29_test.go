package book

import (
	"testing"
)

func TestNo29(t *testing.T) {
	j := "aA"
	s := "aAAbbbb"
	want := 3

	rz := solve29n1(j, s)
	if rz != want {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve29n1(j string, s string) int {
	m := make(map[int32]int)
	for _, c := range s {
		m[c] = m[c] + 1
	}

	result := 0
	for _, c := range j {
		if v, ok := m[c]; ok {
			result += v
		}
	}
	return result
}

func solve29(j string, s string) int {
	m := make(map[int32]int)
	for _, c := range s {
		for _, c2 := range j {
			if c == c2 {
				if no, ok := m[c]; ok {
					m[c] = no + 1
				} else {
					m[c] = 1
				}
			}
		}
	}

	result := 0
	for _, v := range m {
		result = result + v
	}

	return result
}
