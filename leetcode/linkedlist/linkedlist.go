package linkedlist

import (
	"encoding/json"
	"errors"
	"fmt"
)

//ListNode  Singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

//NewLinkedList New a linked list from input string
func NewLinkedList(input string) (*ListNode, error) {
	var vals []int
	err := json.Unmarshal([]byte(input), &vals)
	if err != nil {
		panic(err)
		return nil, errors.New(fmt.Sprintf("input %s is not valid", input))
	}
	if len(vals) == 0 {
		return nil, nil
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
	if head == nil {
		return "[]"
	}
	var vals []int
	visited := make(map[*ListNode]bool)
	for head != nil {
		_, ok := visited[head]
		if ok {
			panic("cycle detected")
		}
		visited[head] = true
		vals = append(vals, head.Val)
		head = head.Next
	}
	res, err := json.Marshal(&vals)
	if err != nil {
		panic(err)
	}
	return string(res)
}
