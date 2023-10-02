package main

type Todo struct {
	Text string `json:"text"`
}

func NewTodo(text string) *Todo {
	return &Todo{Text: text}
}
