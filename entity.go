package main

type Todo struct {
	Text string
}

func NewTodo(text string) *Todo {
	return &Todo{Text: text}
}
