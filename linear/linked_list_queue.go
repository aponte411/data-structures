package linear

import "fmt"

type QueueNode struct {
	Val  int
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

func (q *QueueLinkedList) Enqueue(val int) {
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

func (q *QueueLinkedList) Dequeue() int {
	node := q.Head
	if node == nil {
		fmt.Println("EmptyQueueException")
		return 0
	}
	q.Head = q.Head.Next
	q.size -= 1
	return node.Val
}
