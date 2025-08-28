# Día 5 – Proyecto final y buenas prácticas

## 👉 Objetivo: integrar todo en un mini proyecto completo.

### Teoría (30 min):

- Principios de Go idiomático (nombres, errores, simplicidad).
- Uso de defer.
- Documentación con comentarios.
- Breve vistazo a `go fmt`, `go vet` y `golint`.

### Ejercicio guiado (1h):

- Proyecto: API REST simple con Go usando `net/http`.
- Endpoints:
    - /ping (devuelve “pong”)
    - /tasks (listar tareas)
    - /tasks [POST] (agregar tarea).
- Estructura con struct Task y slice global de tareas.
- Uso de JSON (serialización/deserialización).
- Concurrencia segura con `sync.Mutex`.

### Reto (30 min):

👉 Extiende la API agregando un endpoint /tasks/{id} que devuelva una tarea específica.
