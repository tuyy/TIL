package book

import (
	"math"
	"testing"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func TestNo43(t *testing.T) {
	root := makeTreeNode()
	want := 3

	rz := solve43(root)
	if rz != want {
		t.Fatalf("invalid result. rz:%v want:%v", rz, want)
	}
}

func solve43(root *TreeNode) int {
	var dfs func(*TreeNode) int

	result := -1

	dfs = func(node *TreeNode) int{
		if node == nil {
			return -1
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		result = int(math.Max(float64(result), float64(left + right + 2)))

		return int(math.Max(float64(left), float64(right))) + 1
	}

	dfs(root)
	return result
}

func makeTreeNode() *TreeNode {
	n4 := &TreeNode{Value: 4}
	n5 := &TreeNode{Value: 5}
	n2 := &TreeNode{Value: 2, Left: n4, Right: n5}
	n3 := &TreeNode{Value: 3}
	return &TreeNode{Value: 1, Left: n2, Right: n3}
}
