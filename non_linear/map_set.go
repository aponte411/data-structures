package main

import "fmt"

type Set map[interface{}]bool

func (s *Set) Add(key interface{}) {
	(*s)[key] = true
}

func (s *Set) Remove(key interface{}) {
	delete((*s), key)
}

func (s *Set) Find(key interface{}) bool {
	return (*s)[key]
}

func main() {
	set := make(Set)
	set.Add(1)
	set.Add(2)
	fmt.Println(set.Find(1))
	set.Remove(1)
	fmt.Println(set.Find(1))
}
