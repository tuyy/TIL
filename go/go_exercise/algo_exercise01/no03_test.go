package book

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

// 조건
// 1. 로그 가장 앞 부분은 식별자
// 2. 문자로 구성된 로그가 숫자 로그보다 앞에옴ㅐ
// 3. 식별자는 동일한경우 순서에 영향
// 4. 숫자 로그는 입력 순서대로 한다.
func TestNo03(t *testing.T) {
	cases := []struct {
		input []string
		want  []string
	}{
		{
			[]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit", "let3 art zero"},
			[]string{"let1 art can", "let3 art zero", "let2 own kit", "dig1 8 1 5 1", "dig2 3 6"},
		},
	}

	for _, c := range cases {
		rz := solve03n1(c.input)
		if !reflect.DeepEqual(rz, c.want) {
			t.Fatalf("invalid result. rz:%v want:%v", rz, c.want)
		}
	}
}

func solve03n1(input []string) []string {
	// []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit", "let3 art zero"},

	var strArr []string
	var numArr []string
	for _, s := range input {
		tokens := strings.Split(s, " ")
		if tokens[1][0] >= '0' && tokens[1][0] <= '9' {
			numArr = append(numArr, s)
		} else {
			strArr = append(strArr, s)
		}
	}

	sort.SliceStable(strArr, func(i, j int) bool {
		tokens1 := strings.Split(strArr[i], " ")[1:]
		tokens2 := strings.Split(strArr[j], " ")[1:]
		return strings.Join(tokens1, "") < strings.Join(tokens2, "")
	})

	return append(strArr, numArr...)
}

func solve(input []string) []string {
	var charArr []string
	var numArr []string
	for _, s := range input {
		word := strings.Split(s, " ")[1]
		if word[0] >= 'a' && word[0] <= 'z' {
			charArr = append(charArr, s)
		} else {
			numArr = append(numArr, s)
		}
	}

	for i := 1; i < len(charArr); i++ {
		key, val := getValue(charArr, i), charArr[i]
		j := i - 1
		for j >= 0 && getValue(charArr, j) > key {
			charArr[j+1] = charArr[j]
			j--
		}
		charArr[j+1] = val
	}

	for _, s := range numArr {
		charArr = append(charArr, s)
	}

	return charArr
}

func solve2(input []string) []string {
	var charArr []string
	var numArr []string
	for _, s := range input {
		word := strings.Split(s, " ")[1]
		if word[0] >= 'a' && word[0] <= 'z' {
			charArr = append(charArr, s)
		} else {
			numArr = append(numArr, s)
		}
	}

	sort.SliceStable(charArr, func(i, j int) bool {
		return getValue(charArr, i) < getValue(charArr, j)
	})

	charArr = append(charArr, numArr...)

	return charArr
}

func getValue(charArr []string, i int) string {
	return strings.Join(strings.Split(charArr[i], " ")[1:], "")
}
