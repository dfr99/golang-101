# golang-101

## 📅 Plan de Aprendizaje de Go en 5 días (2h/día)

### Día 1 – Fundamentos esenciales

#### 👉 Objetivo: entender la sintaxis básica y cómo se escribe un programa en Go.

##### Teoría (30 min):

- Estructura de un programa Go (package `main`, `func main()`).
- Tipos básicos: `int`, `float`, `string`, `bool`.
- Variables (`var`, `:=`).
- Funciones y parámetros.
- Control de flujo: `if`, `for`, `switch`.

#####  Ejercicios (1h):

- Imprimir “Hola Go”.
- Calcular el área de un círculo (con funciones).
- Iterar sobre un array con `for`.
- Uso de switch para clasificar edades.

##### Reto (30 min):

- 👉 Escribe un programa que reciba un número y diga si es primo o no.

### Día 2 – Composición y colecciones

#### 👉 Objetivo: dominar colecciones y trabajar con estructuras propias.

##### Teoría (30 min):

- Arrays, slices y mapas.
- Structs: definición y uso.
- Métodos asociados a structs.
- Introducción a punteros (`*`, `&`).

##### Ejercicios (1h):

- Crear un slice dinámico y agregar elementos.
- Usar un mapa para contar ocurrencias de palabras en una lista.
- Definir un struct Persona con nombre y edad, y un método que diga si es mayor de edad.

##### Reto (30 min):

👉 Programa un pequeño gestor de contactos con un mapa (nombre → teléfono) y permite agregar, buscar y listar contactos.

### Día 3 – Concurrencia y manejo de errores

#### 👉 Objetivo: usar las características únicas de Go.

##### Teoría (30 min):

- Error como tipo de retorno.
- Múltiples valores de retorno en funciones.
- Goroutines (`go` keyword).
- Canales (`chan`) y comunicación entre goroutines.
- select para multiplexar canales.

##### Ejercicios (1h):

- Escribir una función que devuelva (resultado, error) al dividir dos números (manejar división por cero).
- Lanzar 3 goroutines que impriman mensajes en paralelo.
- Enviar números a través de un canal y procesarlos en otra goroutine.

##### Reto (30 min):

👉 Implementa un programa concurrente que calcule la suma de cuadrados de una lista de números, distribuyendo el cálculo entre varias goroutines.

### Día 4 – Paquetes, testing y módulos

#### 👉 Objetivo: organizar proyectos y probar código.

##### Teoría (30 min):

- Creación y uso de paquetes (`package`).
- Organización en múltiples archivos.
- Módulos con go mod init.
- Introducción al testing con `go test`.
- Subtests y ejemplos.

##### Ejercicios (1h):

- Separar un programa en dos paquetes (`main` y `mathutils`).
- Crear funciones en mathutils (ej: Factorial, Fibonacci).
- Escribir pruebas unitarias para esas funciones.
- Correr los tests con `go test`.

##### Reto (30 min):

👉 Construye un módulo stringutils con funciones para:

- Invertir una cadena.
- Contar vocales.
- Incluye tests unitarios.

### Día 5 – Proyecto final y buenas prácticas

#### 👉 Objetivo: integrar todo en un mini proyecto completo.

##### Teoría (30 min):

- Principios de Go idiomático (nombres, errores, simplicidad).
- Uso de defer.
- Documentación con comentarios.
- Breve vistazo a `go fmt`, `go vet` y `golint`.

##### Ejercicio guiado (1h):

- Proyecto: API REST simple con Go usando `net/http`.
- Endpoints:
    - /ping (devuelve “pong”)
    - /tasks (listar tareas)
    - /tasks [POST] (agregar tarea).
- Estructura con struct Task y slice global de tareas.
- Uso de JSON (serialización/deserialización).
- Concurrencia segura con `sync.Mutex`.

##### Reto (30 min):

👉 Extiende la API agregando un endpoint /tasks/{id} que devuelva una tarea específica.
