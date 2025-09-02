# Día 4 – Paquetes, testing y módulos

Perfecto — abajo tienes una explicación detallada (con ejemplos prácticos) de cada uno de los temas que listaste, **apoyada en la documentación oficial de Go**. Tras cada bloque encontrarás la(s) referencia(s) oficiales usadas.

---

## 1) Paquetes (`package`): qué son y reglas básicas

* **Concepto**: un *paquete* en Go es una colección de archivos fuente `.go` **en el mismo directorio** que se compilan juntos; eso significa que las funciones/ tipos/variables definidas en un archivo son visibles en los demás archivos del mismo paquete (mismo `package` en todos ellos). ([golang.org][1])
* **Declaración**: cada archivo empieza con `package <nombre>`. Para programas ejecutables se usa `package main` y debe haber una función `func main()`. ([golang.org][1])
* **Exportación**: un identificador (función, tipo, campo, etc.) **es exportado** (accesible desde otros paquetes) **si su primer carácter es una letra mayúscula**; en caso contrario es no-exportado (visión solo dentro del paquete). Esto está definido en la especificación del lenguaje. ([go.dev][2])
* **Convención de tests**: puedes escribir tests en el **mismo paquete** (`package foo`) o en el **paquete externo de pruebas** (`package foo_test`) si quieres probar la API pública desde fuera. ([Go Packages][3])

### Ejemplo mínimo (dos archivos de un mismo paquete)

`mathutil/add.go`

```go
package mathutil

// Add devuelve la suma de a y b — nombre exportado (A) => accesible desde otros paquetes.
func Add(a, b int) int {
	return a + b
}
```

`mathutil/sub.go`

```go
package mathutil

// helper no exportado, visible solo dentro del paquete mathutil
func subtract(a, b int) int {
	return a - b
}
```

(ambos archivos se compilán juntos porque están en la misma carpeta y declaran `package mathutil`). ([golang.org][1])

---

## 2) Organización en múltiples archivos y directorios

* **Regla simple**: *todos* los `.go` de un mismo directorio con el mismo `package` forman ese paquete. Usa subdirectorios para agrupar paquetes distintos (por ejemplo `internal/`, `cmd/app/`, `pkg/`, etc.). ([golang.org][1], [go.dev][4])
* **Múltiples comandos en un repo**: es común tener varios directorios con `package main` (cada uno produce un binario) y un `internal/` para código compartido; la guía de layout oficial muestra patrones recomendados. ([go.dev][4])

---

## 3) Módulos (`go mod init`) — gestión de dependencias moderna

* **Qué es un módulo**: un *módulo* es un árbol de código que tiene un fichero `go.mod` en la raíz; el `go.mod` define el *module path* (normalmente el repositorio) y las dependencias/versiones necesarias. ([go.dev][5])
* **`go mod init`**: crea `go.mod` y establece el path del módulo. Después, comandos como `go build`/`go test` añaden dependencias a `go.mod` según las importaciones; `go mod tidy` limpia dependencias no usadas. ([go.dev][6])

### Flujo típico:

```bash
# dentro del directorio del proyecto
go mod init example.com/mymodule
# trabajar, importar paquetes externos
go test ./...
go build ./...
go mod tidy
```

El `go.mod` resultante contendrá al menos `module example.com/mymodule` y la directiva `go <versión>`. ([go.dev][6])

---

## 4) Testing con `go test` — primeros pasos

* **Dónde poner tests**: los archivos de test deben terminar en `_test.go`. `go test` compila esos archivos junto con el paquete y ejecuta las funciones `func TestXxx(t *testing.T)` donde `Xxx` empieza por mayúscula. ([Go Packages][3])
* **Comando `go test`**: por defecto corre los tests del paquete actual; puedes usar `go test ./...` para recorrer todos los paquetes del módulo/repositorio. Flags habituales: `-v` (verbose), `-run` (filtrar por regexp), `-cover`, `-bench`, `-count=1` (desactivar caché), etc. El comportamiento de modo local vs modo lista y la caché están documentados en la página del comando `go`. ([Go Packages][7])

### Ejemplo de test simple

`mathutil/add_test.go`

```go
package mathutil

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	if got != 5 {
		t.Fatalf("Add(2,3) = %d; want 5", got)
	}
}
```

(archivo `_test.go`; `go test` lo compilará/executará). ([Go Packages][3])

---

## 5) Subtests (table-driven + `t.Run`) y buenas prácticas

