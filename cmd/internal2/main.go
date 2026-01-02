package main

import (
	"fmt"

	"github.com/trng-tr/golab4/internal2/app/dtos"
	imp1 "github.com/trng-tr/golab4/internal2/app/repositories/impl"
	"github.com/trng-tr/golab4/internal2/app/repositories/repo"
	"github.com/trng-tr/golab4/internal2/app/services/contract"
	imp2 "github.com/trng-tr/golab4/internal2/app/services/impl"
	"github.com/trng-tr/golab4/internal2/domain"
)

func main() {
	var userRequest = dtos.UserRequest{
		Firstname: "Placide",
		Lastame:   "Nduwayo",
		Poids:     65,
	}

	var re repo.Repository = imp1.NewRepositoryImpl(func(user *domain.User) string {
		return user.Id
	})
	var createUserService contract.UserCreateService = imp2.NewUserCreateServiceImpl(re)
	useResponse, err := createUserService.Create(userRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	printed := useResponse.String()
	fmt.Println(printed)

	var sendService contract.UserSendSerice = imp2.NewUserSenderServiceImpl()
	var ch chan dtos.UserResponse = make(chan dtos.UserResponse)
	go sendService.Send(ch, useResponse)
	fmt.Println("send user response ", <-ch)

}
