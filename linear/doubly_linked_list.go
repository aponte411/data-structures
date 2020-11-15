package main

import "fmt"

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Head *Node
}
