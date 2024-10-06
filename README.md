# Task Management API in Go

This project implements a simple task management API in Go that supports basic CRUD operations (Create, Read, Update, Delete) for tasks. Tasks are stored in an in-memory slice (no database is required).

## Project Structure

- `main.go`: The main file where all the API logic is implemented.
- Tasks are stored as structs in an in-memory slice.

## Task Structure

Each task is represented by the following struct:

```go
type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      string `json:"status"` // "pending" or "completed"
}
```
## Implement CRUD operations:

- **Create** (POST `/tasks`): Create a new task by accepting a JSON payload with a title and description.

- **Read**:
  - Get all tasks (GET `/tasks`): Return a list of all tasks.
  - Get a task by ID (GET `/tasks/{id}`): Return a specific task by its ID.

- **Update** (PUT `/tasks/{id}`): Update the title, description, or status of an existing task.

- **Delete** (DELETE `/tasks/{id}`): Delete a task by its ID.

## Data storage:

- For this assignment, tasks are stored in an in-memory slice of `Task` structs (no database required).

