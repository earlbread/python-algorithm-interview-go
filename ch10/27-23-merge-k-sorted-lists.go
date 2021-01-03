package ch10

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "container/heap"

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

func mergeKLists(lists []*ListNode) *ListNode {
	q := &NodeHeap{}
	for _, node := range lists {
		if node != nil {
			heap.Push(q, node)
		}
	}

	dummy := &ListNode{}
	curr := dummy
	for q.Len() > 0 {
		node := heap.Pop(q).(*ListNode)

		if node.Next != nil {
			heap.Push(q, node.Next)
		}

		curr.Next = node
		curr = curr.Next
	}

	return dummy.Next
}
