package main

type EntityRepository interface {
	Create(text string) (*Todo, error)
	FindAll() ([]*Todo, error)
}
