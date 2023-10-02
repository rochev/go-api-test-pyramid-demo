package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	s := http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
