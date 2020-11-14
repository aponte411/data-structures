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
    for node.Next != nil {
        node.Next = s.Head
    }
    s.Head = node
}

func (s *SinglyLinkedList) Iterate() {
    var node = *Node
    for node = s.Head; node != nil; node = node.Next {
        fmt.Println(node.Val)
    }
}

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

func (s *SinglyLinkedList) AddToEnd(val int) {
    var node &Node{}
    node.Val = val
    node.Next = nil
    var last *Node
    last = s.LastNode()
    if last != nil {
        last.Next = node
    }
}

func (s *SinglyLinkedList) NodeWithValue(val int) *Node {
    var nodeWith = *Node
    for node := s.Head; node != nil; node = node.Next{
        if node.Val == val {
            nodeWith = node
            break
        }
    }
    return nodeWith
}

func (s *SinglyLinkedList) AddAfter(nodeVal, val int){
    var node = &Node
    node.Val = nodeVal
    nodeWith := s.NodeWithValue(val)
    if nodeWith != nil {
        node.Next = nodeWith.Next
        nodeWith.Next = node
    }
}


func main() {
    list := SinglyLinkedList{}
    list.AddToHead(1)
    list.AddToHead(2)
    list.AddToEnd(5)
    list.AddAfter(1, 7)
    list.Iterate()
}
