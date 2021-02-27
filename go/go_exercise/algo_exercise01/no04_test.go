package book

import (
	"math"
	"strings"
	"testing"
)

func TestNo04(t *testing.T) {
	input := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	want := "ball"

	rz := solve04n1(input, banned)
	if rz != want {
		t.Fatalf("invalid result. rz:%s want:%s", rz, want)
	}
}

func solve04n1(input string, banned []string) string {
	m := make(map[string]int)
	for _, s := range strings.Split(strings.ToLower(input), " ") {
		str := ""
		for _, c := range s {
			if c == ',' || c == '.' {
				continue
			}
			str += string(c)
		}

		if _, ok := m[str]; ok {
			m[str] += 1
		} else {
			m[str] = 1
		}
	}

	max := math.MinInt32
	result := ""
	for text, cnt := range m {
		isBanned := false
		for _, b := range banned {
			if b == text {
				isBanned = true
				break
			}
		}
		if isBanned {
			continue
		}

		if max < cnt {
			max = cnt
			result = text
		}
	}

	return result
}

func solve04(input string, banned []string) string {
	m := make(map[string]int)
	for _, s := range strings.Split(input, " ") {
		word := strings.ToLower(s)
		str := ""
		for _, c := range word {
			if c == '.' || c == ',' {
				continue
			}
			str += string(c)
		}

		isBanned := false
		for _, v := range banned {
			if v == str {
				isBanned = true
				break
			}
		}
		if isBanned {
			continue
		}

		if cnt, ok := m[str]; ok {
			m[str] = cnt + 1
		} else {
			m[str] = 1
		}
	}

	max := 0
	result := ""
	for key, val := range m {
		if max < val {
			result = key
			max = val
		}
	}

	return result
}
