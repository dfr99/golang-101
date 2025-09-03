package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Lanzamos 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 📌 Le decimos al WaitGroup que hay una goroutine pendiente
		go func(id int) {
			defer wg.Done() 
			// 📌 Cuando la goroutine termine, se llama a Done() para avisar que ha acabado

			// Simulamos trabajo con 5 mensajes
			for j := 1; j <= 5; j++ {
				fmt.Printf("Goroutine %d → Mensaje %d\n", id, j)
				time.Sleep(300 * time.Millisecond) // Pausa para ver el paralelismo
			}
		}(i)
	}

	// 📌 Wait() bloquea la ejecución del main hasta que todas las goroutines hayan llamado a Done()
	wg.Wait()
	fmt.Println("✅ Todas las goroutines han terminado")
}
