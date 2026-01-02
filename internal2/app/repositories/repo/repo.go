package repo

import "github.com/trng-tr/golab4/internal2/domain"

type Repository interface {
	Save(user *domain.User) (*domain.User, error)
}