* **Subtests**: Go soporta subtests con `t.Run(name, func(t *testing.T){ ... })`. Esto encaja muy bien con *table-driven tests* (tabla de casos) y permite ejecutar/filtrar subtests individualmente, controlar paralelismo por subtest y obtener salidas más legibles. Introducido en Go 1.7 y documentado en el blog oficial. ([go.dev][8])

### Ejemplo (table-driven con subtests)

```go
func TestAdd_Table(t *testing.T) {
	tests := []struct{
		name string
		a, b int
		want int
	}{
		{"pos", 2, 3, 5},
		{"neg", -1, -2, -3},
	}

	for _, tt := range tests {
		tt := tt // captura para paralelismo seguro
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel() // opcional: ejecutar subtests en paralelo
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Fatalf("Add(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
```

(usar `tt := tt` si activas `t.Parallel()` para evitar problemas de captura de variables). ([go.dev][8])

---

## 6) `Example` — ejemplos ejecutables/documentación

* **Qué son**: las funciones `ExampleXxx()` son *ejemplos* que pueden compilarse y (si incluyen comentario `// Output:`) también ejecutarse y verificarse como parte de la suite de tests. Son útiles porque sirven a la vez como documentación ejecutable y pruebas ligeras. ([go.dev][9], [tip.golang.org][10])

### Ejemplo de `Example`

```go
package mathutil_test

import (
	"fmt"
	"example.com/mymodule/mathutil"
)

func ExampleAdd() {
	fmt.Println(mathutil.Add(1, 2))
	// Output: 3
}
```

Si hay la línea `// Output: 3`, `go test` ejecutará la función y comparará su salida estándar con ese comentario (ignora espacios en blanco al inicio/fin). Si no incluyes `// Output: ...`, la función se compila pero no se ejecuta. ([go.dev][9], [tip.golang.org][10])

---

## 7) Consejos prácticos / checklist rápido

* Cada paquete = todos los `.go` del mismo directorio con el mismo `package`. ([golang.org][1])
* Exporta solo lo necesario (mayúscula inicial) y documenta las entidades exportadas. ([tip.golang.org][11], [go.dev][2])
* Usa `go mod init` en la raíz del repo; `go.mod` controla el path del módulo y las versiones. `go mod tidy` mantiene `go.mod`/`go.sum` limpias. ([go.dev][6])
* Tests: `_test.go`, funciones `TestXxx(t *testing.T)`, subtests con `t.Run`, ejemplos `ExampleXxx` con `// Output:` para verificaciones. Ejecuta `go test ./...` para todo el módulo. ([Go Packages][3], [go.dev][8])

---

## Recursos oficiales (lectura recomendada)

* **How to write Go code (organización / paquetes)**. ([golang.org][1])
* **Tutorial: Create a Go module** (uso de `go mod init` y flujo básico). ([go.dev][6])
* **Referencia de módulos / go.mod**. ([go.dev][12])
* **Paquete `testing` (documentación de tests, subtests, examples)**. ([Go Packages][3], [go.dev][9])
* **Go command (`go test`) — comportamiento, flags, caché**. ([Go Packages][7])

---

[1]: https://golang.org/doc/code.html?utm_source=chatgpt.com "How to Write Go Code"
[2]: https://go.dev/ref/spec?utm_source=chatgpt.com "The Go Programming Language Specification"
[3]: https://pkg.go.dev/testing?utm_source=chatgpt.com "testing package"
[4]: https://go.dev/doc/modules/layout?utm_source=chatgpt.com "Organizing a Go module"
[5]: https://go.dev/ref/mod?utm_source=chatgpt.com "Go Modules Reference"
[6]: https://go.dev/doc/tutorial/create-module?utm_source=chatgpt.com "Tutorial: Create a Go module"
[7]: https://pkg.go.dev/cmd/go?utm_source=chatgpt.com "go command - cmd/go"
[8]: https://go.dev/blog/subtests?utm_source=chatgpt.com "Using Subtests and Sub-benchmarks"
[9]: https://go.dev/blog/examples?utm_source=chatgpt.com "Testable Examples in Go"
[10]: https://tip.golang.org/src/testing/testing.go?utm_source=chatgpt.com "func TestXxx(*testing.T)"
[11]: https://tip.golang.org/doc/comment?utm_source=chatgpt.com "Go Doc Comments"
[12]: https://go.dev/doc/modules/gomod-ref?utm_source=chatgpt.com "go.mod file reference"
