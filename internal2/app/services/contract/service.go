package contract

import (
	"github.com/trng-tr/golab4/internal2/app/dtos"
)

// UserCreateService conttract to create a user
type UserCreateService interface {
	Create(u dtos.UserRequest) (*dtos.UserResponse, error)
}

// UserSendSerice contract to send User via a channel
type UserSendSerice interface {
	Send(chUserResp chan dtos.UserResponse, userResponse *dtos.UserResponse) error
}
