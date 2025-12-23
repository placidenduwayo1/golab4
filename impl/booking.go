package impl

import (
	"fmt"
	"time"
)

// BookTicket interface implementation,book ticket to user
func (User) BookTicket(u User) {
	remainingTickets = remainingTickets - u.NumberOfTickets
	bookings = append(bookings, u)
	fmt.Println("Thank you for your impl!")
	u.PrintUser()
	fmt.Println("list of bookings in the following list:")
	for _, booking := range bookings {
		booking.PrintUser()
	}

	fmt.Println("remaining tickets:", remainingTickets)
}

// SendTicket interface implementation, sending ticket to User
func (User) SendTicket(u User) {
	/* sending ticket take 50 seconds, to book another ticket for another
	User with 'BookTicket', program waits until sending finishes.
	to simulate this, we put instruction of sleeping:
	this means that delay instruction is part of actions of the method */
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v",
		u.NumberOfTickets, u.FirstName, u.LastName)
	fmt.Printf("Sending ticket: %v\n to email address %v\n",
		ticket, u.Email)
}
