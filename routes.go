package main

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/getTodos", corsMiddleware(GetTodos))
	http.HandleFunc("/addTodo", corsMiddleware(AddTodo))
	http.HandleFunc("/editTodo", corsMiddleware(EditTodo))
	http.HandleFunc("/deleteTodo", corsMiddleware(DeleteTodo))
	http.HandleFunc("/completeTodo", corsMiddleware(CompleteTodo))
	http.HandleFunc("/uncompleteTodo", corsMiddleware(UncompleteTodo))
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return funcName(next)
}

func funcName(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Allow any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow specific HTTP methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// Allow specific headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}
