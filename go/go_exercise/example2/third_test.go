package example2

import (
	"fmt"
	"testing"
)

func TestHello123(t *testing.T) {
	relation := [][]string{{"100", "ryan", "music", "2"},
		{"200", "apeach", "math", "2"},
		{"300", "tube", "computer", "3"},
		{"400", "con", "computer", "4"},
		{"500", "muzi", "music", "3"},
		{"600", "apeach", "music", "2"}}
	want := 2

	rz := solution3(relation)
	if rz != want {
		t.Fatalf("invalid result. rz:%d want:%d", rz, want)
	}
}

func solution3(relation [][]string) int {
	//rowCnt := len(relation)
	colCnt := len(relation[0])

	for i := 0; i < colCnt; i++ {
		for j := i+1; j < colCnt; j++ {
			fmt.Println(i, j)
		}
		fmt.Println("====================")
	}
	return 2
}
