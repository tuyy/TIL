package book

import (
	"reflect"
	"sort"
	"testing"
)

func TestNo38(t *testing.T) {
	cases := []struct {
		input [][]string
		want  []string
	}{
		{
			[][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			[]string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			[][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			[]string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
	}

	for _, c := range cases {
		rz := solve38n2(c.input)
		if !reflect.DeepEqual(rz, c.want) {
			t.Fatalf("invalid result. rz:%v want:%v", rz, c.want)
		}
	}
}

func solve38n2(input [][]string) []string {
	graph := makeGraph(input)

	var traced []string

	var dfs func(string)

	dfs = func(air string) {
		traced = append(traced, air)

		for len(graph[air]) > 0{
			dest := graph[air][0]
			graph[air] = graph[air][1:]

			dfs(dest)
		}
	}

	dfs("JFK")
	return traced
}

func makeGraph(input [][]string) map[string][]string {
	graph := make(map[string][]string)
	for _, v := range input {
		graph[v[0]] = append(graph[v[0]], v[1])
	}
	for k := range graph {
		sort.Strings(graph[k])
	}
	return graph
}

type stack []string

func (s *stack) Pop() string {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func solve38n1(input [][]string) []string {
	m := make(map[string][]string)
	for _, v := range input {
		m[v[0]] = append(m[v[0]], v[1])
		sort.Strings(m[v[0]])
		for i, j := 0, len(m[v[0]])-1; i < j; i, j = i+1, j-1 {
			m[v[0]][i], m[v[0]][j] = m[v[0]][j], m[v[0]][i]
		}
	}

	var result []string
	s := stack{"JFK"}

	for len(s) > 0 {
		v := s.Pop()

		for len(m[v]) > 0 {
			dest := m[v][0]
			m[v] = m[v][1:]

			s.Push(dest)
		}

		result = append(result, v)
	}

	return result
}

func solve38(input [][]string) []string {
	var result []string

	graph := make(map[string][]string)
	for _, v := range input {
		graph[v[0]] = append(graph[v[0]], v[1])
		sort.Strings(graph[v[0]])
	}

	var dfs func(string)

	dfs = func(a string) {
		for len(graph[a]) > 0 {
			d := graph[a][0]
			graph[a] = graph[a][1:]
			dfs(d)
		}
		result = append(result, a)
	}

	dfs("JFK")
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
