// Los ejecutables usan siempre package main;
package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

var edad int;

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Introduce un único número entero: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numero, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("❌ Error: lo ingresado no es un entero válido.")
		return
	}

	switch {
	case numero < 18:
		fmt.Println("Eres menor de edad.")
	case numero >= 18 && numero < 25:
		fmt.Println("Eres un adulto joven.")
	case numero >= 25 && numero < 40:
		fmt.Println("Eres un adulto.")
	case numero >= 40:
		fmt.Println("Eres una persona madura.")
	}
}
