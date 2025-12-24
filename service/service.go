package service

import "github.com/trng-tr/golab4/impl"

type BookingService interface {
	BookTicket(u impl.User)
	SendTicket(u impl.User)
	PrintUser(u impl.User)
}
