package dtos

import (
	"fmt"
)

// UserResponse structure
type UserResponse struct {
	Id        string
	Firstname string
	Lastame   string
	Email     string
	Poids     float64
}

// NewUserResponse : function construtor or structure
func NewUserResponse(id, firstname, lastname, email string, poids float64) (*UserResponse, error) {
	return &UserResponse{
		Id:        id,
		Firstname: firstname,
		Lastame:   lastname,
		Email:     email,
		Poids:     poids,
	}, nil
}

// String method of structure
func (u *UserResponse) String() string {
	return fmt.Sprintf(`User {
 Id: %s
 Firstname: %s
 Lastname: %s
 Email: %s
 Poids: %.2f
}`, u.Id, u.Firstname, u.Lastame, u.Email, u.Poids)
}
