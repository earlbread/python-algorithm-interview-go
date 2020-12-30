package ch08

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := head
	even_head := head.Next

	for node != nil && node.Next != nil && node.Next.Next != nil {
		odd := node
		even := node.Next

		odd.Next = even.Next
		even.Next = even.Next.Next
		node = odd.Next
	}

	node.Next = even_head

	return head
}
