package main

import "fmt"

// century returns century number from year
func century(year int) int {
	if year%100 == 0 {
		return year / 100
	}

	return year/100 + 1
}

func main() {
	fmt.Println(century(1705))
	fmt.Println(century(1900))
	fmt.Println(century(1601))
	fmt.Println(century(2000))
}
