package main

import "fmt"

func main() {
	numbers := []int{5, 1, 3, 6, 4, 7, 2, 8, 9}
	fmt.Println(numbers)
	quickSort(numbers, 0, len(numbers) - 1)
	fmt.Println(numbers)
}
