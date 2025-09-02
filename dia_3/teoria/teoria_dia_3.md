Perfecto — aquí tienes una explicación detallada (en español) de los puntos que pides, con ejemplos prácticos y referencias a la documentación oficial de Go para que puedas consultar cada detalle.

# Día 3 — Concurrencia y manejo de errores (explicación detallada)

---

## 1) **Error como tipo de retorno** — concepto y práctica idiomática

En Go `error` es un **tipo predeclarado** (una interfaz):

```go
type error interface {
    Error() string
}
```

Esto significa que cualquier valor que implemente `Error() string` satisface `error`. En Go la convención idiomática es devolver un `error` como **último** valor de retorno para indicar fallo en la operación; por ejemplo `func Read() (T, error)`. Comprobar `err != nil` es la forma habitual de decidir si continuar o manejar el fallo. ([go.dev][1])

**Patrones y prácticas importantes**

* Comprueba `if err != nil { ... }` inmediatamente tras la llamada. ([go.dev][2])
* Envuelve errores para mantener contexto usando `fmt.Errorf("%w", err)` y luego distingue/inspecciona con `errors.Is` / `errors.As` (introducido y documentado en las notas de Go 1.13). Esto permite propagar y reconocer causas concretas sin perder la traza/causa original. ([go.dev][2], [Go Packages][3])
* Evita `panic` para control de errores normal — `panic` se usa sólo en situaciones excepcionales (errores irreparables). Para diseño de API, prefiere errores como valores (pattern *errors are values*). ([go.dev][4])

**Ejemplo corto:**

```go
func LoadConfig(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("abrir config %s: %w", path, err)
    }
    defer f.Close()
    return io.ReadAll(f)
}

data, err := LoadConfig("cfg.json")
if err != nil {
    // manejar o devolver el error
}
```

(Nota: `errors` como valores y `fmt.Errorf("%w", ...)` están cubiertos en la documentación oficial). ([go.dev][2], [Go Packages][3])

---

## 2) **Múltiples valores de retorno en funciones**

Go permite que una función devuelva varios valores en su firma, por ejemplo `func Div(a, b int) (int, error)`. Este rasgo facilita devolver un resultado **y** un `error` sin necesidad de estructuras ad-hoc. También se usan para operaciones que naturalmente devuelven pares (p. ej. `key, value`, `x, y` o `result, ok`). ([go.dev][5])

**Puntos clave y ejemplos**

* Sintaxis:

```go
func Swap(a, b int) (int, int) { return b, a }
```

* Patrón resultado+error:

```go
func Div(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("división por cero")
    }
    return a / b, nil
}
```

* Puedes ignorar valores con el *blank identifier* `_`: `x, _ := someFunc()`. ([go.dev][5])

Nota: la convención de poner `error` como último valor es idiomática (no impuesta por el compilador), pero seguirla mejora la legibilidad y la interoperabilidad con herramientas y ejemplos en el ecosistema Go. ([go.dev][5])

---

## 3) **Goroutines (`go` keyword)** — hilos ligeros gestionados por el runtime

La instrucción `go <expr>` arranca una **goroutine**, que es una unidad de ejecución ligera programada por el runtime de Go (no es un hilo OS por sí sola). Ejemplos:

```go
go doWork()             // llama a doWork() en una nueva goroutine
go func() { fmt.Println("hola") }()  // función literal en goroutine
```

Las goroutines son baratas en recursos comparadas con hilos OS y permiten escribir concurrencia con muchas rutinas activas. ([go.dev][6])

**Sincronización inicial / memoria**
El modelo de memoria de Go define garantías sobre la sincronización entre la goroutine que crea y la creada: **la sentencia `go` está sincronizada antes del inicio de la goroutine**, lo que asegura ciertas visibilidades iniciales (pero aún necesitas sincronización explícita para coordinar datos cambiantes). Para órdenes/visibilidad más finas usa canales, `sync` o `context`. ([go.dev][7])

**Cosas prácticas a tener en cuenta**

* Evitar fugas de goroutines: si una goroutine queda bloqueada por siempre (p. ej. esperando en un canal que nunca recibe), es una fuga. Usa canales, `context.Context` o `WaitGroup` para coordinar terminación.
* El `panic` en una goroutine no afecta a otras goroutines a menos que no se capture; si quieres evitar abortos globales, usa `recover` dentro de la propia goroutine. ([go.dev][6])

---

## 4) **Canales (`chan`) y comunicación entre goroutines**

Los **canales** son conductos tipados para enviar/recibir valores entre goroutines. Se crean con `make`:

```go
ch := make(chan int)     // canal sin buffer (bloqueante)
buf := make(chan int, 3) // canal con buffer de capacidad 3
```

Operadores básicos: `ch <- v` (enviar), `v := <-ch` (recibir). Usar `range ch` para iterar hasta que el canal se cierre. ([go.dev][8])

**Semántica importante**

* Canales **sin buffer**: el `send` bloquea hasta que hay un receptor listo y el `receive` bloquea hasta que hay un valor — sincronización de mano a mano (rendezvous).
* Canales **buffered**: el `send` sólo bloquea si el buffer está lleno; el `receive` bloquea si el buffer está vacío. Esto permite amortiguar productores/consumidores. ([go.dev][8])
* Cerrar canales: `close(ch)` indica que no habrá más envíos; los receptores pueden detectar cierre con la forma `v, ok := <-ch` (si `!ok` canal cerrado). **Sólo el remitente debe cerrar** el canal. La documentación builtin describe el comportamiento tras el cierre (los `receive` sucesivos devuelven el valor cero y `ok=false`). ([go.dev][9])

