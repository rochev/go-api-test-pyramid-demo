package main

type EntityRepository interface {
	Create(text string) (*Todo, error)
	FindAll() ([]*Todo, error)
}

type inMemoryRepository struct {
}

func (r *inMemoryRepository) Create(text string) (*Todo, error) {
	return NewTodo(text), nil
}

func (r *inMemoryRepository) FindAll() ([]*Todo, error) {
	return []*Todo{
		NewTodo("foo"),
		NewTodo("bar"),
		NewTodo("baz"),
	}, nil
}

func NewInMemoryRepository() EntityRepository {
	return &inMemoryRepository{}
}
