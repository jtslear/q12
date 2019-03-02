package main

import (
	"fmt"
)

func theQuestion(i int) [][]int {
	fmt.Printf("Input: %d\n", i)
	r := [][]int{}

	// We'll try when allowed 1 step at a time
	attempt := []int{}
	for s := 1; s <= i; s++ {
		attempt = append(attempt, 1)
	}
	r = append(r, attempt)

	// We'll try when allowed 2 steps at a time
	attempt = []int{}
	for s := 1; s <= i; s += 2 {
		attempt = append(attempt, 2)
	}
	r = append(r, attempt)

	return r
}

func main() {
}
