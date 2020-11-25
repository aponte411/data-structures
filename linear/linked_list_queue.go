package main

type Node struct {
	Val  int
	Next *Node
}

type QueueLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

func (q *QueueLinkedList) Size() int {
	return q.size
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueLinkedList) Enqueue(val int) {
	node := &Node{val, nil}
	if q.Head == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
	q.size += 1
}
