package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Creamos un canal de enteros
	numeros := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2) // vamos a lanzar dos goroutines

	// Goroutine que envía números
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			numeros <- i*2
			time.Sleep(time.Millisecond * 1000) // Simula trabajo
		}
		close(numeros) // Cerramos el canal al terminar
	}()

	// Goroutine que procesa números
	go func() {
		defer wg.Done()
		for n := range numeros {
			fmt.Printf("Procesando número: %d\n", n)
		}
		fmt.Println("No hay más números, terminando...")
	}()

	// Esperamos a que ambas goroutines acaben
	wg.Wait()
}
