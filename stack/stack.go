package main

import "fmt"  
 
type stack struct{
    myStack []string
}

func newStack() *stack {
    return &stack{make([]string,0)}
}

func (s *stack) Push(item string){
    s.myStack = append(s.myStack, item)
    fmt.Printf("Add %s\n", item)
}

func (s *stack) Len() *int{
    l := len(s.myStack) 
    return &l
}

func (s *stack) Pop(){
    l := len(s.myStack)
    fmt.Printf("Removing %v\n",s.myStack[0])
    s.myStack = s.myStack[:l-1]
}

func (s * stack) Print(){
    for _,j := range s.myStack {
        fmt.Printf("%s, ",j)
    }   
}
