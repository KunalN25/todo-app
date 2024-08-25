package store

import (
	"encoding/json"
	"os"
	"todo-app/todos"
)

const filename = "store/todo-data.json"

func LoadTodos() ([]todos.Todo, error) {
	var todoData []todos.Todo
	file, err := os.Open(filename)
	if err == nil {
		// File exists, so read the existing todos
		defer file.Close()
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&todoData); err != nil {
			return nil, err // Return error if decoding fails
		}
	} else if !os.IsNotExist(err) {
		return nil, err // Return error if file cannot be opened for reasons other than not existing
	}

	return todoData, nil
}

func SaveTodos(todoData []todos.Todo) error {
	file, err := os.Create(filename)
	if err != nil {
		return err // Return error if file creation fails
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Optional: for pretty-printing
	if err := encoder.Encode(todoData); err != nil {
		return err // Return error if encoding fails
	}

	return nil // Return nil if successful
}
