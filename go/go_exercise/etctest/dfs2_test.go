package etctest

import (
	"fmt"
	"testing"
)

var graph2 = map[int][]int{
	1: {2, 3, 4},
	2: {5},
	3: {5},
	4: {},
	5: {6, 7},
	6: {},
	7: {3},
}

func TestDfs22(t *testing.T) {
	var discovered []int
	dfs222(1, &discovered)

	for _, v := range discovered {
		fmt.Printf("%d ", v)
	}

	fmt.Println()
	dfsWithStack()
}

func TestBfs22(t *testing.T) {
	discovered := bfs22(1)
	fmt.Println(discovered)
}

func bfs22(start int) []int {
	discovered := []int{start}
	que := []int{start}

	for len(que) > 0 {
		v := que[0]
		que = que[1:]

		for _, n := range graph2[v] {
			isExists := false
			for _, d := range discovered {
				if d == n {
					isExists = true
					break
				}
			}

			if !isExists {
				discovered = append(discovered, n)
				que = append(que, n)
			}
		}
	}
	return discovered
}
func dfs222(v int, discovered *[]int) {
	*discovered = append(*discovered, v)

	for _, n := range graph2[v] {
		isExists := false
		for _, d := range *discovered {
			if n == d {
				isExists = true
				break
			}
		}

		if !isExists {
			dfs222(n, discovered)
		}
	}
}

func dfsWithStack() {
	var stack []int
	stack = append(stack, 1)

	var discovered []int

	for len(stack) > 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		discovered = append(discovered, v)

		for _, n := range graph2[v] {
			isExists := false
			for _, d := range discovered {
				if n == d {
					isExists = true
					break
				}
			}

			if !isExists {
				stack = append(stack, n)
			}
		}
	}
	fmt.Println(discovered)
}
