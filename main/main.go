package main

import (
	"fmt"

	"github.com/trng-tr/golab4/impl"
	"github.com/trng-tr/golab4/reader"
	"github.com/trng-tr/golab4/service"
)

func main() {
	user, _ := reader.GetUserInput()
	var svc service.BookingService = impl.User{}
	svc.BookTicket(user)
	impl.Wg.Add(2)
	go svc.SendTicket(user)    // start a new goroutine (thread)
	go impl.DoOtherWorks(user) // start a new goroutine (thread)
	fmt.Println("##############################################")
	fmt.Println("Conference: ", impl.ConferenceName)
	fmt.Println("other instructions")
	impl.Wg.Wait()
}
