Perfecto — te dejo una explicación detallada, con ejemplos y **referencias a la documentación oficial** para cada punto del Día 5. Iré sección por sección (idiomática, `defer`, comentarios/documentación y las herramientas `gofmt`/`go vet`/`golint`), y al final un mini checklist práctico que puedes aplicar al mini-proyecto.

# Principios de Go *idiomático* (nombres, errores, simplicidad)

**Resumen corto:** Go favorece nombres sencillos y consistentes, manejo explícito de errores (errores como valores), APIs pequeñas/claras y la simplicidad por encima de las abstracciones complejas. Estas pautas aparecen en *Effective Go*, en las “Code Review Comments” y en las proverbios de Go. ([go.dev][1], [go-proverbs.github.io][2])

**Nombres**

* Paquetes: nombres cortos, en minúsculas y normalmente una sola palabra (ej. `http`, `json`).
* Identificadores multi-palabra: usar `MixedCaps` (exportado: `ExportedName`, no\_exportado: `unexportedName`) en vez de `snake_case`.
* Interfaces de un solo método: nombrarlas por el verbo + `-er` (ej. `Reader`, `Writer`).
* Getters: si el nombre del campo es `owner`, el getter se llama `Owner()` en vez de `GetOwner()`.
  Estas convenciones están en *Effective Go* y en el wiki de revisiones de código. ([go.dev][1])

Ejemplo corto:

```go
package store // paquete corto y descriptivo

type Saver interface { // interfaz de un método -> verbo+er
    Save(data []byte) error
}

type db struct{} // receptor no exportado: mixedCaps
func (d *db) Save(data []byte) error { /* ... */ return nil }
```

**Errores (idiomática)**

* No uses `panic` para errores esperables; devuelve `error` y deja que el llamador decida. (Uso de `panic` reservado para condiciones verdaderamente excepcionales). ([go.dev][1])
* Desde Go 1.13 hay wrapping y helpers estándar: `fmt.Errorf("...: %w", err)` para envolver; `errors.Is` y `errors.As` para inspeccionar. Usa `errors` y `fmt` de la stdlib en lugar de soluciones externas salvo necesidad. ([Go Packages][3], [go.dev][4])

Ejemplo:

```go
if err := doWork(); err != nil {
    return fmt.Errorf("doWork failed: %w", err) // envolver con contexto
}

if errors.Is(err, sql.ErrNoRows) {
    // manejo especial si la causa es ErrNoRows
}
```

**Simplicidad**

* Prefiere APIs pequeñas (pequeñas interfaces, funciones claras), composición sobre herencia compleja y evitar “sobre-ingeniería”. Las *Go Proverbs* resumen esto (“The bigger the interface, the weaker the abstraction”, “Make the zero value useful”, etc.). ([go-proverbs.github.io][2], [go.dev][1])

---

# Uso de `defer`

**Qué hace `defer`:** programa una llamada para que se ejecute justo antes de que la función que la declaró retorne. Se usa comúnmente para cerrar recursos y liberar locks. Las llamadas `defer` se ejecutan en orden LIFO (última `defer` ejecutada primero) y los argumentos de la llamada se evalúan en el momento que se declara el `defer`. ([go.dev][1])

Ejemplo clásico:

```go
func copyFile(src, dst string) error {
    f, err := os.Open(src)
    if err != nil { return err }
    defer f.Close()  // se ejecutará al salir de la función

    // ... leer f, escribir en dst ...
    return nil
}
```

Puntos prácticos:

* Coloca el `defer` justo después de obtener el recurso (mejora legibilidad y evita fugas). ([go.dev][1])
* **Evaluación de argumentos:** en `defer fmt.Println(x)` el valor de `x` se toma **en el momento** del `defer`, no cuando se ejecuta al final.
* Si necesitas liberar en un bucle de alto rendimiento, evalúa la frecuencia/overhead: `defer` es ideal para claridad y seguridad de limpieza; en caminos extremadamente críticos de rendimiento podrías preferir liberar manualmente (si mediste y comprobaste). (La documentación oficial explica la semántica y su uso típico; para decisiones de micro-optimización, medir es la regla). ([go.dev][1])

---

# Documentación con comentarios (godoc / pkg.go.dev)

**Reglas clave (oficiales):**

