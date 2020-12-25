package non_linear

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	counter := make(Counter)
	counter.Add("a")
	counter.Add("b")
	counter.Add("a")
	fmt.Println(counter.Find("a"))
	fmt.Println(counter.Get("a"))
}
