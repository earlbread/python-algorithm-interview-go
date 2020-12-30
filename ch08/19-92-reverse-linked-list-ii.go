package ch08

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{0, head}
	start := dummy
	for i := 0; i < m-1; i++ {
		start = start.Next
	}

	end := start.Next
	for i := 0; i < n-m; i++ {
		startNext := start.Next

		start.Next = end.Next
		end.Next = end.Next.Next
		start.Next.Next = startNext
	}

	return dummy.Next
}
