### Explicación detallada de ChatGPT

¡Vamos por partes! Aquí tienes los conceptos básicos de Go con ejemplos y enlaces a la documentación oficial.

# Estructura de un programa Go (`package main`, `func main()`)

Un ejecutable en Go debe estar en el paquete `main` y definir una función `main()` como punto de entrada. Suele incluir importaciones y otras funciones auxiliares:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hola, Go 👋")
}
```

* Los ejecutables usan **siempre** `package main`; la función `main()` es la que se ejecuta al iniciar el programa. ([Go][1])
* La guía de inicio y el tutorial muestran cómo compilar y ejecutar (`go build`, `go run`). ([Go][2])

# Tipos básicos: `int`, `float`, `string`, `bool`

Go tiene tipos básicos predeclarados:

* `bool`
* `string`
* Numéricos con signo: `int`, `int8`, `int16`, `int32`, `int64`
* Sin signo: `uint`, `uint8` (`byte`), `uint16`, `uint32`, `uint64`, `uintptr`
* Flotantes: `float32`, `float64`
* Complejos: `complex64`, `complex128`

```go
var a bool = true
var b string = "gopher"
var c int = 42
var d float64 = 3.14
```

La *Tour of Go* resume los tipos y anota detalles de tamaño habituales; la especificación formal los define. ([Go][3])

# Variables (`var`, `:=`)

Hay dos formas idiomáticas de declarar variables:

```go
// 1) Declaración con 'var' (opcionalmente con valor inicial)
var x int
var y = 10        // el tipo se infiere (int)
var s string = "" // tipo explícito

// 2) Declaración corta dentro de funciones
z := 3.5          // solo dentro de funciones, no a nivel de paquete
```

Puntos clave:

* `var` funciona en cualquier ámbito (paquete o función). La inferencia de tipo aplica si das un inicializador. ([Go][4])
* `:=` es **declaración corta** (crea e inicializa) y **solo puede usarse dentro de funciones**. ([Go][5])
* Si no inicializas, la variable toma su **valor cero**: `false`, `0`, `""`, `nil` según el tipo. ([Go][6])

# Funciones y parámetros

Sintaxis básica y rasgos del lenguaje:

```go
// Firma con parámetros y valor de retorno
func add(a, b int) int {
    return a + b
}

// Varios valores de retorno
func divmod(a, b int) (q, r int) {
    q = a / b
    r = a % b
    return // "named result parameters"
}

// Variádicas
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

* Declaración de funciones, parámetros, resultados múltiples y resultados *nombrados* están definidos en la especificación y ejemplificados en la Tour. ([Go][7])
* Go también permite *closures* (funciones que capturan variables externas) cuando lo necesites. ([Go][8])

# Control de flujo: `if`, `for`, `switch`

## `if`

```go
if x > 0 {
    // ...
} else if x == 0 {
    // ...
} else {
    // ...
}

// 'if' con sentencia corta (inicialización previa al condicional):
if v := compute(); v < 100 {
    // v está en el ámbito de este if/else
}
```

El `if` puede incluir una **sentencia corta** antes de la condición; las variables declaradas ahí viven solo dentro del `if`. ([Go][7])

## `for`

Go tiene **un solo** bucle: `for`, con varias formas:

```go
// Tradicional (init; cond; post)
for i := 0; i < 3; i++ { /* ... */ }

// Estilo "while"
for x > 0 { /* ... */ }

// Bucle infinito
for { /* ... */ }

// Con 'range' sobre colecciones
for i, v := range []int{2, 3, 5} {
    _ = i; _ = v
}
```

La especificación detalla las tres formas y el uso de `range`. (Desde Go 1.22 cada iteración de `for` con `range` usa variables nuevas, evitando capturas sorpresivas en closures.) ([Go][7])
La Tour introduce todas las variantes de `for`. ([Go][9])

## `switch`

```go
// Expresión 'switch' (sin 'fallthrough' implícito)
switch day {
case "Mon", "Tue":
    // ...
case "Sat", "Sun":
    // ...
default:
    // ...
}

// 'switch' sin expresión (equivale a 'switch true')
switch {
case x < 0:
    // ...
case x == 0:
    // ...
default:
    // ...
}

// Type switch (sobre interfaces)
var any interface{} = 3.14
switch v := any.(type) {
case int:
    _ = v // v es int
case string:
    _ = v // v es string
default:
    // otros tipos
}
```

* `switch` evalúa casos de arriba a abajo y **no** hace *fallthrough* por defecto (puedes usar la palabra clave `fallthrough` explícitamente). ([Go][10])
* Existe el `switch` sin expresión (equivale a `switch true`) y el **type switch** para descubrir el tipo dinámico de una interfaz. ([Go][11])
* Las reglas completas están en la especificación. ([Go][7])

---

## Referencias oficiales (para seguir profundizando)

* **Cómo escribir código en Go**: estructura de ejecutables (`package main`/`func main`). ([Go][1])
* **Especificación del lenguaje (go1.25)**: declaraciones de funciones, `if`, `for`, `switch`, valores cero, declaración corta, inicialización de paquetes. ([Go][7])
* **A Tour of Go**: tipos básicos, variables (`var`, `:=`), `if` con sentencia corta, `for`, `switch`, *type switch*. ([Go][3])
* **Tutoriales oficiales**: “Getting started”, compilar e instalar. ([Go][2])

Si quieres, te preparo una chuleta imprimible con estos ejemplos y notas.

[1]: https://go.dev/doc/code?utm_source=chatgpt.com "How to Write Go Code - The Go Programming Language"
[2]: https://go.dev/doc/tutorial/getting-started?utm_source=chatgpt.com "Tutorial: Get started with Go"
[3]: https://go.dev/tour/basics/11?utm_source=chatgpt.com "Basic types"
[4]: https://go.dev/tour/basics/8?utm_source=chatgpt.com "Variables"
[5]: https://go.dev/tour/basics/10?utm_source=chatgpt.com "Short variable declarations"
[6]: https://go.dev/tour/basics/12?utm_source=chatgpt.com "Zero values"
[7]: https://go.dev/ref/spec "The Go Programming Language Specification - The Go Programming Language"
[8]: https://go.dev/tour/moretypes/25?utm_source=chatgpt.com "Function closures"
[9]: https://go.dev/tour/flowcontrol/1?utm_source=chatgpt.com "Loops"
[10]: https://go.dev/tour/flowcontrol/10?utm_source=chatgpt.com "Switch evaluation order"
[11]: https://go.dev/tour/flowcontrol/11?utm_source=chatgpt.com "Switch with no condition"
