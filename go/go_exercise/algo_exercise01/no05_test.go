package book

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestNo05(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	want := [][]string {
		{"ate", "eat", "tea"},
		{"nat", "tan"},
		{"bat"},
	}

	rz := solve05(input)
	if !reflect.DeepEqual(rz, want) {
	    t.Fatalf("invalid result. rz:%s want:%s", rz, want)
	}
}

func solve05(input []string) [][]string {
	m := make(map[string][]string)
	for _, s := range input {
		tokens := strings.Split(s, "")
		sort.Strings(tokens)
		str := strings.Join(tokens, "")
		if arr, ok := m[str]; ok {
			m[str] = append(arr, s)
		} else {
			m[str] = []string{s}
		}
	}

	var result [][]string
	for _, v := range m {
		sort.Strings(v)
		result = append(result, v)
	}
	return result
}
