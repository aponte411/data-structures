package queue

import "testing"

func TestQueueWithStacks(t *testing.T) {
	queue := InitQueueWithStacks()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	param_2 := queue.Pop()
	if param_2 != 1 {
		t.Errorf("Expected 1, got %v", param_2)
	}

	param_3 := queue.Peek()
	if param_3 != 2 {
		t.Errorf("Expected 2, got %v", param_3)
	}

}
