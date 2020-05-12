package reverse_nodes_in_k_group

import (
	. "github.com/sirkay777/leetcode-challenge/leetcode/linkedlist"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil || head.Next == nil {
		return head
	}
	count := 1
	current := head
	var prev *ListNode
	for ; current != nil && count <= k; count++ {
		current, prev = current.Next, current
	}
	if count <= k {
		return head
	} else {
		prev.Next = nil
		newHead := reverseLinkedList(head)
		head.Next = reverseKGroup(current, k)
		return newHead
	}
}

func reverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		current, current.Next, prev = current.Next, prev, current
	}
	return prev
}
