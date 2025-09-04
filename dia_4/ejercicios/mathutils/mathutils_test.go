package mathutils

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		n, expected int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{7, 5040},
	}

	for _, tt := range tests {
		result := Factorial(tt.n)
		if result != tt.expected {
			t.Errorf("Factorial(%d) = %d; want %d", tt.n, result, tt.expected)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n, expected int
	}{
		{0, 0},
		{1, 1},
		{5, 5},
		{10, 55},
	}

	for _, tt := range tests {
		result := Fibonacci(tt.n)
		if result != tt.expected {
			t.Errorf("Fibonacci(%d) = %d; want %d", tt.n, result, tt.expected)
		}
	}
}

func TestFactorialIter(t *testing.T) {
	tests := []struct {
		n, expected int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{7, 5040},
	}

	for _, tt := range tests {
		result := FactorialIter(tt.n)
		if result != tt.expected {
			t.Errorf("FactorialIter(%d) = %d; want %d", tt.n, result, tt.expected)
		}
	}
}

func TestFibonacciIter(t *testing.T) {
	tests := []struct {
		n, expected int
	}{
		{0, 0},
		{1, 1},
		{5, 5},
		{10, 55},
	}

	for _, tt := range tests {
		result := FibonacciIter(tt.n)
		if result != tt.expected {
			t.Errorf("FibonacciIter(%d) = %d; want %d", tt.n, result, tt.expected)
		}
	}
}

// ---------- Factorial ----------
func BenchmarkFactorialRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(15) // tamaño moderado para evitar overflow
	}
}

func BenchmarkFactorialIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialIter(15)
	}
}

// ---------- Fibonacci ----------
func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(30) // recursivo explota muy rápido, mejor un valor no tan grande
	}
}

func BenchmarkFibonacciIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciIter(30)
	}
}
