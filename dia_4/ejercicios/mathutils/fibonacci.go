package mathutils

// Fibonacci devuelve el enésimo número de Fibonacci.
// Definición: Fib(0)=0, Fib(1)=1.
func Fibonacci(n int) int {
	if n < 0 {
		return 0 // Igual que en Factorial, evitamos negativos.
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

package mathutils

// FibonacciIter devuelve el enésimo número de Fibonacci de forma iterativa.
func FibonacciIter(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
