package ch08

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	backward := prev
	forward := head
	for backward != nil && backward.Val == forward.Val {
		backward = backward.Next
		forward = forward.Next
	}

	return backward == nil
}

// Using slice
func isPalindrome2(head *ListNode) bool {
	l := make([]int, 0)

	node := head
	for node != nil {
		l = append(l, node.Val)
		node = node.Next
	}

	for left, right := 0, len(l)-1; left < right; left, right = left+1, right-1 {
		if l[left] != l[right] {
			return false
		}
	}

	return true
}
