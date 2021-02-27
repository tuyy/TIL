package kakao2018

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestHappy8(t *testing.T) {
	cases := []struct{
		str1, str2 string
		want int
	}{
		{"FRANCE",	"french",	16384},
		{"handshake",	"shake hands",	65536},
		{"aa1+aa2",	"AAAA12",	43690},
		{"E=M*C^2",	"e=m*c^2",	65536},
	}

	for _, c := range cases {
		rz := calcResult(c.str1, c.str2)
		if rz != c.want {
			t.Fatalf("invalid result. want:%d rz:%d(%s)", c.want, rz, c.str1)
		}
	}
}

func calcResult(str1 string, str2 string) int {
	arr1 := getSorted(strings.ToUpper(str1))
	arr2 := getSorted(strings.ToUpper(str2))

	if len(arr1) == 0 && len(arr2) == 0 {
		return 65536
	}

	interCnt := 0
	for _, v1 := range arr1 {
		for i, v2 := range arr2 {
			if v1 == v2 {
				interCnt++
				arr2[i] = ""
				break
			}
		}
	}

	unionCnt := len(arr1)

	for _, v := range arr2 {
		if v != "" {
			unionCnt += 1
		}
	}

	return int((float64(interCnt) / float64(unionCnt)) * float64(65536))
}

func getSorted(str string) []string{
	var rz []string
	for i := 1; i < len(str); i++ {
		value := fmt.Sprintf("%c%c", str[i-1], str[i])
		if value[0] < 'A' || value[0] > 'Z' {
			continue
		}
		if value[1] < 'A' || value[1] > 'Z' {
			continue
		}
		rz = append(rz, value)
	}
	sort.Strings(rz)
	return rz
}
