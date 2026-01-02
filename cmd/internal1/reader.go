package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/mail"
	"os"
	"strconv"
	"strings"

	"github.com/trng-tr/golab4/internal1/app/services/config"
	"github.com/trng-tr/golab4/internal1/domain"
)

var streamReader = bufio.NewReader(os.Stdin)

func getUserInput() (domain.User, error) {
	fmt.Print("input user firstname: ")
	firstname, err := streamReader.ReadString('\n')
	if err != nil {
		return domain.User{}, errors.New("error reading user firstname")
	}
	firstname = strings.TrimSpace(firstname)
	if firstname == "" {
		return domain.User{}, errors.New("error empty user firstname")
	}
	fmt.Print("input user lastname: ")
	lastname, err := streamReader.ReadString('\n')
	if err != nil {
		return domain.User{}, errors.New("error reading user lastname")
	}
	lastname = strings.TrimSpace(lastname)
	if lastname == "" {
		return domain.User{}, errors.New("error empty user lastname")
	}
	fmt.Print("input user email address: ")
	email, err := streamReader.ReadString('\n')
	if err != nil {
		return domain.User{}, errors.New("error reading user email")
	}
	email = strings.TrimSpace(email)
	if email == "" {
		return domain.User{}, errors.New("error empty user email")
	}
	if _, emailParseErr := mail.ParseAddress(email); emailParseErr != nil {
		return domain.User{}, errors.New("error invalid email")
	}
	fmt.Print("input user number of tickets: ")
	tickets, err := streamReader.ReadString('\n')
	if err != nil {
		return domain.User{}, errors.New("error reading user number of tickets")
	}
	tickets = strings.TrimSpace(tickets)
	if tickets == "" {
		return domain.User{}, errors.New("error empty user number of tickets")
	}
	nbTickets, err := strconv.Atoi(tickets)
	if err != nil {
		return domain.User{}, errors.New("error parsing user number of tickets")
	}
	if nbTickets >= int(config.ConferenceTickets) {
		return domain.User{}, errors.New("error: Invalid number of tickets")
	}
	return domain.NewUser(firstname, lastname, email, uint(nbTickets)), nil
}
