package ch17

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		prev = slow
		fast = fast.Next.Next
		slow = slow.Next
	}

	prev.Next = nil

	l1 := sortList(head)
	l2 := sortList(slow)

	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}

	root := l1
	node := root
	l1 = l1.Next

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}

		node = node.Next
	}

	if l1 == nil {
		node.Next = l2
	} else {
		node.Next = l1
	}

	return root
}
