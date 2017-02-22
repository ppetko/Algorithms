# Data Structure - Merge Sorting Implementation

## Sample main 

```
  1 package main
  2 
  3 func main() {
  4         s := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
  5         MergeSort(s)
  6 }
```
## Run the code

```
go vet main.go mergesort.go 
go build main.go mergesort.go 
./main 
[4]
[1]
[1]
[1 3]
[1]
[1 3]
[1 3 4]
[1 3 4 6]
[2]
[7]
[5]
[2]
[2 5]
[2 5 7]
[2 5 7 8]
[1]
[1 2]
[1 2 3]
[1 2 3 4]
[1 2 3 4 5]
[1 2 3 4 5 6]
[1 2 3 4 5 6 7]
[1 2 3 4 5 6 7 8]
[1 2 3 4 5 6 7 8 9]
```

## Running time in Big(O)

```
			Best 		Average		Wost		
Mergesort		Ω(n log(n))	Θ(n log(n))	O(n log(n))
```
