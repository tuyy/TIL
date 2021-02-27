package book_test

import (
	"strings"
	"testing"
)

// 펠린드폼
func TestNo01(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal : Panama", true},
		{"race a car", false},
		{"123 4 321", true},
	}

	for _, c := range cases {
		rz := solve01n1(c.input)
		if rz != c.want {
			t.Fatalf("invalid result. rz:%v want:%v", rz, c.want)
		}
	}
}

func solve01n1(text string) bool {
	text = strings.ToLower(text)
	i, j := 0, len(text)-1
	for i < j {
		left := text[i]
		for ((left >= 'a' && left <= 'z') || (left >= '0' && left <= '9')) == false {
			i++
			left = text[i]
			if i > j {
				break
			}
		}

		right := text[j]
		for ((right >= 'a' && right <= 'z') || (right >= '0' && right <= '9')) == false {
			j--
			right = text[j]
			if i > j {
				break
			}
		}
		if left != right {
			return false
		}
		i, j = i+1, j-1
	}

	return true
}

func solve(input string) bool {
	var str string
	for _, c := range strings.ToLower(input) {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			str += string(c)
		}
	}

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true
}
