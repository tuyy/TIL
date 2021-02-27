package book

import (
	"container/heap"
	"testing"
)

func TestNo40(t *testing.T) {
	// 다익스트라 알고리즘, 연습이 필요하다.
	times := [][]int{
		{2, 1, 1}, // 출발지, 도착지, 소요시간
		{2, 3, 1},
		{3, 4, 1},
	}
	n := 4 // 노드 수
	k := 2 // 시작점
	want := 2

	rz := solve40n1(times, n, k)
	if rz != want {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve40n1(times [][]int, n int, k int) int {
	graph := makeGraph40(times)
	q := []NodeAndWeight{{k, 0}}
	dist := make(map[int]int)

	for len(q) > 0 {
		n, w := q[0].Node, q[0].Weight
		q = q[1:]

		if _, ok := dist[n]; !ok {
			dist[n] = w
			for _, nodeAndWeight := range graph[n] {
				q = append(q, NodeAndWeight{nodeAndWeight.Node, nodeAndWeight.Weight+w})
			}
		}
	}

	return 0
}

func solve40(times [][]int, n, k int) int {
	graph := makeGraph40(times)
	que := &HeapQue{{k, 0}}
	heap.Init(que)
	dist := make(map[int]int)

	for que.Len() > 0 {
		dw := heap.Pop(que).(NodeAndWeight)
		if _, ok := dist[dw.Node]; !ok {
			dist[dw.Node] = dw.Weight
			for _, n := range graph[dw.Node] {
				alt := dw.Weight + n.Weight
				heap.Push(que, NodeAndWeight{n.Node, alt})
			}
		}
	}

	if len(dist) == n {
		max := 0
		for _, v := range dist {
			if max < v {
				max = v
			}
		}
		return max
	}
	return -1
}

type HeapQue []NodeAndWeight

func (h HeapQue) Len() int {
	return len(h)
}

func (h HeapQue) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h HeapQue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapQue) Push(x interface{}) {
	*h = append(*h, x.(NodeAndWeight))
}

func (h *HeapQue) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

type NodeAndWeight struct {
	Node, Weight int
}

func makeGraph40(times [][]int) map[int][]NodeAndWeight {
	result := make(map[int][]NodeAndWeight)

	for _, v := range times {
		result[v[0]] = append(result[v[0]], NodeAndWeight{v[1], v[2]})
	}

	return result
}
