package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Task representa una tarea sencilla.
// Se serializa/deserializa en JSON usando tags (ej: {"id":1, "title":"..."}).
type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// Variables globales para almacenar las tareas en memoria.
// Se usa un slice y un contador de IDs.
// sync.Mutex asegura concurrencia segura cuando hay múltiples requests.
var (
	mu     sync.Mutex
	tasks  = make([]Task, 0)
	nextID = 1
)

// writeJSON serializa la respuesta en JSON y la envía con el status indicado.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// writeError devuelve un error en formato JSON.
func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]any{"error": msg})
}

// pingHandler responde "pong" → sirve como chequeo rápido del servidor.
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte("pong"))
}

// tasksHandler maneja /tasks.
// - GET: lista todas las tareas.
// - POST: crea una nueva tarea.
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		w.Header().Set("Allow", "GET, POST")
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
	}
}

// taskByIDHandler maneja /tasks/{id}.
// - GET: devuelve una tarea específica según su ID.
func taskByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		writeError(w, http.StatusMethodNotAllowed, "método no permitido")
		return
	}

	// Extraer el ID desde la URL.
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/tasks/"), "/")
	if len(parts) < 1 || parts[0] == "" {
		writeError(w, http.StatusBadRequest, "ID requerido")
		return
	}
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	// Buscar la tarea con ese ID en el slice protegido por mutex.
	mu.Lock()
	defer mu.Unlock()
	for _, t := range tasks {
		if t.ID == id {
			writeJSON(w, http.StatusOK, t)
			return
		}
	}

	// Si no existe, devolver 404.
	writeError(w, http.StatusNotFound, fmt.Sprintf("tarea con id %d no encontrada", id))
}

// listTasks devuelve todas las tareas como un slice en JSON.
func listTasks(w http.ResponseWriter, _ *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	writeJSON(w, http.StatusOK, tasks)
}

// taskInput define el payload para crear/actualizar una tarea.
type taskInput struct {
	Title string `json:"title"`
	Done  *bool  `json:"done,omitempty"` // opcional al crear
}

// validateTitle asegura que el título no esté vacío ni sea demasiado largo.
func validateTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("title no puede estar vacío")
	}
	if len(title) > 200 {
		return errors.New("title demasiado largo (máx 200)")
	}
	return nil
}

// createTask crea una nueva tarea y la agrega al slice global.
func createTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var in taskInput

	// Deserializar el JSON recibido.
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "JSON inválido")
		return
	}
	if err := validateTitle(in.Title); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Proteger acceso concurrente al slice de tareas.
	mu.Lock()
	id := nextID
	nextID++
	newTask := Task{
		ID:        id,
		Title:     in.Title,
		Done:      false,
		CreatedAt: time.Now().UTC(),
	}
	tasks = append(tasks, newTask)
	mu.Unlock()

	// Responder con 201 Created y la tarea recién creada.
	w.Header().Set("Location", "/tasks/"+strconv.Itoa(id))
	writeJSON(w, http.StatusCreated, newTask)
}

func main() {
	// Crear router y asignar handlers.
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/tasks", tasksHandler)
	mux.HandleFunc("/tasks/", taskByIDHandler)

	// Configuración del servidor HTTP.
	addr := ":8080"
	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Servidor escuchando en http://localhost%v\n", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error al iniciar servidor: %v", err)
	}
}
