//go:build e2e

package main

import (
	"github.com/gavv/httpexpect/v2"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodoHandlers_CreateTodoHandler(t *testing.T) {
	er := NewInMemoryRepository()
	h := NewHandlers(er)
	r := mux.NewRouter()
	r.HandleFunc("/todos", h.CreateTodoHandler).Methods(http.MethodPost)

	// run server using httptest
	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	e.POST("/todos").
		Expect().
		Status(http.StatusCreated).
		JSON().Object().ContainsKey("text").HasValue("text", "New Todo")
}

func TestTodoHandlers_FindAllTodosHandler(t *testing.T) {
	er := NewInMemoryRepository()
	_, _ = er.Create("foo")
	_, _ = er.Create("bar")
	_, _ = er.Create("baz")
	h := NewHandlers(er)
	r := mux.NewRouter()
	r.HandleFunc("/todos", h.FindAllTodosHandler).Methods(http.MethodGet)

	// run server using httptest
	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.Default(t, server.URL)
	js := e.GET("/todos").
		Expect().
		Status(http.StatusOK).
		JSON()
	js.Array().Length().IsEqual(3)
	js.Array().Value(0).Object().HasValue("text", "foo")
	js.Array().Value(1).Object().HasValue("text", "bar")
	js.Array().Value(2).Object().HasValue("text", "baz")
}
