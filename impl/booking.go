package impl

/* implementation of interface methods*/
import (
	"fmt"
	"time"
)

func (User) BookTicket(u User) {
	conferenceTickets = conferenceTickets - u.NumberOfTickets
	bookings = append(bookings, u)
	fmt.Println("Thank you for your impl!")
	// PrintUser implements interface BookingService 👇
	var svc = User{}
	svc.PrintUser(u)
	fmt.Println("list of bookings in the following list:")
	for _, user := range bookings {
		svc.PrintUser(user)
	}
	fmt.Println("remaining tickets:", conferenceTickets)
}

func (User) SendTicket(u User) {
	/* sending ticket take 50 seconds, to book another ticket for another User with 'BookTicket',
	program waits until sending finishes.to simulate this, we put instruction of sleeping:
	this means that delay instruction is part of actions of the method */
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v",
		u.NumberOfTickets, u.FirstName, u.LastName)
	fmt.Printf("Sending ticket: %v\n to email address %v\n",
		ticket, u.Email)
	Wg.Done()
}

func (User) PrintUser(u User) {
	fmt.Println("{")
	fmt.Println(" FirstName:", u.FirstName)
	fmt.Println(" LastName:", u.LastName)
	fmt.Println(" Email:", u.Email)
	fmt.Println(" NumberOfTickets:", u.NumberOfTickets)
	fmt.Println("}")
}
