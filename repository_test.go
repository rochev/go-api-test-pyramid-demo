//go:build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryRepository_Create(t *testing.T) {
	r := NewInMemoryRepository()
	todos, err := r.FindAll()
	assert.Nil(t, err)
	assert.Len(t, todos, 0)
	newTodo, err := r.Create("foo bar")
	assert.Nil(t, err)
	assert.Equal(t, "foo bar", newTodo.Text)
	todos, err = r.FindAll()
	assert.Nil(t, err)
	assert.Len(t, todos, 1)
}

func NewInMemoryRepositoryStub() EntityRepository {
	return &InMemoryRepository{
		todos: []*Todo{
			NewTodo("foo"),
			NewTodo("bar"),
			NewTodo("baz"),
		},
	}
}

func TestInMemoryRepository_FindAll(t *testing.T) {
	r := NewInMemoryRepositoryStub()
	todos, err := r.FindAll()
	assert.Nil(t, err)
	assert.Len(t, todos, 3)
	assert.Equal(t, "foo", todos[0].Text)
	assert.Equal(t, "bar", todos[1].Text)
	assert.Equal(t, "baz", todos[2].Text)
}
