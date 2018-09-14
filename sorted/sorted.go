package main

import (
	"fmt"
)

// InAscOrder returns true if numbers are in ascending order and false
// otherwise.
func InAscOrder(numbers []int) bool {

	if len(numbers) < 1 {
		return false
	}

	curr := numbers[0]

	for _, elem := range numbers {
		if elem < curr {
			return false
		}
		curr = elem
	}

	return true
}

func main() {
	fmt.Println(InAscOrder([]int{1, 2, 4, 7, 19}))
	fmt.Println(InAscOrder([]int{1, 2, 3, 4, 5}))
	fmt.Println(InAscOrder([]int{1, 6, 10, 18, 2, 4, 20}))
	fmt.Println(InAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}))
}
