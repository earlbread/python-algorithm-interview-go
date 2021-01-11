package ch14

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	prev := math.MinInt32
	result := math.MaxInt32

	stack := []*TreeNode{}
	node := root

	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node, stack = stack[len(stack)-1], stack[:len(stack)-1]

		result = min(result, node.Val-prev)
		prev = node.Val
		node = node.Right
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
