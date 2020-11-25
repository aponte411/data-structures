package main

import "fmt"

const Capacity = 100

type Queue struct {
	size  int
	data  [Capacity]interface{}
	front int
	back  int
}

func (q *Queue) Enqueue(val interface{}) {
	if q.size >= Capacity {
		return
	}
	q.size += 1
	q.data[q.back] = val
	q.back = (q.back + 1) % (Capacity - 1)
}

func (q *Queue) Dequeue() interface{} {
	var val interface{}
	if q.size <= 0 {
		return val
	}
	q.size -= 1
	val = q.data[q.front]
	q.front = (q.front + 1) % (Capacity - 1)
	return val
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
func (q *Queue) Length() int {
	return q.size
}

func main() {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Dequeue())
}
