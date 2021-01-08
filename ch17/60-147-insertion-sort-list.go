package ch17

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy

	for head != nil {
		for node.Next != nil && node.Next.Val < head.Val {
			node = node.Next
		}

		nextHead := head.Next
		node.Next, head.Next = head, node.Next

		head = nextHead

		if head != nil && head.Val < node.Next.Val {
			node = dummy
		}
	}

	return dummy.Next
}
