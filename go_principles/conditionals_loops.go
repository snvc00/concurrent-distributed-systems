package main

import "fmt"

func main() {
	// Homework 3: Conditionals and loops
	zodiac()
	eNumber()
}

func zodiac() {
	var (
		day   = 0
		month = 0
	)

	fmt.Scan(&day)
	fmt.Scan(&month)

	switch {
	case (day >= 21 && day <= 31 && month == 3) || (day >= 1 && day <= 20 && month == 4):
		fmt.Println("aries")
	case (day >= 21 && day <= 30 && month == 4) || (day >= 1 && day <= 20 && month == 5):
		fmt.Println("tauro")
	case (day >= 21 && day <= 31 && month == 5) || (day >= 1 && day <= 21 && month == 6):
		fmt.Println("geminis")
	case (day >= 22 && day <= 30 && month == 6) || (day >= 1 && day <= 22 && month == 7):
		fmt.Println("cancer")
	case (day >= 23 && day <= 31 && month == 7) || (day >= 1 && day <= 22 && month == 8):
		fmt.Println("leo")
	case (day >= 23 && day <= 31 && month == 8) || (day >= 1 && day <= 22 && month == 9):
		fmt.Println("virgo")
	case (day >= 23 && day <= 30 && month == 9) || (day >= 1 && day <= 22 && month == 10):
		fmt.Println("libra")
	case (day >= 23 && day <= 31 && month == 10) || (day >= 1 && day <= 22 && month == 11):
		fmt.Println("escorpio")
	case (day >= 23 && day <= 30 && month == 11) || (day >= 1 && day <= 21 && month == 12):
		fmt.Println("sagitario")
	case (day >= 22 && day <= 31 && month == 12) || (day >= 1 && day <= 20 && month == 1):
		fmt.Println("capricornio")
	case (day >= 21 && day <= 31 && month == 1) || (day >= 1 && day <= 18 && month == 2):
		fmt.Println("acuario")
	case (day >= 19 && day <= 29 && month == 2) || (day >= 1 && day <= 20 && month == 3):
		fmt.Println("piscis")
	}
}

func eNumber() {
	var e float64
	for i := 0.0; i < 100.0; i++ {
		e += 1.0 / factorial(i)
	}

	fmt.Println("E =", e)
}

func factorial(n float64) float64 {
	f := 1.0
	for i := 1.0; i <= n; i++ {
		f *= i
	}

	return f
}
