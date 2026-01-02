package domain

import "fmt"

type User struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

func NewUser(fn string, ln string, em string, nbT uint) User {
	return User{
		FirstName:       fn,
		LastName:        ln,
		Email:           em,
		NumberOfTickets: nbT,
	}
}

func (u *User) String() string {
	return fmt.Sprintf(`User {
 Firstname: %s
 Lastname: %s
 Email: %s
 Number of tickets: %d
}
	`, u.FirstName, u.FirstName, u.Email, u.NumberOfTickets)
}
