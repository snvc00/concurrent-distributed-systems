package main

import (
	"fmt"
	"math"
)

// Homework 4: Arrays and slices
func main() {
	// Size of for the slice
	var n int
	fmt.Scanln(&n)

	// Validation for invalid size
	if n < 1 || n > math.MaxInt32 {
		fmt.Println("Invalid size")
		return
	}

	// Slice to store numbers to sum
	s := make([]int, n)

	// Get numbers from input
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Scanln(&tmp)
		s = append(s, tmp)
	}

	// Calculate sum of the elements
	var r int
	for _, v := range s {
		r += v
	}

	// Print result
	fmt.Println(r)
}