package ch14

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestDepth(node *TreeNode, longest *int) int {
	if node == nil {
		return 0
	}

	left := longestDepth(node.Left, longest)
	right := longestDepth(node.Right, longest)

	if node.Left != nil && node.Left.Val == node.Val {
		left++
	} else {
		left = 0
	}

	if node.Right != nil && node.Right.Val == node.Val {
		right++
	} else {
		right = 0
	}

	*longest = max(*longest, left+right)
	return max(left, right)
}

func longestUnivaluePath(root *TreeNode) int {
	longest := 0
	longestDepth(root, &longest)

	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
