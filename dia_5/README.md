# API REST simple en Go

Este proyecto implementa una **API REST minimalista** en Go utilizando únicamente el paquete estándar `net/http`. Permite gestionar un listado de tareas en memoria.

---

## 🚀 Endpoints disponibles

### `GET /ping`

Verifica que el servidor esté en funcionamiento.

```bash
curl http://localhost:8080/ping
```

Respuesta:

```
pong
```

### `GET /tasks`

Lista todas las tareas.

```bash
curl http://localhost:8080/tasks
```

Respuesta inicial (sin tareas):

```json
[]
```

### `POST /tasks`

Crea una nueva tarea.

```bash
curl -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"Aprender Go"}'
```

Respuesta:

```json
{
  "id": 1,
  "title": "Aprender Go",
  "done": false,
  "created_at": "2025-09-04T15:00:00Z"
}
```

### `GET /tasks/{id}`

Obtiene una tarea específica por ID.

```bash
curl http://localhost:8080/tasks/1
```

Respuesta:

```json
{
  "id": 1,
  "title": "Aprender Go",
  "done": false,
  "created_at": "2025-09-04T15:00:00Z"
}
```

Si el ID no existe:

```json
{"error":"tarea con id 99 no encontrada"}
```

---

## 🛠️ Cómo ejecutar

1. Asegúrate de tener [Go](https://go.dev/dl/) instalado (v1.20+).
2. Clona este repositorio o copia el archivo `main.go`.
3. Ejecuta:

   ```bash
   go run main.go
   ```
4. El servidor escuchará en `http://localhost:8080`.

---

## 📦 Dependencias

Este proyecto **no usa librerías externas**, únicamente la biblioteca estándar de Go:

* `net/http` → servidor HTTP.
* `encoding/json` → serialización JSON.
* `sync.Mutex` → concurrencia segura.
* `time` → timestamps.

---

## 📌 Notas

* Los datos se guardan solo en memoria (se pierden al reiniciar el servidor).
* `sync.Mutex` asegura que múltiples clientes puedan usar la API al mismo tiempo sin conflictos.
* Se valida el campo `title` para evitar entradas vacías o demasiado largas.

---

## 🔮 Posibles mejoras

* Soporte para `PUT /tasks/{id}` (actualizar una tarea).
* Soporte para `DELETE /tasks/{id}` (eliminar una tarea).
* Persistencia en una base de datos.
* Organización del código en múltiples archivos (`handlers.go`, `models.go`, etc.).
