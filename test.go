package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Result", sum(1, 2, 3, 4, 5, 6, 8, 9, 10))
}
