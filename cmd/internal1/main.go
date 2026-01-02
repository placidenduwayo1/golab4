package main

import (
	"fmt"

	"github.com/trng-tr/golab4/internal1/app/services/config"
	"github.com/trng-tr/golab4/internal1/app/services/impl"
	"github.com/trng-tr/golab4/internal1/app/services/service"
)

func main() {
	user, _ := getUserInput()
	var svc service.BookingService = impl.NewBookingServiceImpl()
	svc.BookTicket(user)
	config.Wg.Add(2)
	go svc.SendTicket(user)   // start a new goroutine (thread)
	go svc.DoOtherWorks(user) // start a new goroutine (thread)
	fmt.Println("##############################################")
	fmt.Println("Conference: ", config.ConferenceName)
	fmt.Println("other instructions")

	config.Wg.Wait()
}
