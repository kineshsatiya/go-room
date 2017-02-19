package main

import "fmt"

type Node struct {
    val interface{}
    next* Node
}

type Stack struct {
    head* Node
    size int
}

func (s *Stack) Push(value interface{}) {
    var n = &Node{value, nil}
    if s.head == nil {
        s.head = n
    } else {
        t := s.head
        n.next = t
        s.head = n
    }
    s.size++
}

func (s *Stack) Pop() (value interface{}) {
    if s.size > 0 {
        value = s.head.val
        s.head = s.head.next
        s.size--
        return
    }
    return
}

func main() {
    var s Stack = Stack{nil , 0}
    s.Push(2)
    s.Push(3)
    s.Push(4)
    fmt.Println(s)
    fmt.Println(s.Pop().(int))
    fmt.Println(s.Pop().(int))
    fmt.Println(s.Pop().(int))
    fmt.Println(s.Pop().(int))
}
