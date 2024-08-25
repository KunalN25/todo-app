package main

import (
	"fmt"
	"net/http"
)

func main() {
	RegisterRoutes()
	startServer()
}

func startServer() {
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

/*
 Few things can add

* Filter todos by title, completed status, priority
*/
