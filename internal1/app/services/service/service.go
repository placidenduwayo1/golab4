package service

import (
	"github.com/trng-tr/golab4/internal1/domain"
)

type BookingService interface {
	BookTicket(u domain.User)
	SendTicket(u domain.User)   // will take much time, so will be a goroutine
	DoOtherWorks(u domain.User) // will take much time,so will be a goroutine
}
