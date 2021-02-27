package dfs_ex

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var graph = map[int][]int{
	1: {2, 3, 4},
	2: {5},
	3: {5},
	4: {},
	5: {6, 7},
	6: {},
	7: {3},
}

func dfs33() {
	stack := []int{1}

	var visited []int

	for len(stack) > 0{
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		visited = append(visited, i)
		fmt.Println(i)
		for _, v := range graph[i] {
			if !isVisited(visited, v) {
				stack = append(stack, v)
			}
		}
	}
}

func isVisited(visited []int, k int) bool {
	for _, v := range visited {
		if v == k {
			return true
		}
	}
	return false
}

func TestDfs1(t *testing.T) {
	dfs33()
	//want := []int{1, 2, 5, 6, 7, 3, 4}

	//rz := dfs1(1, nil)
	//if !reflect.DeepEqual(rz, want) {
	//	t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	//}
}

func dfs1(v int, discovered []int) []int {
	discovered = append(discovered, v)
	for _, w := range graph[v] {
		if !isExists(discovered, w) {
			discovered = dfs1(w, discovered)
		}
	}
	return discovered
}

func isExists(discovered []int, w int) bool {
	for _, v := range discovered {
		if v == w {
			return true
		}
	}
	return false
}

func TestSolution32(t *testing.T) {
	input := [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
	}
	want := 2

	rz := solution32(input)
	if rz != want {
		t.Fatalf("invalid result. rz:%d want:%d", rz, want)
	}
}

func solution32(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				dfs2(i, j, grid)
				count++
			}
		}
	}
	return count
}

func dfs2(i int, j int, grid [][]int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0

	dfs2(i+1, j, grid)
	dfs2(i, j+1, grid)
	dfs2(i-1, j, grid)
	dfs2(i, j-1, grid)
}

var arr []int = []int{1, 2, 3, 4}

var result [][]int

func TestDfs3(t *testing.T) {
	n := 4
	k := 2
	want := [][]int{
		{2, 4},
		{3, 4},
		{2, 3},
		{1, 2},
		{1, 3},
		{1, 4},
	}

	rz := solution35(n, k)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

var results [][]int
var n int = 4

func solution35(n int, k int) [][]int {
	dfs(nil, 0, 2)
	return results
}

func dfs(elem []int, start, k int) {
	if k == 0 {
		var result []int
		// TODO copy의 비밀부터 파헤쳐라
		copy(result, elem)
		fmt.Println(elem)
		results = append(results, result)
		return
	}

	for i := start; i < n+1; i++ {
		elem = append(elem, i)
		dfs(elem, i+1, k-1)
		elem = elem[:len(elem)-1]
	}
}

type Stack struct {
	arr []int
}

func (s *Stack) Pop() (int, error) {
	if len(s.arr) > 0 {
		fmt.Println(len(s.arr))
		v := s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return v, nil
	}
	return 0, errors.New("empty stack size")
}

func NewStack() *Stack {
	s := &Stack{}
	return s

}

func (s *Stack) Push(v int) {
	s.arr = append(s.arr, v)
}

var myresults [][]int

func TestMap(t *testing.T) {
	m := make(map[string]string)
	if _, ok := m["hello"]; ok {

	}
}

func TestSlide(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	a2 := arr[:len(arr)-1]
	fmt.Printf("arr:%v\n", arr)
	fmt.Printf("a2:%v\n", a2)

	tmp := make([]int, len(arr))
	fmt.Printf("copy:%d\n", copy(tmp, arr))
	myresults = append(myresults, tmp)
	myresults[0][0] = 10
	fmt.Printf("myresults:%v\n", myresults)

	fmt.Printf("arr:%v\n", arr)
	fmt.Printf("a2:%v\n", a2)
}

func TestDfs4(t *testing.T) {
	arr := []string{"a", "b", "c", "d"}
	k := 3
	want := []string{"abc", "abd", "acd", "bcd"}

	s := NewSolution(arr, k)
	rz := s.solve()
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

type Solution struct {
	strings  []string
	k        int
	results  [][]string
	visitMap map[string]bool
}

func (s *Solution) solve() []string {
	s.dfs(nil, 0, s.k)
	result := make([]string, len(s.results))
	for i, v := range s.results {
		result[i] = strings.Join(v, "")
	}
	fmt.Println(result)
	return result
}

func (s *Solution) dfs(elems []string, start, k int) {
	if k == 0 {
		str := strings.Join(elems, "")
		if _, ok := s.visitMap[str]; ok {
			return
		} else {
			s.visitMap[str] = true
			tmp := make([]string, len(elems))
			copy(tmp, elems)
			s.results = append(s.results, tmp)
		}
	}

	fmt.Printf("m:%v\n", s.visitMap)

	for i := start; i < len(s.strings); i++ {
		elems = append(elems, s.strings[i])
		s.dfs(elems, i+1, k-1)
		elems = elems[:len(elems)-1]
	}
}

func NewSolution(strings []string, k int) *Solution {
	m := make(map[string]bool)
	s := Solution{strings, k, nil, m}
	return &s
}

func TestPalindrome(t *testing.T) {
	cases := []struct {
	    input string
	    want bool
	}{
	    {"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for _, c := range cases {
	    rz := solution(c.input)
	    if rz != c.want {
	        t.Fatalf("invalid result. rz:%v want:%v", rz, c.want)
	    }
	}
}

func solution(input string) bool {
	str := ""
	input = strings.ToLower(input)
	for _, c := range input {
		if c >= 'a' && c <= 'z' {
			str += string(c)
		}
	}

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			fmt.Printf("i:%d j:%d\n", i, j)
			return false
		}
	}

	return true
}


