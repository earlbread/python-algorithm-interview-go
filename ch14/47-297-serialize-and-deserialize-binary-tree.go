package ch14

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	serialized := []string{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q[0], q = nil, q[1:]

		if node == nil {
			serialized = append(serialized, "null")
		} else {
			serialized = append(serialized, strconv.Itoa(node.Val))
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
	}

	return strings.Join(serialized, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	splited := strings.Split(data, ",")
	nodes := []*TreeNode{}
	for _, s := range splited {
		if s == "null" {
			nodes = append(nodes, nil)
		} else {
			value, _ := strconv.Atoi(s)
			nodes = append(nodes, &TreeNode{Val: value})
		}
	}

	root := nodes[0]
	q := []*TreeNode{root}
	nodes[0], nodes = nil, nodes[1:]

	for len(q) > 0 {
		node := q[0]
		q[0], q = nil, q[1:]

		if node == nil {
			continue
		}

		if len(nodes) > 0 {
			node.Left = nodes[0]
			nodes[0], nodes = nil, nodes[1:]
			q = append(q, node.Left)
		}

		if len(nodes) > 0 {
			node.Right = nodes[0]
			nodes[0], nodes = nil, nodes[1:]
			q = append(q, node.Right)
		}
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
