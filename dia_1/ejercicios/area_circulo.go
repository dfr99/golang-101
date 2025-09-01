// Los ejecutables usan siempre package main;
package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
	"math"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
  fmt.Print("¿Cuál es el radio de tu círculo?: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var radio float64
	if num, err := strconv.Atoi(input); err == nil {
    fmt.Println("✅ Es un número entero")
    radio = float64(num)
  } else if numf, err := strconv.ParseFloat(input, 64); err == nil {
    fmt.Println("✅ Es un número flotante")
    radio = numf
  } else {
    fmt.Println("❌ Error: lo ingresado no es un número válido.")
    return
  }

  fmt.Println("El área de tu círculo es", math.Pi*radio*radio)
}
