package linear

import "fmt"

const Capacity = 100

type ArrayQueue struct {
	size  int
	data  [Capacity]interface{}
	front int
	back  int
}

func (a *ArrayQueue) Enqueue(val interface{}) {
	if a.size >= Capacity {
		return
	}
	a.size += 1
	a.data[a.back] = val
	a.back = (a.back + 1) % (Capacity - 1)
}

func (a *ArrayQueue) Dequeue() interface{} {
	if a.size == 0 {
		fmt.Println("EmptyQueueException")
		return nil
	}
	a.size -= 1
	val := a.data[a.front]
	a.front = (a.front + 1) % (Capacity - 1)
	return val
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.size == 0
}
func (a *ArrayQueue) Length() int {
	return a.size
}
