package book

import (
	"testing"
)

func TestNo06(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for _, c := range cases {
		rz := solve06(c.input)
		if rz != c.want {
			t.Fatalf("invalid result. rz:%s want:%s", rz, c.want)
		}
	}
}


func solve06(input string) string {
	var results []string
	for k := 2; k <= 3; k++ {
		for i := 0; i < len(input)-k; i++ {
			results = getPalindrome(input, i, k, results)
		}
	}

	max := ""
	for _, v := range results {
		if len(max) < len(v) {
			max = v
		}
	}
	return max
}

func getPalindrome(input string, i, k int, results []string) []string{
	if i < 0 || (i+k) > len(input) {
		return results
	}

	if isPalindrome(input[i : i+k]) {
		results = append(results, input[i:i+k])
		getPalindrome(input, i-1, k+2, results)
	}
	return results
}

func isPalindrome(input string) bool {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		if input[i] != input[j] {
			return false
		}
	}
	return true
}
