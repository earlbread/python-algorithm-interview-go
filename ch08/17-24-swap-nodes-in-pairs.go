package ch08

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	node := dummy

	for node.Next != nil && node.Next.Next != nil {
		first := node.Next
		second := node.Next.Next

		first.Next = second.Next
		second.Next = first
		node.Next = second

		node = first
	}

	return dummy.Next
}
