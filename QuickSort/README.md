# Data Structure - Quick Sort Implementation

## Sample main 

```
1 package main
2 
3 func main() {
4         numbers := []int{5, 1, 3, 6, 4, 7, 2, 8, 9}
5         fmt.Println(numbers)
6         quickSort(numbers, 0, len(numbers) - 1)
7         fmt.Println(numbers)
8 }
```
## Run the code

```
go vet main.go quicksort.go 
go build main.go quicksort.go 
./main 
[5 1 3 6 4 7 2 8 9]
[1 2 3 4 5 6 7 8 9]
```

## Running time in Big(O)

```
		Best		Average		Worst	
Quicksort	Ω(n log(n))	Θ(n log(n))	O(n^2)
```
