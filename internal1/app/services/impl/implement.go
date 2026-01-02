package impl

import (
	"fmt"
	"time"

	"github.com/trng-tr/golab4/internal1/app/services/config"
	"github.com/trng-tr/golab4/internal1/domain"
)

type BookingServiceImpl struct {
}

// NewBookingServiceImpl : function contructor
func NewBookingServiceImpl() *BookingServiceImpl {
	return &BookingServiceImpl{}
}

// BookTicket implementation of interface service.BookingService
func (bsi *BookingServiceImpl) BookTicket(u domain.User) {
	config.ConferenceTickets = config.ConferenceTickets - u.NumberOfTickets
	config.Bookings = append(config.Bookings, u)
	fmt.Println("Thank you for your impl!")
	fmt.Println("list of bookings in the following list:")
	for _, user := range config.Bookings {
		fmt.Println(user.String())
	}
	fmt.Println("remaining tickets:", config.ConferenceTickets)
}

// SendTicket implementation of interface service.BookingService
func (bsi *BookingServiceImpl) SendTicket(u domain.User) {
	/* sending ticket take 50 seconds, to book another ticket for another User with 'BookTicket',
	program waits until sending finishes. To simulate this situation, we put sleep instruction:
	this means that delay instruction is part of actions of the method */
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v",
		u.NumberOfTickets, u.FirstName, u.LastName)
	fmt.Printf("Sending ticket: %v\n to email address %v\n",
		ticket, u.Email)
	config.Wg.Done()
}

// DoOtherWorks implementation of interface service.BookingService
func (bsi *BookingServiceImpl) DoOtherWorks(u domain.User) {
	time.Sleep(30 * time.Second)
	fmt.Printf("%v %v do other actions\n", u.FirstName, u.LastName)
	config.Wg.Done()
}
