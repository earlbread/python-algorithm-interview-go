package ch08

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
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
