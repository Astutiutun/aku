package main

import "fmt"

type Node struct {
    value int
    next  *Node
}

type LinkedList struct {
    head *Node
}

func (l *LinkedList) Add(value int) {
    newNode := &Node{value: value}
    if l.head == nil {
        l.head = newNode
        return
    }
    current := l.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}

func (l *LinkedList) AddToFront(value int) {
    newNode := &Node{value: value, next: l.head}
    l.head = newNode
}

func (l *LinkedList) Remove(value int) {
    if l.head == nil {
        return
    }
    if l.head.value == value {
        l.head = l.head.next
        return
    }
    current := l.head
    for current.next != nil && current.next.value != value {
        current = current.next
    }
    if current.next != nil {
        current.next = current.next.next
    }
}

func (l *LinkedList) Print() {
    current := l.head
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
    fmt.Println("nil")
}

func main() {
    list := LinkedList{}

    list.Add(1)
    list.Add(2)
    list.Add(3)
    list.Print() // Output: 1 -> 2 -> 3 -> nil
    list.AddToFront(0)
    list.Print() // Output: 0 -> 1 -> 2 -> 3 -> nil
    list.Remove(2)
    list.Print() // Output: 0 -> 1 -> 3 -> nil

    list.Remove(0)
    list.Print() // Output: 1 -> 3 -> nil

    list.Remove(3)
    list.Print() // Output: 1 -> nil
   }
