package linear

import "testing"

func TestArrayQueue(t *testing.T){

  	q := new(ArrayQueue)
   	q.Enqueue(1)
   	q.Enqueue(2)
   	q.Enqueue(3)
    val := q.Dequeue()
    if val != 1 {
        t.Errorf("Expected 1, got %v", )
    }
}
