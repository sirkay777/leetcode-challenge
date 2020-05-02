package reverselinkedlist

//ListNode singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		current.Next, prev, current = prev, current, current.Next
	}
	return prev
}
