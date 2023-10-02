package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	er := NewInMemoryRepository()
	h := NewHandlers(er)
	r := mux.NewRouter()
	r.HandleFunc("/todos", h.FindAllTodosHandler).Methods(http.MethodGet)
	r.HandleFunc("/todos", h.CreateTodoHandler).Methods(http.MethodPost)
	s := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
