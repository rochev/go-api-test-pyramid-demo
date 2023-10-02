package main

import (
	"encoding/json"
	"net/http"
)

type TodoHandlers struct {
	er EntityRepository
}

func NewHandlers(er EntityRepository) *TodoHandlers {
	return &TodoHandlers{
		er: er,
	}
}

func (th *TodoHandlers) CreateTodoHandler(w http.ResponseWriter, _ *http.Request) {
	todo, err := th.er.Create("New Todo")
	commonHandler(w, http.StatusCreated, todo, err)
}

func (th *TodoHandlers) FindAllTodosHandler(w http.ResponseWriter, _ *http.Request) {
	todo, err := th.er.FindAll()
	commonHandler(w, http.StatusOK, todo, err)
}

func commonHandler(w http.ResponseWriter, status int, data any, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	marshal, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(marshal)
}
