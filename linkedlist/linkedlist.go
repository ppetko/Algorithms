package main

import "fmt"

type Node struct {
        key interface{}
        next *Node
}

type List struct {
        head *Node
}

func (L *List) add (key interface{}) {
        list := &Node{key:key}

        if L.head == nil {
                L.head = list
                return
        }

        current := L.head
        for current.next != nil {
                current = current.next
        }
        current.next = list
}

func (L *List) print () {
        cur := L.head
        for cur !=nil {
                fmt.Printf("%v ->",cur.key)
                cur = cur.next
        }
        fmt.Printf("\n")
}

func (L *List) search (key interface{}) {
        current := L.head
        for current.next != nil {
                current = current.next
                if current.key == key {
                        fmt.Printf("Found the key %v\n", key)
                }
        }
}
func (L *List) delete (key interface{}) {
        current := L.head
        prev := L.head
        for (current.next !=nil && current.key !=key) {
                prev = current
                current = current.next
        }

        if current == prev {
                prev.key = 0
                L.head = current.next
                return
        }
        prev.next = current.next
        current.key = nil
}
