package book

import "testing"

func TestNo30(t *testing.T) {
	cases := []struct {
	    input string
	    want int
	}{
		{"abcakdbb", 5},
		{"abcabcbb", 3},
		{"abcabcdbb", 4},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, c := range cases {
		rz := solve30(c.input)

		if rz != c.want {
			t.Fatalf("invalid result. rz:%d want:%d", rz, c.want)
		}
	}
}

func solve30(input string) int {
	max := 0
	for i := 0; i < len(input) - 1; i++ {
		m := make(map[uint8]bool)
		for j := i; j < len(input); j++ {
			if _, ok := m[input[j]]	; ok {
				if max < len(m) {
					max = len(m)
				}
				break
			} else {
				m[input[j]] = true
			}
		}
	}
	return max
}
