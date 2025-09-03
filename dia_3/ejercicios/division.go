package main

import (
	"errors"
	"fmt"
)

// Divide recibe dos números float64 y devuelve el resultado y un error
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}
	return a / b, nil
}

func main() {
	resultado, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}

	resultado, err = Divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}
