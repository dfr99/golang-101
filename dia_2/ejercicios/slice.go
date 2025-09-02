// Los ejecutables usan siempre package main;
package main

import "fmt"

func main() {
		// Creación de slice dinámico sin elementos
    var pares []int
		// make(<data_type>, <initial_elements>, <max_capacity>)
		impares := make([]int, 0, 2)

    // Llenar dinámicamente con números pares
    for i := 0; i < 10; i++ {
      pares = append(pares, i*2)
    }

		// Llenar dinámicamente con números impares
    for i := 1; i < 10; i++ {
      impares = append(impares, i*2 - 1)
    }

    fmt.Println("Números pares:", pares)
		fmt.Println("Números impares:", impares)
}
