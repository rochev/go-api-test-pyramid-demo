package main

type EntityRepository interface {
	Create(text string) (*Todo, error)
	FindAll() ([]*Todo, error)
}

type InMemoryRepository struct {
}

func (r *InMemoryRepository) Create(text string) (*Todo, error) {
	return NewTodo(text), nil
}

func (r *InMemoryRepository) FindAll() ([]*Todo, error) {
	return []*Todo{
		NewTodo("foo"),
		NewTodo("bar"),
		NewTodo("baz"),
	}, nil
}

func NewInMemoryRepository() EntityRepository {
	return &InMemoryRepository{}
}
