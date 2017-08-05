package models

type Todo struct {
	Title string
	Done  bool
}

func NewTodo(title string) *Todo {
	return &Todo{title, false}
}
