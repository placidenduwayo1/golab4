package impl

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
