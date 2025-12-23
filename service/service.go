package service

import "github.com/trng-tr/golab3/impl"

type BookingService interface {
	BookTicket(u impl.User)
	SendTicket(u impl.User)
}
