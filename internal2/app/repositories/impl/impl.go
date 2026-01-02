package impl

import (
	"errors"
)

// RepositoryImpl struct to implement RepositoryImpl
type RepositoryImpl[T any] struct {
	GetId func(*T) string
	data  map[string]*T
}

// NewRepositoryImpl func constructor, DI par constructeur
func NewRepositoryImpl[T any](getId func(*T) string) *RepositoryImpl[T] {
	return &RepositoryImpl[T]{
		GetId: getId,
		data:  make(map[string]*T, 0),
	}
}

// Save: implement interface
func (ri *RepositoryImpl[T]) Save(t *T) (*T, error) {
	if t == nil {
		return nil, errors.New("error: object user is nil")
	}
	objectId := ri.GetId(t)
	if objectId == "" {
		return nil, errors.New("error: id is empty")
	}
	ri.data[objectId] = t

	return t, nil
}
