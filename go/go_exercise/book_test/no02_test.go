package book

import "testing"

func TestNo02(t *testing.T) {
	cases := []struct {
	    input string
	    want string
	}{
	    {"hello", "olleh"},
		{"123456", "654321"},
	}

	for _, c := range cases {
	    rz := solve02n1(c.input)
	    if rz != c.want {
	        t.Fatalf("invalid result. rz:%s want:%s", rz, c.want)
	    }
	}
}

func solve02n1(input string) string {
	r := []rune(input)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func reversStr(input string) string {
	r := []rune(input)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
