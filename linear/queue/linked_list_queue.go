package queue

import "fmt"

type QueueNode struct {
	Val  interface{}
	Next *QueueNode
}

type QueueLinkedList struct {
	Head *QueueNode
	Tail *QueueNode
	size int
}

func (q *QueueLinkedList) Size() int {
	return q.size
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueLinkedList) Enqueue(val interface{}) {
	node := &QueueNode{val, nil}
	if q.Head == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
	q.size += 1
}

func (q *QueueLinkedList) Dequeue() interface{} {
	node := q.Head
	if node == nil {
		fmt.Println("EmptyQueueException")
		return nil
	}
	q.Head = q.Head.Next
	q.size -= 1
	return node.Val
}

func (q *QueueLinkedList) Peek() interface{} {
	if q.Head == nil {
		fmt.Println("EmptyQueueException")
		return nil
	}
	return q.Head.Val
}
