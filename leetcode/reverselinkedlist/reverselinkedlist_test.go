package reverselinkedlist

import (
	"fmt"
	"testing"
)

func BenchmarkReverseList(b *testing.B) {
	head := sliceToList([]int{1, 2, 3, 4, 5})
	for i := 0; i < b.N; i++ {
		ReverseList(head)
	}
}

func TestReverseLinkedList(t *testing.T) {
	data := []struct {
		input    []int
		expected string
	}{
		{[]int{1, 2, 3, 4, 5}, "[5,4,3,2,1]"},
		{[]int{1}, "[1]"},
		{[]int{}, "[]"},
	}

	for _, d := range data {
		head := sliceToList(d.input)
		reversedHead := ReverseList(head)
		got := printListToString(reversedHead)
		if got != d.expected {
			t.Errorf("got: %s, expected: %s", got, d.expected)
		}
	}
}

func sliceToList(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}
	head := &ListNode{Val: slice[0], Next: nil}
	tail := head
	for _, val := range slice[1:] {
		tail.Next = &ListNode{Val: val, Next: nil}
		tail = tail.Next
	}
	return head
}

func printListToString(head *ListNode) string {
	if head == nil {
		return "[]"
	}
	current := head
	s := "["
	for current != nil {
		s += fmt.Sprintf("%d,", current.Val)
		current = current.Next
	}
	s = s[:len(s)-1] + "]"
	return s
}
