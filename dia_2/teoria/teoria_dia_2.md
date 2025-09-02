Perfecto — te doy una **explicación detallada** (orientada a 30 min de teoría) de Arrays, Slices, Maps, Structs, Métodos y Punteros en Go, con ejemplos prácticos y **referencias a la documentación oficial**. Cada bloque contiene la idea clave, comportamiento importante, ejemplo mínimo y “puntos de cuidado”.

---

# 1) Arrays — definición y semántica

**Idea clave:** Un array en Go es una secuencia numerada con longitud fija; **la longitud forma parte del tipo** (`[N]T`). Las asignaciones copian todo el array (semántica por valor).

* Declaración: `var a [3]int` o `b := [3]int{1,2,3}`.
* `len(a)` devuelve la longitud; `cap(a)` devuelve lo mismo para arrays.
* Pasar un array a una función copia todo el array; para mutarlo desde la función pasa `*[N]T`.

Ejemplo:

```go
package main
import "fmt"

func modifyCopy(a [3]int) { a[0] = 99 }     // copia: no cambia el original
func modifyPtr(a *[3]int) { a[0] = 99 }    // cambia el original

func main() {
    arr := [3]int{1,2,3}
    modifyCopy(arr)
    fmt.Println(arr) // [1 2 3]
    modifyPtr(&arr)
    fmt.Println(arr) // [99 2 3]
}
```

Puntos a cuidar: los arrays son útiles cuando necesitas **semántica por valor** o tamaño fijo; en la práctica se usan menos que slices. ([go.dev][1])

---

# 2) Slices — descriptor, len/cap, append y compartir memoria

**Idea clave:** Un *slice* es un descriptor (puntero + len + cap) que referencia un segmento contiguo de un array subyacente. Su tipo es `[]T`. El valor cero es `nil`.

* Crear: literal `[]T{...}`, `make([]T, len)` o `make([]T, len, cap)`.
* Slicing: `a[low:high]` (half-open). `len(s)` y `cap(s)` devuelven longitud y capacidad.
* `append(s, x)` añade elementos; si `cap` no alcanza, `append` asigna un **nuevo** array subyacente y copia elementos. Dos slices que compartían el mismo array **pueden divergir** después de un `append` que provoque realloc.
* Compartir memoria: `t := s[:2]` comparte el mismo array; modificar `s[i]` se ve en `t` (mientras compartan el mismo array).

Ejemplo:

```go
s := []int{1,2,3}
t := s[:2]   // comparte el array subyacente
s[0] = 10
fmt.Println(s, t) // [10 2 3] [10 2]

s = append(s, 4)  // puede realocar; entonces t y s dejarán de compartir memoria
```

Puntos a cuidar:

* `var s []int` → `s == nil`, `len==0`, `cap==0`. `make([]int,0)` produce slice no-nil con `len==0`.
* Eficiencia: `append` amortiza crecimientos, pero la re-asignación puede copiar datos.
  Documentación oficial y explicación de internals. ([go.dev][2])

---

# 3) Maps — tabla hash, inicialización, operaciones y concurrencia

**Idea clave:** `map[Key]Value` implementa una tabla hash. La *zero value* de un map es `nil`; para insertar necesitas `make` (o literal). Las claves deben ser de tipos **comparable** (no slices, maps, funciones).

* Crear: `m := make(map[string]int)` o `m := map[string]int{"x":1}`.
* Leer: `v := m["k"]` (si k no existe devuelves zero value); forma idiomática: `v, ok := m["k"]`.
* Borrar: `delete(m, "k")`.
* Iteración: `for k,v := range m { ... }` — **el orden no está garantizado**.
* **Concurrencia:** las operaciones simultáneas (writes concurrentes) en un `map` **no son seguras** y pueden producir panic `concurrent map read and map write` — si necesitas acceso concurrente usa sincronización (`sync.Mutex`, `sync.RWMutex`) o `sync.Map` (estructura segura para concurrencia).

Ejemplo:

```go
m := make(map[string]int)
m["a"] = 1
v, ok := m["b"]
if !ok { fmt.Println("clave ausente") }
delete(m, "a")
for k, v := range m { fmt.Println(k, v) }
```

Puntos a cuidar:

* No uses `m["k"]=x` si `m` es `nil` (panic). Inicializa con `make`.
* Para concurrencia: `sync.Map` o proteger con mutex. ([go.dev][3], [Go Packages][4])

---

# 4) Structs — definición, literales, exportación y tags

**Idea clave:** `struct` agrupa campos (posiblemente de distintos tipos). Se declara con `type T struct { ... }`. Los nombres que empiezan con mayúscula son **exportados** (visibles fuera del paquete).

* Literales: `p := Person{Name: "Ana", age: 30}` (keyed) o `Person{"Ana", 30}` (positional — menos claro).
* Campos anónimos y “promoted fields” existen (embebimiento).
* **Struct tags**: se usan para metadata (por ejemplo `json:"name,omitempty"`), muy comunes con `encoding/json` u otros paquetes.

Ejemplo:

```go
type Person struct {
    Name string `json:"name"`
    age  int    // no exportado
}

p := Person{Name: "Ana", age: 30}
q := &Person{}      // puntero a struct
```

Puntos a cuidar:

* La visibilidad (export) depende de la primera letra.
* Los tags son cadenas literales y no tienen semántica salvo que las use un paquete concreto (ej. `encoding/json`). ([golang.google.cn][5], [go.dev][2])

