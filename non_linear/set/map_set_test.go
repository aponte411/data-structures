package map_set

import (
	"testing"
)

func TestMapSet(t *testing.T) {
	set := make(MapSet)
	set.Add(1)
	set.Add(2)
	res := set.Find(1)
	if !res {
		t.Errorf("Missing 1 from set")
	}
	set.Remove(1)
	res2 := set.Find(1)
	if res2 {
		t.Errorf("1 should have been deleted")
	}
}
