package main

import (
	"fmt"
	"github.com/golang-collections/collections/set"
)

func main() {
	set := set.New()
	set.Insert(1)
	set.Insert(2)
	fmt.Println(set.Has(1))
	set.Remove(1)
	fmt.Println(set.Has(1))
}
