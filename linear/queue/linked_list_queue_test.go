package queue

import "testing"

func TestQueueLinkedList(t *testing.T) {
	queue := new(QueueLinkedList)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	res := queue.Dequeue()
	if res != 1 {
		t.Errorf("Expected 1, got %v", res)
	}
}
