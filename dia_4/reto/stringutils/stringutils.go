package stringutils

import "strings"

// InvertirCadena devuelve la cadena invertida.
func InvertirCadena(s string) string {
	runes := []rune(s) // Para soportar caracteres Unicode
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

// ContarVocales cuenta el número de vocales en la cadena (mayúsculas y minúsculas).
func ContarVocales(s string) int {
	vocales := "aeiouáéíóú"
	count := 0
	for _, r := range strings.ToLower(s) {
		if strings.ContainsRune(vocales, r) {
			count++
		}
	}
	return count
}
