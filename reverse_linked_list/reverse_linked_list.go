package reverse_linked_list

import . "github.com/sirkay777/leetcode-challenge/leetcode/linkedlist"

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		current.Next, prev, current = prev, current, current.Next
	}
	return prev
}
