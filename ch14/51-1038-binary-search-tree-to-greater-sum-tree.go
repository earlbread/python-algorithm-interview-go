package ch14

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Iteration
func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	stack := []*TreeNode{}
	node := root
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Right
		}

		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		node.Val += sum
		sum = node.Val
		node = node.Left
	}

	return root
}

// Recursion
func bstToGst2(root *TreeNode) *TreeNode {
	sum := 0
	helper(root, &sum)
	return root
}

func helper(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	helper(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	helper(root.Left, sum)
}
