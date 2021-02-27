package example2

import (
	"reflect"
	"testing"
)

func TestHappy333334(t *testing.T) {
	n := 4
	stages := []int{4,4,4,4,4}
	want := []int{4,1,2,3}

	rz := solution2(n, stages)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

type StageAndFailure struct {
	Stage   int
	Failure float64
}

func solution2(n int, stages []int) []int {
	arr := make([]StageAndFailure, n)
	var nextStages []int
	for i := 1; i <= n; i++ {
		cnt := 0
		for _, s := range stages {
			if s == i {
				cnt++
			} else {
				nextStages = append(nextStages, s)
			}
		}
		arr[i-1] = StageAndFailure{i, float64(cnt) / float64(len(stages))}
		stages = nextStages
		nextStages = nil
	}

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j].Failure < key.Failure  {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	result := make([]int, n)
	for i, v := range arr {
		result[i] = v.Stage
	}

	return result
}

func TestInsertFunc(t *testing.T) {
	input := []int{3, 4, 5, 6, 1, 1, 2}
	want := []int{1, 1, 2, 3, 4, 5, 6}

	sortInsert(input)
	if !reflect.DeepEqual(input, want) {
	    t.Fatalf("invalid result. rz:%v want:%v", input, want)
	}
}

func sortInsert(arr []int) {
	for i, v := range arr {
		key := v
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j + 1] = arr[j]
			j--
		}
		arr[j + 1] = key
	}
}
