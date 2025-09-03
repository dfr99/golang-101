¡Vamos línea por línea! 👇

```go
package main
```

* Define el paquete principal. En Go, el paquete `main` indica que este archivo compila a un ejecutable con una función `main`.

```go
import (
	"context"
	"fmt"
	"sync"
	"time"
)
```

* Importa paquetes estándar:

  * `context`: para cancelación/timeout y propagación de señales.
  * `fmt`: formato e impresión por consola.
  * `sync`: primitivas de sincronización (aquí, `WaitGroup`).
  * `time`: manejo de tiempos y duraciones.

```go
// worker: procesa trabajos hasta que el contexto se cancele o se acaben
```

* Comentario que describe la función `worker`.

```go
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
```

* Declaración del **worker**:

  * `ctx`: contexto para cancelar.
  * `id`: identificador del worker (útil para logs/depurar).
  * `jobs <-chan int`: canal **solo lectura** de trabajos (números).
  * `results chan<- int`: canal **solo escritura** de resultados.
  * `wg *sync.WaitGroup`: para avisar cuándo este worker termina.

```go
	defer wg.Done()
```

* Asegura que, al salir de la función (por la razón que sea), se llama a `wg.Done()` para decrementar el contador del `WaitGroup`.

```go
	for {
```

* Bucle infinito: el worker procesa trabajos hasta que reciba cancelación o no queden trabajos.

```go
		select {
```

* Multiplexa entre varias “esperas” (cancelación o nuevos trabajos).

```go
		case <-ctx.Done():
			// Cancelación: salimos inmediatamente
			return
```

* Si el contexto se cancela (timeout o cancelación manual), el worker termina de inmediato.

```go
		case num, ok := <-jobs:
```

* Intenta leer un trabajo del canal `jobs`.

  * `num` es el valor leído.
  * `ok` indica si el canal sigue abierto (`true`) o ya fue cerrado (`false`).

```go
			if !ok {
				// No más trabajos
				return
			}
```

* Si `jobs` está cerrado y no hay más trabajos, el worker finaliza.

```go
			results <- num * num
```

* Calcula el cuadrado y lo envía al canal `results`.

```go
		}
	}
}
```

* Cierre del `select`, del `for` y de la función `worker`.

```go
func main() {
```

* Punto de entrada del programa.

```go
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
```

* Lista de números a procesar.

```go
	jobs := make(chan int, len(numeros))
	results := make(chan int, len(numeros))
```

* Crea canales con **buffer**:

  * `jobs`: cola de trabajos con capacidad igual al número de tareas.
  * `results`: cola de resultados con la misma capacidad (evita bloqueos innecesarios).

```go
	numWorkers := 3
```

* Tamaño del pool: cuántos workers se lanzan en paralelo.

```go
	var wg sync.WaitGroup
```

* `WaitGroup` para esperar a que **todos** los workers terminen.

```go
	// Creamos un contexto con timeout de 2 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
```

* Crea un contexto derivado de `Background()` que se **cancela automáticamente** tras 2 segundos.

```go
	defer cancel()
```

* Garantiza liberar recursos del contexto (y permite cancelación manual si se sale antes).

```go
	// Lanzamos los workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, results, &wg)
	}
```

* Inicia `numWorkers` goroutines:

  * `wg.Add(1)`: incrementa el contador por cada worker lanzado.
  * `go worker(...)`: arranca el worker concurrentemente.

```go
	// Enviamos los números a procesar
	for _, n := range numeros {
		jobs <- n
	}
```

* Publica cada número como un “job” en el canal `jobs`.

```go
	close(jobs)
```

* Cierra `jobs` para indicar “no habrá más trabajos”.
  Los workers que sigan leyendo verán `ok == false` y terminarán cuando consuman todo lo pendiente o al recibir cancelación.

```go
	// Cerramos results cuando todos los workers terminen
	go func() {
		wg.Wait()
		close(results)
	}()
```

* Goroutine de cierre:

  * Espera a que **todos** los workers llamen `wg.Done()`.
  * Luego **cierra** `results`.
    Esto permite que el `range` sobre `results` (más abajo) termine de forma limpia.

```go
	// Sumamos los resultados
	suma := 0
	for r := range results {
		suma += r
	}
```

* Acumulador: lee del canal `results` hasta que se cierre y suma cada cuadrado.

```go
	// Mostramos el resultado
	if ctx.Err() != nil {
		fmt.Println("Operación cancelada:", ctx.Err())
	} else {
		fmt.Printf("La suma de cuadrados es: %d\n", suma)
	}
}
```

* Si el contexto terminó con error (por ejemplo, **timeout**), informa cancelación.
* Si no hubo error (se completó antes del timeout), imprime la suma final.

---

### Notas clave

* **Cancelación vs. completado**: si se dispara el timeout, los workers salen por `ctx.Done()` y puede que no se procesen todos los trabajos; sumarás solo resultados parciales.
* **Direccionalidad de canales** en la firma del worker (`<-chan`, `chan<-`) evita usos incorrectos en tiempo de compilación.
* **Cierre ordenado**: se cierra `jobs` cuando no hay más trabajos y `results` solo después de que **todos** los workers han acabado, evitando lecturas/escrituras sobre canales cerrados.
