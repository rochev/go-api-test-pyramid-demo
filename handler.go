package main

import "net/http"

func CreateTodoHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("TODO"))
}

func FindAllTodosHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("TODO"))
}
