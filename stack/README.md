# Data Structure - Stack Implementation using slices

## Sample main 

```
  1 package main
  2 
  3 import "fmt"
  4 
  5 func main(){
  6     p := newStack()
  7     p.Push("California")
  8     p.Push("Chicago")
  9     p.Push("New York")
 10     fmt.Printf("My Stack: %s \n", p)
 11     fmt.Printf("My len is %d \n",*p.Len())
 12     p.Pop()
 13     fmt.Printf("My len is %d\n",*p.Len())
 14 }
```
## Run the code
```
go vet main.go stack.go 
go build main.go stack.go 
./main 
Add California
Add Chicago
Add New York
My Stack: &{[California Chicago New York]} 
My len is 3 
Removing California
My len is 2
```

# Running time in BigO
```
        Access  Search  Insertion Deletion  Access    Search   Insert   Deletion Worst
Stack   Θ(n)    Θ(n)      Θ(1)    Θ(1)      O(n)        O(n)    O(1)    O(1)    O(n)

        O(fmtn) < O(n) < O(nlogn) < O(n^2) < O(2^n) < O(n!)
```
