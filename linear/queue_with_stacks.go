package main

import "fmt"

type MyQueue struct {
	in_stack  Stack
	out_stack Stack
}

type Stack struct {
	data []int
}

func (s *Stack) push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) pop() int {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) peek() int {
	return s.data[len(s.data)-1]
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		in_stack:  Stack{data: []int{}},
		out_stack: Stack{data: []int{}},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in_stack.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.out_stack.isEmpty() {
		for !this.in_stack.isEmpty() {
			val := this.in_stack.pop()
			this.out_stack.push(val)
		}
	}
	return this.out_stack.pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.out_stack.isEmpty() {
		for !this.in_stack.isEmpty() {
			val := this.in_stack.pop()
			this.out_stack.push(val)
		}
	}
	return this.out_stack.peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.out_stack.isEmpty() && this.in_stack.isEmpty()
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	param_2 := queue.Pop()
	fmt.Println(param_2)
	param_3 := queue.Peek()
	fmt.Println(param_3)
	param_4 := queue.Empty()
	fmt.Println(param_4)
}
