package domain

import (
	"errors"
)

// User structure
type User struct {
	Id        string
	Firstname string
	Lastname  string
	Email     string
	Poids     float64
}

// NewUser : function construtor or structure
func NewUser(id, firstname, lastname, email string, poids float64) (*User, error) {
	if firstname == "" || lastname == "" || email == "" || poids <= 50 {
		return &User{}, errors.New("error: one or more fields invalid")
	}
	return &User{
		Id:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Poids:     poids,
	}, nil
}
