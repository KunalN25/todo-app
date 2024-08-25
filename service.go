package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"todo-app/store"
	"todo-app/todos"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	filters := todos.GetTodosFilters{
		Title:    r.URL.Query().Get("title"),
		Priority: todos.PriorityLevel(r.URL.Query().Get("priority")),
	}

	// Handle the "completed" filter as a special case (it can be true, false, or not set)
	if completed := r.URL.Query().Get("completed"); completed != "" {
		completedValue := completed == "true"
		filters.Completed = &completedValue
	}

	todoData, err := store.LoadTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Filter the todos based on the provided filters
	filteredTodos := []todos.Todo{}
	for _, todo := range todoData {
		if filters.Title != "" && !strings.Contains(todo.Title, filters.Title) {
			continue
		}
		if filters.Priority != "" && todo.Priority != filters.Priority {
			continue
		}
		if filters.Completed != nil && todo.Completed != *filters.Completed {
			continue
		}
		filteredTodos = append(filteredTodos, todo)
	}
	todosJSONResponse, err := json.Marshal(filteredTodos)
	if err != nil {
		http.Error(w, "Failed to encode todos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(todosJSONResponse)
	if writeErr != nil {
		return
	}
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todoRequest todos.AddTodoRequest

	err := json.NewDecoder(r.Body).Decode(&todoRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo, err := todos.NewTodo(todoRequest.Title, todoRequest.Priority)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Todos, err := store.LoadTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Todos = append(Todos, todo)

	err = store.SaveTodos(Todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func findTodoById(Todos []todos.Todo, id string) (*todos.Todo, int, error) {
	for idx, _ := range Todos {
		if Todos[idx].Id == id {
			return &Todos[idx], idx, nil
		}
	}
	return nil, -1, errors.New("todo not found")
}

func EditTodo(w http.ResponseWriter, r *http.Request) {
	var todoRequest todos.EditTodoRequest
	err := json.NewDecoder(r.Body).Decode(&todoRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Todos, _ := store.LoadTodos()
	todo, _, err := findTodoById(Todos, todoRequest.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if todoRequest.Title != "" {
		todo.Title = todoRequest.Title
	}
	if todoRequest.Priority != "" {
		err := todo.ChangePriority(todoRequest.Priority)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	err = store.SaveTodos(Todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var deleteTodoRequest todos.DeleteTodoRequest
	err := json.NewDecoder(r.Body).Decode(&deleteTodoRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Todos, _ := store.LoadTodos()
	_, idx, err := findTodoById(Todos, deleteTodoRequest.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Todos = append(Todos[:idx], Todos[idx+1:]...)
	err = store.SaveTodos(Todos)
}

func CompleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var completeTodoRequest todos.CompleteTodoRequest
	err := json.NewDecoder(r.Body).Decode(&completeTodoRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Todos, _ := store.LoadTodos()
	todo, _, err := findTodoById(Todos, completeTodoRequest.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.CompleteTodo()
	err = store.SaveTodos(Todos)
}

func UncompleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var completeTodoRequest todos.CompleteTodoRequest
	err := json.NewDecoder(r.Body).Decode(&completeTodoRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Todos, _ := store.LoadTodos()
	todo, _, err := findTodoById(Todos, completeTodoRequest.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.UncompleteTodo()
	err = store.SaveTodos(Todos)
}
