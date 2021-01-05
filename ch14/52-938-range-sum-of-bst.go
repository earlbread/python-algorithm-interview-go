package ch14

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// DFS
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	result := 0

	if root.Val > high {
		result += rangeSumBST(root.Left, low, high)
	}

	if root.Val < low {
		result += rangeSumBST(root.Right, low, high)
	}

	if root.Val >= low && root.Val <= high {
		result += root.Val
		result += rangeSumBST(root.Left, low, high)
		result += rangeSumBST(root.Right, low, high)
	}

	return result
}

// BFS
func rangeSumBST2(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}
	result := 0

	for len(q) > 0 {
		node := q[0]

		if node != nil {
			if node.Val >= low && node.Val <= high {
				result += node.Val
				q = append(q, node.Left)
				q = append(q, node.Right)
			} else if node.Val < low {
				q = append(q, node.Right)
			} else if node.Val > high {
				q = append(q, node.Left)
			}
		}

		q = q[1:]
	}

	return result
}
