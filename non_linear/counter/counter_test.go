package counter

import (
	"testing"
)

func TestCounter(t *testing.T) {
	counter := make(Counter)
	counter.Add("a")
	counter.Add("b")
	counter.Add("a")
	res1 := counter.Find("a")
	if !res1 {
		t.Errorf("Expected a to in counter, %v", res1)
	}

	res2, _ := counter.Get("a")
	if res2 != 2 {
		t.Errorf("Expected a, got %v", res2)
	}
}
