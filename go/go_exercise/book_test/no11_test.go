package book

import (
	"reflect"
	"testing"
)

func TestNo11(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	want := []int{24, 12, 8, 6}

	rz := solve11n2(nums)
	if !reflect.DeepEqual(rz, want) {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve11n2(nums []int) []int {
	left := make([]int, len(nums))
	num := 1
	for i := 1; i < len(nums); i++ {
		num *= nums[i-1]
		left[i] = num
	}
	left[0] = 1

	result := make([]int, len(nums))
	num = 1
	for i := len(nums)-1; i >= 0 ; i-- {
		result[i] = left[i] * num
		num *= nums[i]
	}
	return result
}

func solve11n1(nums []int) []int {
	p := make([]int, len(nums))
	val := 1
	for i, num := range nums {
		val = num * val
		p[i] = val
	}

	result := make([]int, len(nums))
	preRight := 1
	for i := len(nums) - 1; i >= 0; i-- {

		left := 1
		if i > 0 {
			left = p[i-1]
		}

		right := 1
		if i < len(nums)-1 {
			right = preRight
		}
		result[i] = left * right
		preRight *= nums[i]
	}
	return result
}
