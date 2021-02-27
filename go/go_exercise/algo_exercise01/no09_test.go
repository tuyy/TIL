package book

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestNo09(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	want := [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}

	rz := solve09n2(nums)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

//TODO 투 포인트 한번 더 보자
func solve09n2(nums []int) [][]int {
	var result [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0{
				left++
			} else {
				fmt.Printf("%d %d %d\n", nums[i], nums[left], nums[right])
				result = append(result, []int{nums[i], nums[left], nums[right]})
				break
			}
		}
	}

	return result
}

func solve09n1(nums []int) [][]int {
	sort.Ints(nums)

	m := make(map[string]bool)

	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if (nums[i] + nums[j] + nums[k]) == 0 {
					key := fmt.Sprintf("%d%d%d", nums[i], nums[j], nums[k])
					if _, ok := m[key]; ok {
						continue
					} else {
						m[key] = true
						result = append(result, []int{nums[i], nums[j], nums[k]})
					}
				}
			}
		}
	}
	return result
}
