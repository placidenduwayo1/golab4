package impl

import "fmt"

type User struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

func NewUserData(fn string, ln string, em string, nbT uint) User {
	return User{
		FirstName:       fn,
		LastName:        ln,
		Email:           em,
		NumberOfTickets: nbT,
	}
}
func (u User) PrintUser() {
	fmt.Println("{")
	fmt.Println(" FirstName:", u.FirstName)
	fmt.Println(" LastName:", u.LastName)
	fmt.Println(" Email:", u.Email)
	fmt.Println(" NumberOfTickets:", u.NumberOfTickets)
	fmt.Println("}")
}
