// Los ejecutables usan siempre package main
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func limpiarTexto(texto string) string {
	var sb strings.Builder
	for _, r := range texto {
		// Solo dejamos letras, números y espacios
		if unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSpace(r) {
			sb.WriteRune(unicode.ToLower(r)) // Convertimos a minúscula
		}
	}
	return sb.String()
}

func main() {
	// Pedir al usuario que escriba una frase
	fmt.Print("Escribe una frase: ")

	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n') // Lee hasta presionar Enter

	// Limpiar puntuación y convertir a minúsculas
	frase = limpiarTexto(frase)

	// Convertimos el texto en slice de palabras
	palabras := strings.Fields(frase)

	// Mapa para contar ocurrencias
	ocurrencias := make(map[string]int)

	// Contar
	for _, palabra := range palabras {
		ocurrencias[palabra]++
	}

	// Imprimir resultados
	fmt.Println("\nConteo de palabras:")
	for palabra, count := range ocurrencias {
		fmt.Printf("%s: %d\n", palabra, count)
	}
}
