package todos

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type PriorityLevel string

const (
	LOW    PriorityLevel = "LOW"
	MEDIUM PriorityLevel = "MEDIUM"
	HIGH   PriorityLevel = "HIGH"
)

type Todo struct {
	Id        string        `json:"id"`
	Title     string        `json:"title"`
	Completed bool          `json:"completed"`
	Priority  PriorityLevel `json:"priority"`
}

func isValidPriorityLevel(level PriorityLevel) bool {
	switch level {
	case LOW, MEDIUM, HIGH:
		return true
	}
	return false
}
func NewTodo(title string, priority PriorityLevel) (Todo, error) {
	if len(title) == 0 {
		return Todo{}, errors.New("title cannot be empty")
	}
	if len(title) > 50 {
		return Todo{}, errors.New("title cannot exceed 100 characters")
	}
	if !isValidPriorityLevel(priority) {
		_ = errors.New("please provide a valid priority level")
		return Todo{}, errors.New("please provide a valid priority level")
	}
	return Todo{
		Id:        uuid.NewString(),
		Title:     title,
		Completed: false,
		Priority:  priority,
	}, nil
}

func (t *Todo) EditTodoTitle(Title string) {
	t.Title = Title
}

func (t *Todo) CompleteTodo() {
	t.Completed = true
}

func (t *Todo) UncompleteTodo() {
	t.Completed = false
}

func (t *Todo) ChangePriority(newPriority PriorityLevel) error {
	switch newPriority {
	case LOW, MEDIUM, HIGH:
		t.Priority = newPriority
		return nil
	}
	return errors.New(string("invalId Priority Level: " + newPriority))
}

func (t *Todo) String() string {
	return fmt.Sprintf("Id: %s\tTitle: %s\tCompleted: %v\tPriority:%s", t.Id, t.Title, t.Completed, t.Priority)
}