* Cada paquete debería tener un *package comment* (normalmente en `doc.go`) cuyo primer enunciado empiece con `Package <name> ...`.
* Los nombres exportados deben tener doc comments con oraciones completas que expliquen el propósito y comportamiento público. Los comentarios que docan una declaración deben ir inmediatamente antes de la declaración (sin líneas en blanco).
* Puedes usar funciones `Example...` para ejemplos ejecutables que aparecen en la documentación y se prueban con `go test`. La nueva página “Go Doc Comments” recoge las normas actuales. ([tip.golang.org][5], [go.dev][6])

Ejemplo:

```go
// Package calc provides simple arithmetic helpers.
package calc

// Add returns the sum of a and b.
func Add(a, b int) int {
    return a + b
}

// ExampleAdd demonstrates Add usage.
func ExampleAdd() {
    fmt.Println(Add(2, 3))
    // Output: 5
}
```

Consejos:

* Escribe la primera frase como resumen breve; el resto puede contener detalles de comportamiento, errores esperados, invariantes y ejemplos de uso.
* Mantén la documentación junto al código para que evolucione con él; usa `pkg.go.dev` y `gopls` en el editor para feedback inmediato. ([tip.golang.org][5], [go.dev][7])

---

# Breve vistazo: `gofmt`, `go vet` y `golint`

**`gofmt`**

* Formatea el código de manera canonical: tabs para indentación y espacios para alineación. Úsalo como primera medida: `gofmt -w .` o configura tu editor para formatear al guardar. Esto elimina debates de estilo y mejora la coherencia. ([Go Packages][8], [go.dev][9])

**`go vet`**

* Herramienta de análisis estático mantenida por el equipo Go; detecta construcciones sospechosas (ej.: `Printf` con argumentos erróneos) y otras heurísticas útiles que el compilador no detecta. Se invoca con `go vet ./...` o `go vet package`. No sustituye tests, pero encuentra bugs sutiles. ([Go Packages][10], [go.dev][11])

**`golint`**

* `golint` (en `golang.org/x/lint`) fue una herramienta popular para sugerencias de estilo, pero actualmente está **deprecada / frozen**; el equipo recomienda usar `go vet` y linters modernos (ej. `staticcheck` o `golangci-lint`) para análisis más completos. En resumen: *no* confíes en golint como la única regla; prefieres herramientas actuales en CI. ([Go Packages][12], [golangci-lint.run][13])

Comando de ejemplo para CI (sugerencia):

```sh
# formato
gofmt -l .   # lista archivos sin formatear
gofmt -w .

# análisis rápido
go vet ./...

# linter más completo (instalar previamente)
golangci-lint run
```

---

# Mini-checklist práctico para tu mini-proyecto

1. `gofmt` automático (editor o precommit). ([Go Packages][8])
2. Documenta el paquete (`doc.go`) y todos los símbolos exportados; añade ejemplos ejecutables. ([tip.golang.org][5])
3. Maneja errores explícitamente (devuelve `error`), envuélvelos con `%w` cuando añades contexto; usa `errors.Is/As` para inspección. ([Go Packages][3], [go.dev][4])
4. Usa `defer` para cerrar/limpiar justo después de reservar recursos. ([go.dev][1])
5. Añade `go vet` + `golangci-lint` o `staticcheck` en CI. ([Go Packages][10], [golangci-lint.run][13])

---

[1]: https://go.dev/doc/effective_go?utm_source=chatgpt.com "Effective Go"
[2]: https://go-proverbs.github.io/?utm_source=chatgpt.com "Go Proverbs"
[3]: https://pkg.go.dev/errors?utm_source=chatgpt.com "errors package"
[4]: https://go.dev/doc/tutorial/handle-errors?utm_source=chatgpt.com "Return and handle an error"
[5]: https://tip.golang.org/doc/comment?utm_source=chatgpt.com "Go Doc Comments"
[6]: https://go.dev/doc/?utm_source=chatgpt.com "Documentation"
[7]: https://go.dev/wiki/Comments?utm_source=chatgpt.com "Go Wiki: Comments"
[8]: https://pko.dev/cmd/gofmt?utm_source=chatgpt.com "cmd/gofmt"
[9]: https://go.dev/src/cmd/gofmt/doc.go?utm_source=chatgpt.com "Gofmt formats Go programs."
[10]: https://pkg.go.dev/cmd/vet?utm_source=chatgpt.com "cmd/vet - command"
[11]: https://go.dev/src/cmd/vet/README?utm_source=chatgpt.com "cmd/vet/README"
[12]: https://pkg.go.dev/golang.org/x/lint?utm_source=chatgpt.com "lint package - golang.org/x/lint - Go ..."
[13]: https://golangci-lint.run/docs/linters/configuration/?utm_source=chatgpt.com "Settings"
