package book

import "testing"

func TestNo41(t *testing.T) {
	n := 3
	edges := [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src := 0
	dst := 2
	k := 0
	want := 500

	rz := solve41n1(n, edges, src, dst, k)
	if rz != want {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve41n1(n int, edges [][]int, src int, dst int, K int) int {
	graph := makeGraph41(edges)
	que := [][3]int{{0, src, K}} // price, node, 경유지

	for len(que) > 0 {
		price, node, k := que[0][0],que[0][1], que[0][2]
		que = que[1:]

		if node == dst {
			return price
		}
		if k >= 0 {
			for _, nw := range graph[node] {
				alt := price + nw.Weight
				que = append(que, [3]int{alt, nw.Node, k-1})
			}
		}
	}

	return -1
}

func solve41(n int, edges [][]int, src, dst, k int) int {
	graph := makeGraph41(edges)
	dist := make(map[int][]int)
	que := []NodeAndWeight{{Node: src, Weight: 0}}

	for len(que) > 0 {
		node, weight := que[0].Node, que[0].Weight
		que = que[1:]

		dist[node] = append(dist[node], weight)

		for _, nw := range graph[node] {
			que = append(que, NodeAndWeight{Node: nw.Node, Weight: nw.Weight + weight})
		}
	}

	if v, ok := dist[dst]; ok {
		if len(v) > k {
			return v[k]
		}
		return -1
	} else {
		return -1
	}
}

func makeGraph41(edges [][]int) map[int][]NodeAndWeight {
	result := make(map[int][]NodeAndWeight)

	for _, v := range edges {
		result[v[0]] = append(result[v[0]], NodeAndWeight{Node: v[1], Weight: v[2]})
	}

	return result
}
