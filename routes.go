package main

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/getTodos", GetTodos)
	http.HandleFunc("/addTodo", AddTodo)
	http.HandleFunc("/editTodo", EditTodo)
	http.HandleFunc("/deleteTodo", DeleteTodo)
	http.HandleFunc("/completeTodo", CompleteTodo)
	http.HandleFunc("/uncompleteTodo", UncompleteTodo)
}
