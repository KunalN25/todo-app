package todos

type AddTodoRequest struct {
	Title    string        `json:"title"`
	Priority PriorityLevel `json:"priority"`
}

type EditTodoRequest struct {
	Id        string        `json:"id"`
	Title     string        `json:"title"`
	Completed bool          `json:"completed"`
	Priority  PriorityLevel `json:"priority"`
}

type DeleteTodoRequest struct {
	Id string `json:"id"`
}

type CompleteTodoRequest struct {
	Id string `json:"id"`
}
