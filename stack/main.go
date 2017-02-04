package main 

import "fmt"  

func main(){
    p := newStack()
    p.Push("California")
    p.Push("Chicago")
    p.Push("New York")
    fmt.Printf("My Stack: %s \n", p)
    fmt.Printf("My len is %d \n",*p.Len())
    p.Pop()
    fmt.Printf("My len is %d\n",*p.Len())
}
