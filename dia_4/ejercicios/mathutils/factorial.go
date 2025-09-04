package mathutils

// Factorial calcula el factorial de un número n.
// Si n es 0 o 1, devuelve 1.
func Factorial(n int) int {
	if n < 0 {
		return 0 // Podríamos manejar error, pero aquí devolvemos 0.
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// FactorialIter calcula el factorial de un número n de manera iterativa.
func FactorialIter(n int) int {
	if n < 0 {
		return 0
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
