package main

import (
	"fmt"
	"strings"
)

// RepeatStr repeats value parameter repetition parameter times
func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

func main() {
	fmt.Println(RepeatStr(6, "I"))
	fmt.Println(RepeatStr(5, "Hello"))
}
