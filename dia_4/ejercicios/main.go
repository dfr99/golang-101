package main

import (
	"fmt"

	"demo-math/mathutils"
)

func main() {
	fmt.Println("Factorial(5):", mathutils.Factorial(5))
	fmt.Println("FactorialIter(5):", mathutils.FactorialIter(5))

	fmt.Println("Fibonacci(10):", mathutils.Fibonacci(10))
	fmt.Println("FibonacciIter(10):", mathutils.FibonacciIter(10))
}
