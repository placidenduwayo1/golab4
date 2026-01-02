package config

import (
	"sync"

	"github.com/trng-tr/golab4/internal1/domain"
)

const ConferenceName string = "Go conference"

var ConferenceTickets uint = 50
var Bookings []domain.User = make([]domain.User, 0)
var Wg = sync.WaitGroup{}
