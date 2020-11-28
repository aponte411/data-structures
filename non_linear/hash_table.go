package main

import "fmt"

type HashTable struct {
    array []*Node
    size int
}

type Node struct {
    key int
    value int
    next *Node
}
