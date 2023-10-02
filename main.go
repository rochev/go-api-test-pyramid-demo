package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", FindAllTodosHandler).Methods(http.MethodGet)
	r.HandleFunc("/todos", CreateTodoHandler).Methods(http.MethodPost)
	s := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
