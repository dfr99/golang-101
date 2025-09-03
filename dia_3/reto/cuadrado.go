package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: cancelado\n", id)
			return
		case num, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: no hay más trabajos, terminando\n", id)
				return
			}
			fmt.Printf("Worker %d: procesando %d\n", id, num)
			time.Sleep(200 * time.Millisecond) // Simulamos trabajo
			results <- num * num
		}
	}
}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	jobs := make(chan int, len(numeros))
	results := make(chan int, len(numeros))

	numWorkers := 3
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Lanzamos los workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, results, &wg)
	}

	// Enviamos los trabajos
	for _, n := range numeros {
		jobs <- n
	}
	close(jobs)

	// Cerramos results cuando los workers terminen
	go func() {
		wg.Wait()
		close(results)
	}()

	// Sumamos los resultados
	suma := 0
	for r := range results {
		suma += r
	}

	if ctx.Err() != nil {
		fmt.Println("Operación cancelada:", ctx.Err())
	} else {
		fmt.Printf("La suma de cuadrados es: %d\n", suma)
	}
}