**Ejemplo patrón productor/consumidor:**

```go
func producer(out chan<- int) {
    for i := 0; i < 5; i++ { out <- i }
    close(out)
}

func consumer(in <-chan int) {
    for v := range in { fmt.Println(v) } // termina cuando channel se cierra
}

ch := make(chan int)
go producer(ch)
consumer(ch)
```

(Patrón documentado en el blog oficial "Go Concurrency Patterns — pipelines"). ([go.dev][10])

---

## 5) **`select` para multiplexar canales**

`select` permite esperar en múltiples operaciones de canal y ejecutar la que esté lista. Sintaxis básica:

```go
select {
case v := <-ch1:
    fmt.Println("recibido", v)
case ch2 <- 1:
    fmt.Println("enviado a ch2")
default:
    fmt.Println("ninguna comunicación lista")
}
```

Comportamientos clave:

* `select` bloquea hasta que **alguna** de sus cláusulas pueda proceder; si varias pueden, la elección es **pseudo-aleatoria** (uniforme entre las listas preparadas). Esto evita sesgos entre casos simultáneos. ([Go Packages][11], [go.dev][12])
* `default` permite que el `select` sea *no bloqueante*: si ninguna comunicación está lista, se ejecuta el `default`.
* Patrón de timeout con `time.After` (canal que se activa tras un retraso) es muy común para abortar operaciones que tardan demasiado:

```go
select {
case res := <-resultCh:
    return res, nil
case <-time.After(2 * time.Second):
    return nil, fmt.Errorf("timeout")
}
```

`select` es la herramienta principal para orquestar goroutines y canales de forma limpia y sin locks en muchos casos. ([Go Packages][11], [go.dev][10])

---

## Buenas prácticas resumidas (rápido)

* **Errores**: devuélvelos como valores (`... , error`), comprueba `err != nil`, envuélvelos (`%w`) y usa `errors.Is`/`errors.As`. Evita `panic` salvo casos excepcionales. ([go.dev][4])
* **Múltiples retornos**: úsalos para resultado+error; el orden `(..., error)` es convención. ([go.dev][5])
* **Goroutines**: son baratas, pero coordina su finalización (no dejes goroutines bloqueadas). Usa `context`, canales o `sync.WaitGroup`. ([go.dev][7])
* **Canales**: elige buffered vs unbuffered según el acoplamiento deseado; no cierres canales desde el receptor; usa `range` sobre canales cerrados para terminar bucles de consumo. ([go.dev][8])
* **select**: ideal para multiplexar e implementar timeouts, cancelación y patrones de orquestación. ([Go Packages][11], [go.dev][10])

---

## Referencias oficiales (principal)

* *The error type* — blog/documentación sobre manejo de errores en Go. ([go.dev][1])
* *Errors are values* (blog oficial, explicación idiomática). ([go.dev][4])
* *Working with Errors in Go 1.13* (wrapping, Is/As). ([go.dev][2])
* *Package errors* (std lib). ([Go Packages][3])
* *Effective Go* (estilo e idioms, múltiples retornos). ([go.dev][5])
* *The Go Programming Language Specification* (select, canales, tipos, etc.). ([go.dev][12])
* *Tour of Go — Concurrency* (canales, ejemplos). ([go.dev][8])
* *The Go memory model* (sincronización y garantías al arrancar goroutines). ([go.dev][7])
* *Go Concurrency Patterns — Pipelines* (patrones de cierre/pipe). ([go.dev][10])
* *reflect.Select / implementación de select (comportamiento pseudo-aleatorio)*. ([Go Packages][11])

---

[1]: https://go.dev/blog/error-handling-and-go?utm_source=chatgpt.com "Error handling and Go"
[2]: https://go.dev/blog/go1.13-errors?utm_source=chatgpt.com "Working with Errors in Go 1.13"
[3]: https://pkg.go.dev/errors?utm_source=chatgpt.com "errors package"
[4]: https://go.dev/blog/errors-are-values?utm_source=chatgpt.com "Errors are values"
[5]: https://go.dev/doc/effective_go?utm_source=chatgpt.com "Effective Go"
[6]: https://go.dev/blog/defer-panic-and-recover?utm_source=chatgpt.com "Defer, Panic, and Recover"
[7]: https://go.dev/ref/mem?utm_source=chatgpt.com "The Go Memory Model"
[8]: https://go.dev/tour/concurrency/2?utm_source=chatgpt.com "v // Send v to channel ch. v"
[9]: https://go.dev/pkg/builtin/?m=old&utm_source=chatgpt.com "Package builtin"
[10]: https://go.dev/blog/pipelines?utm_source=chatgpt.com "Go Concurrency Patterns: Pipelines and cancellation"
[11]: https://pkg.go.dev/reflect?utm_source=chatgpt.com "reflect package"
[12]: https://go.dev/ref/spec?utm_source=chatgpt.com "The Go Programming Language Specification"
