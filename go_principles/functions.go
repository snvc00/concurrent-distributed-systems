package main

import "fmt"

const (
	FibonacciLimit = 610
	NumbersLimit   = 12
)

// Homework 5: Functions and recursion
func main() {
	// Bubble sort on slice
	fmt.Println("Bubble sort")
	s := []int{5, 4, 3, 2, 1, 0, 0}
	fmt.Println(s)
	bubbleSort(s)
	fmt.Println(s)

	// Recursive Fibonacci
	fmt.Println("Fibonacci")
	fibonacci(0, 1)

	// Generate odd numbers
	fmt.Println("Odd numbers")
	generateOdd := oddGenerator()
	for i := 0; i < NumbersLimit-1; i++ {
		fmt.Print(generateOdd(), ", ")
	}
	fmt.Println(generateOdd())

	// Swap two numbers using pointers
	fmt.Println("Swap two numbers")
	var a, b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	swap(&a, &b)
	fmt.Println(a, "\b,", b)
}

func bubbleSort(s []int) {
	var tmp int
	for i := 1; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j+1] < s[j] {
				tmp = s[j]
				s[j] = s[j+1]
				s[j+1] = tmp
			}
		}
	}
}

func fibonacci(x0 int, x1 int) {
	fmt.Print(x0, ", ")
	t := x0
	x0 = x1
	x1 += t

	if x0 > FibonacciLimit {
		fmt.Println("\b\b")
		return
	}

	fibonacci(x0, x1)
}

func oddGenerator() func() uint {
	i := uint(1)

	return func() uint {
		var odd = i
		i += 2
		return odd
	}
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}
