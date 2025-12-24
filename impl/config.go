package impl

import "sync"

const ConferenceName string = "Go conference"

var conferenceTickets uint = 50
var bookings []User = make([]User, 0)
var Wg = sync.WaitGroup{}
