package dtos

import "errors"

// UserRequest structure
type UserRequest struct {
	Firstname string
	Lastame   string
	Poids     float64
}

// NewUserRequest : function construtor or structure
func NewUserRequest(id, firstname, lastname string, poids float64) (*UserRequest, error) {
	if firstname == "" || lastname == "" || poids <= 50 {
		return &UserRequest{}, errors.New("error: one or more fields invalid")
	}
	return &UserRequest{
		Firstname: firstname,
		Lastame:   lastname,
		Poids:     poids,
	}, nil
}
