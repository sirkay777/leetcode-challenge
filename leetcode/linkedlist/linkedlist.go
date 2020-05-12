package linkedlist

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//ListNode  Singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

//NewLinkedList New a linked list from input string
func NewLinkedList(input string) (*ListNode, error) {
	filtered := strings.ReplaceAll(input, " ", "")
	if len(filtered) == 0 || string(filtered[0]) != "[" ||
		string(filtered[len(filtered)-1]) != "]" {
		return nil, errors.New(fmt.Sprintf("input %s is not valid", input))
	}

	if filtered == "[]" {
		return nil, nil
	}
	list := strings.Split(filtered[1:len(filtered)-1], ",")
	vals := []int{}
	for _, v := range list {
		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("input %s is not valid", input))
		}
		vals = append(vals, val)
	}
	head := &ListNode{Val: vals[0], Next: nil}
	tail := head
	for _, v := range vals[1:] {
		tail.Next = &ListNode{Val: v, Next: nil}
		tail = tail.Next
	}
	return head, nil
}

//LinkedListToString convert a linked list to its string representation
func LinkedListToString(head *ListNode) string {
	var vals []string
	visited := make(map[*ListNode]bool)
	for head != nil {
		_, ok := visited[head]
		if ok {
			panic("cycle detected")
		}
		visited[head] = true
		val := strconv.Itoa(head.Val)
		vals = append(vals, val)
		head = head.Next
	}
	return "[" + strings.Join(vals, ",") + "]"
}
