package ch14

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return depth(root) != -1
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := depth(node.Left)
	right := depth(node.Right)

	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
