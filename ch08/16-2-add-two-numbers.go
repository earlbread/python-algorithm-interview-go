package ch08

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	node := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		nextVal := 0

		if l1 != nil {
			nextVal += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			nextVal += l2.Val
			l2 = l2.Next
		}

		nextVal += carry
		nextVal, carry = nextVal%10, nextVal/10

		node.Next = &ListNode{Val: nextVal}
		node = node.Next
	}

	return dummy.Next
}
