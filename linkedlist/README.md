# Data Structure - LinkedList Implementation

## Sample main 

```
  1 package main
  2 
  3 func main(){
  4         list := &List{}
  5         list.add(5)
  6         list.add(10)
  7         list.add(15)
  8         list.add(20)
  9         list.add(25)
 10         list.add(30)
 11         list.print()
 12         list.delete(5)
 13         list.delete(30)
 14         list.print()
 15         list.search(10)
 16 }
```
## Run the code

```
go  vet main.go linkedlist.go 
go build main.go linkedlist.go 
./main 
5 ->10 ->15 ->20 ->25 ->30 ->
10 ->15 ->20 ->25 ->

```

## Running time in Big(O)

```
			Access	Search	Inserrt Deletion Access	Search Insertion Deletion  Worst 
Singly-Linked List	Θ(n)	Θ(n)	Θ(1)	Θ(1)	  O(n)	O(n)	O(1)	   O(1)	   O(n)
```
