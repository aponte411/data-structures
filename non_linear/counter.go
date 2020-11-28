package main

import "fmt"

type Counter map[interface{}]int

func (c *Counter) Add(key interface{}) {
	(*c)[key] += 1
}

func (c *Counter) Find(key interface{}) bool {
	_, ok := (*c)[key]
	return ok
}

func (c *Counter) Get(key interface{}) (int, bool) {
	val, ok := (*c)[key]
	return val, ok
}

func main() {
	counter := make(Counter)
	counter.Add("a")
	counter.Add("b")
	counter.Add("a")
	fmt.Println(counter.Find("a"))
	fmt.Println(counter.Get("a"))
}
