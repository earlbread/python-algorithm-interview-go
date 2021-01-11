package ch14

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type NodeInfo struct {
	Diameter int
	Depth    int
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func diameter(node *TreeNode) NodeInfo {
	if node == nil {
		return NodeInfo{0, 0}
	}

	left := diameter(node.Left)
	right := diameter(node.Right)

	currentDiameter := left.Depth + right.Depth
	longest := currentDiameter

	if left.Diameter > longest {
		longest = left.Diameter
	}

	if right.Diameter > longest {
		longest = right.Diameter
	}

	return NodeInfo{longest, max(left.Depth, right.Depth) + 1}
}

func diameterOfBinaryTree(root *TreeNode) int {
	return diameter(root).Diameter
}
