package impl

import (
	"errors"

	"github.com/trng-tr/golab4/internal2/domain"
)

// RepositoryImpl struct to implement RepositoryImpl
type RepositoryImpl struct {
	GetId func(*domain.User) string
	users map[string]*domain.User
}

// NewRepositoryImpl func constructor, DI par constructeur
func NewRepositoryImpl(getId func(*domain.User) string) *RepositoryImpl {
	return &RepositoryImpl{
		GetId: getId,
		users: make(map[string]*domain.User, 0),
	}
}

// Save: implement interface
func (ri *RepositoryImpl) Save(user *domain.User) (*domain.User, error) {
	if user == nil {
		return nil, errors.New("error: object user is nil")
	}
	userId := ri.GetId(user)
	if userId == "" {
		return nil, errors.New("error: id is empty")
	}
	ri.users[userId] = user

	return user, nil
}
