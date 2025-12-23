package main

import (
	"github.com/trng-tr/golab3/impl"
	"github.com/trng-tr/golab3/reader"
	"github.com/trng-tr/golab3/service"
)

func main() {
	for {
		user, _ := reader.GetUserInput()
		var svc service.BookingService = impl.User{}
		svc.BookTicket(user)
		svc.SendTicket(user)
	}
}
