package main

type EntityRepository interface {
	Create(text string) (*Todo, error)
	FindAll() ([]*Todo, error)
}

type InMemoryRepository struct {
	todos []*Todo
}

func (r *InMemoryRepository) Create(text string) (*Todo, error) {
	t := NewTodo(text)
	r.todos = append(r.todos, t)
	return t, nil
}

func (r *InMemoryRepository) FindAll() ([]*Todo, error) {
	return r.todos, nil
}

func NewInMemoryRepository() EntityRepository {
	return &InMemoryRepository{
		todos: make([]*Todo, 0),
	}
}
