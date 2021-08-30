package main

import (
	"fmt"
	"math"
)

func main() {
	// Homework 1: Learning to compile and run
	helloWorld()

	// Homework 2: Variables
	areaOfSquare()
	areaOfTriangle()
	areaOfCircle()
	fahrenheitToCelsius()
}

func helloWorld() {
	fmt.Println("Hello world!")
}

func areaOfSquare() {
	var width float64

	fmt.Print("\nArea of a square\nWidth in cm: ")
	fmt.Scanln(&width)
	fmt.Println("That square has an area of:", width*width, "cm")
}

func areaOfTriangle() {
	var (
		height = 0.0
		base   = 0.0
	)

	fmt.Print("\nArea of a triangle\nHeight in cm: ")
	fmt.Scanln(&height)
	fmt.Print("Base in cm: ")
	fmt.Scanln(&base)
	fmt.Println("That square has an area of:", (base*height)/2, "cm")
}

func areaOfCircle() {
	var radius float64

	fmt.Print("\nArea of a circle\nRadius in cm: ")
	fmt.Scanln(&radius)
	fmt.Println("That circle has an area of:", math.Pi*(radius*radius), "cm")
}

func fahrenheitToCelsius() {
	var temperature float64

	fmt.Print("\nFahrenheit to Celsius\nTemperature in degrees Fahrenheit: ")
	fmt.Scanln(&temperature)
	fmt.Println("That is equivalent to", (temperature-32)*5/9, "degrees Celsius")
}
