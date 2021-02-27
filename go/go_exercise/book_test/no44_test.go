package book

import (
	"math"
	"testing"
)

func TestNo44(t *testing.T) {
	root := makeTreeNode()
	want := 2

	rz := solve44(root)
	if rz != want {
	    t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve44(root *TreeNode) int {
	var dfs func(*TreeNode) int

	var result int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if node.Left != nil && node.Value == node.Left.Value {
			left += 1
		} else {
			left = 0
		}
		if node.Right != nil && node.Value == node.Right.Value {
			right += 1
		} else {
			right = 0
		}

		result = int(math.Max(float64(result), float64(left + right)))

		return int(math.Max(float64(left), float64(right)))
	}
	dfs(root)

	return result
}

func copyAndAppend(v int, path []int) []int {
	copied := make([]int, len(path))
	copy(copied, path)
	copied = append(copied, v)
	return copied
}
