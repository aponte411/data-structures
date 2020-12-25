package non_linear

import (
	"fmt"
	"testing"
)

func TestMapSet(t *testing.T) {
	set := make(Set)
	set.Add(1)
	set.Add(2)
	fmt.Println(set.Find(1))
	set.Remove(1)
	fmt.Println(set.Find(1))

}
