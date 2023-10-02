//go:build unit

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTodo(t *testing.T) {
	todo := NewTodo("foo bar")
	assert.Equal(t, "foo bar", todo.Text)
}
