package stringutils

import "testing"

func TestInvertirCadena(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hola", "aloh"},
		{"", ""},
		{"a", "a"},
		{"murciélago", "ogaléicrum"},
	}

	for _, tt := range tests {
		result := InvertirCadena(tt.input)
		if result != tt.expected {
			t.Errorf("InvertirCadena(%q) = %q; esperado %q", tt.input, result, tt.expected)
		}
	}
}

func TestContarVocales(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"hola", 2},
		{"HOLA", 2},
		{"xyz", 0},
		{"áéíóú", 5},
		{"murciélago", 5},
	}

	for _, tt := range tests {
		result := ContarVocales(tt.input)
		if result != tt.expected {
			t.Errorf("ContarVocales(%q) = %d; esperado %d", tt.input, result, tt.expected)
		}
	}
}
