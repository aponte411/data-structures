package main

import (
    "fmt"
)

// Singly linked list node
type Node struct {
    Val int
    Next *Node
}
// Singly linked list
type SinglyLinkedList struct {
    Head *Node
}

// AddToHead method
func (s *SinglyLinkedList) AddToHead(val int) {
    var node = &Node{}
    node.Val = val
    node.Next = nil

    if node.Head != nil {
        node.Next = s.Head
    }
    s.Head = node
}

// Iterate through nodes and print to stdout
func (s *SinglyLinkedList) Iterate() {
    var node = *Node
    for node = s.Head; node != nil; node = node.Next {
        fmt.Println(node.Val)
    }
}

// LastNode method
func (s *SinglyLinkedList) LastNode() *Node {
    var node *Node
    var last *Node
    for node = s.Head; node != nil; node = node.Next {
        if node == nil {
            last = node
        }
    }
    return last
}

// AddToEnd method
func (s *SinglyLinkedList) AddToEnd(val int) {
    var node = &Node{}
    node.Val = val
    node.Next = nil

    var last *Node
    last = s.LastNode()

    if last != nil {
        last.Next = node
    }
}

// NodeWithValue method
func (s *SinglyLinkedList) NodeWithValue(val int) *Node {
    var nodeWith = *Node
    var node = *Node
    for node = s.Head; node != nil; node = node.Next{
        if node.Val == val {
            nodeWith = node
            break
        }
    }
    return nodeWith
}
// AddAfter method
func (s *SinglyLinkedList) AddAfter(nodeVal, val int){
    var node = &Node
    node.Val = val
    node.Next = nil

    var nodeWith = *Node
    nodeWith = s.NodeWithValue(nodeVal)
    if nodeWith != nil {
        node.Next = nodeWith.Next
        nodeWith.Next = node
    }
}

// main
func main() {
    var list SinglyLinkedList
    list = SinglyLinkedList{}
    list.AddToHead(1)
    list.AddToHead(2)
    list.AddToEnd(4)
    list.AddAfter(1, 9)
    list.Iterate()
}