---

# 5) Métodos asociados a structs — receivers (valor vs puntero) y reglas

**Idea clave:** Los métodos se declaran con un *receiver* antes del nombre; puede ser de valor `func (p Person) Foo()` o de puntero `func (p *Person) Bar()`.

* Sintaxis:

```go
func (p Person) Greet() string { return "Hola " + p.Name }   // value receiver
func (p *Person) SetName(n string) { p.Name = n }           // pointer receiver
```

* **Semántica:**

  * *Value receiver* recibe una **copia** del valor; modificaciones no afectan al original.
  * *Pointer receiver* recibe puntero y puede modificar el valor original.
* **Reglas de method set / interfaces (resumen útil):**

  * El *method set* de `T` contiene los métodos con receptor `T`.
  * El *method set* de `*T` contiene métodos con receptor `T` **y** `*T`.
  * Por conveniencia, si `x` es **addressable** (por ejemplo una variable), puedes llamar a un método con receiver pointer aunque escribas `x.m()` — el compilador aplica `&x` automáticamente. Pero hay casos (p. ej. valores no addressable como elementos de un `map`) donde esa conversión implícita **no** es posible.
* **Práctica recomendada:** Si alguno de los métodos necesita modificar el receptor o el struct es grande (coste de copia), usa *pointer receivers* de forma consistente en ese tipo. Esto también evita sorpresas con interfaces (si el método requerido tiene receiver `*T`, sólo `*T` implementa la interfaz). ([go.dev][6])

---

# 6) Punteros (`*`, `&`) — introducción práctica

**Idea clave:** `*T` es tipo “puntero a T”. `&v` toma la dirección; `*p` desreferencia (accede al valor). La zero value de un puntero es `nil`. `new(T)` devuelve `*T` inicializado al zero value.
Ejemplo:

```go
func inc(x *int) { *x++ }

func main() {
    v := 10
    inc(&v)
    fmt.Println(v) // 11

    p := new(int)  // *int apuntando al 0 inicial
    *p = 5
    fmt.Println(*p) // 5
}
```

Puntos a cuidar:

* Desreferenciar un `nil` causa panic (runtime).
* `&T{...}` es la forma idiomática de crear un puntero a un struct con valores iniciales. ([go.dev][7])

---

# Consejos prácticos y trampas comunes

* **Slices vs Arrays:** usa slices casi siempre (dinámicos y cómodos). Usa arrays cuando **quieras la semántica por valor** o tamaños fijos por diseño. ([go.dev][1])
* **Comparabilidad de claves en maps:** tipos con `==` (int, string, pointers, structs cuyos campos sean comparables) — no se pueden usar slices, maps ni funciones como claves. ([go.dev][3])
* **Concurrencia y maps:** nunca permitas escrituras concurrentes sin sincronización — usa `sync.Mutex`/`sync.RWMutex` o `sync.Map`. ([go.dev][8], [Go Packages][4])
* **Receivers consistentes:** evita mezclar receptores valor y puntero en el mismo tipo salvo motivos claros; preferencia: si necesitas mutar, usa puntero para todos los métodos. ([go.dev][9])

---

# Mini-ejercicios (para practicar en 15–30 min)

1. Crear un `[]int`, hacer `append` hasta superar la capacidad inicial y comprobar si los slices comparten memoria antes/después de append.
2. Declarar un `map[string][]int`, añadir pares, iterar y borrar claves. Comprobar comportamiento con `v, ok := m["x"]`.
3. Definir `type Counter struct{ n int }` con métodos `Inc()` y `Value() int`. Implementar `Inc()` como pointer receiver y `Value()` como value receiver. Probar llamadas desde `Counter` y `*Counter`.
4. Escribir función que reciba `[3]int` y otra que reciba `*[3]int` y observar diferencias.

Si quieres, te doy las soluciones al final o las corro en un playground y te devuelvo el código listo para copiar.

---

# Referencias oficiales (selección)

* Especificación del lenguaje (definitions: slices, arrays, types, method sets). ([go.dev][2])
* A Tour of Go — lecciones prácticas sobre arrays, slices, maps, structs, métodos y punteros. ([go.dev][1], [golang.google.cn][5])
* Go blog (slices intro, maps in action) — contexto y detalles de implementación útiles. ([go.dev][10])
* `sync.Map` (uso en concurrencia). ([Go Packages][4])

---

[1]: https://go.dev/tour/moretypes/6?utm_source=chatgpt.com "Arrays"
[2]: https://go.dev/ref/spec?utm_source=chatgpt.com "The Go Programming Language Specification"
[3]: https://go.dev/tour/moretypes/19?utm_source=chatgpt.com "Maps"
[4]: https://pkg.go.dev/sync?utm_source=chatgpt.com "sync package"
[5]: https://golang.google.cn/tour/list?utm_source=chatgpt.com "A Tour of Go"
[6]: https://go.dev/doc/go1.17_spec?utm_source=chatgpt.com "The Go Programming Language Specification"
[7]: https://go.dev/tour/moretypes?utm_source=chatgpt.com "Pointers"
[8]: https://go.dev/blog/maps?utm_source=chatgpt.com "Go maps in action"
[9]: https://go.dev/tour/methods/8?utm_source=chatgpt.com "Choosing a value or pointer receiver"
[10]: https://go.dev/blog/slices-intro?utm_source=chatgpt.com "Go Slices: usage and internals"
