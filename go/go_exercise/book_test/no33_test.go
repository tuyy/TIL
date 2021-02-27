package book

import (
	"reflect"
	"sort"
	"testing"
)

func TestNo33(t *testing.T) {
	input := "23"
	want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}

	rz := solve33n2(input)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve33n2(input string) []string {
	var result []string

	s := []StackVal{{0, ""}}

	for len(s) > 0 {
		v := s[len(s)-1]
		s = s[:len(s)-1]

		if len(input) == len(v.path) {
			result = append(result, v.path)
			continue
		}

		for _, c := range m[input[v.idx]] {
			s = append(s, StackVal{v.idx + 1, v.path + c})
		}
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

var m = map[uint8][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
}

type StackVal struct {
	idx  int
	path string
}

func solve33n1(input string) []string {
	s := []StackVal{{0, ""}}

	var result []string

	for len(s) > 0 {
		v := s[len(s)-1]
		s = s[:len(s)-1]

		if len(v.path) == len(input) {
			result = append(result, v.path)
			continue
		}

		for _, c := range m[input[v.idx]] {
			s = append(s, StackVal{v.idx + 1, v.path + c})
		}
	}
	sort.Strings(result)
	return result
}

func solve33(input string) []string {
	var result []string

	var dfs func(int, string)

	dfs = func(idx int, path string) {
		if len(path) == len(input) {
			result = append(result, path)
			return
		}

		for _, v := range m[input[idx]] {
			dfs(idx+1, path+v)
		}
	}

	dfs(0, "")

	return result
}
