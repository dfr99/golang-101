// Los ejecutables usan siempre package main;
package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

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

	var nums []int
	for i := 0; i < numero; i++ {
		nums = append(nums, i)
	}

		for num := range nums {
			fmt.Println("Número", num+1)
		}
}
