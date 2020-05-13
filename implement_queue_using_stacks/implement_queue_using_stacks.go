package implement_queue_using_stacks

import . "github.com/sirkay777/leetcode-challenge/leetcode/stack"

type MyQueue struct {
	first  Stack
	second Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.first.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for !this.first.Empty() {
		this.second.Push(this.first.Pop())
	}
	result := this.second.Pop()
	for !this.second.Empty() {
		this.first.Push(this.second.Pop())
	}
	return result
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	for !this.first.Empty() {
		this.second.Push(this.first.Pop())
	}
	result := this.second.Peek()
	for !this.second.Empty() {
		this.first.Push(this.second.Pop())
	}
	return result
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.first.Empty()
}
