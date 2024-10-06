package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Define Task struct
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // "pending" or "completed"
}

// Creating In-memory storage (simulating a database)
var tasks = []Task{}
var nextID = 1

// TaskHandler handles all the CRUD operations for tasks
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the ID from the URL (if present)
	//idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	idStr := r.URL.Path[7:]
	id, _ := strconv.Atoi(idStr)
	//fmt.Println(id) //Test extracted id

	switch r.Method {
	case http.MethodGet:
		if id > 0 {
			getTaskByID(w, r, id)
		} else {
			getAllTasks(w, r)
		}
	case http.MethodPost:
		createTask(w, r)
	case http.MethodPut:
		updateTask(w, r, id)
	case http.MethodDelete:
		deleteTask(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Get all tasks
func getAllTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

// Get task by ID
// func getTaskByID(w http.ResponseWriter, r *http.Request, id int) {
// 	task, found := getTask(id)
// 	if !found {
// 		http.Error(w, "Task not found", http.StatusNotFound)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(task)
// }

// New Get task by ID
func getTaskByID(w http.ResponseWriter, r *http.Request, id int) {
	for _, task := range tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

// Create a new task
func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)
}

// Update an existing task
func updateTask(w http.ResponseWriter, r *http.Request, id int) {
	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	updatedTask.ID = id
	if updateExistingTask(&updatedTask) {
		json.NewEncoder(w).Encode(updatedTask)
	} else {
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}

// Delete a task by ID
func deleteTask(w http.ResponseWriter, r *http.Request, id int) {
	if deleteExistingTask(id) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Task not found", http.StatusNotFound)
	}
}

// In-memory functions (simulating database actions)

// func getTask(id int) (Task, bool) {
// 	for _, task := range tasks {
// 		if task.ID == id {
// 			return task, true
// 		}
// 	}
// 	return Task{}, false
// }

// UpdateTask updates an existing task
func updateExistingTask(updatedTask *Task) bool {
	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = *updatedTask
			return true
		}
	}
	return false
}

// DeleteTask deletes a task by ID
func deleteExistingTask(id int) bool {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	// Initialize the HTTP router and routes
	http.HandleFunc("/tasks", TaskHandler)  // Handles all task routes
	http.HandleFunc("/tasks/", TaskHandler) // Handles specific task actions

	// Start the HTTP server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
