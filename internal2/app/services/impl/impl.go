package impl

import (
	"errors"

	"github.com/google/uuid"
	"github.com/trng-tr/golab4/internal2/app/dtos"
	"github.com/trng-tr/golab4/internal2/app/repositories/repo"
)

// UserCreateServiceImpl structurecto implement UserCreateService
type UserCreateServiceImpl struct {
	repository repo.Repository
}

// NewUserCreateServiceImpl constructor function
func NewUserCreateServiceImpl(repo repo.Repository) *UserCreateServiceImpl {
	return &UserCreateServiceImpl{
		repository: repo,
	}
}

// UserSenderServiceImpl structurecto implement UserSenderService
type UserSenderServiceImpl struct{}

// NewUserSenderServiceImpl constructor function
func NewUserSenderServiceImpl() *UserSenderServiceImpl {
	return &UserSenderServiceImpl{}
}

// Create implement interface  UserCreateService
func (ucsi *UserCreateServiceImpl) Create(u dtos.UserRequest) (*dtos.UserResponse, error) {
	if ucsi == nil {
		return nil, errors.New("service is nil")
	}
	if ucsi.repository == nil {
		return nil, errors.New("repository is nil (service not initialized)")
	}
	var user = dtos.ToUser(u)
	user.Id = uuid.NewString()[:5]
	user.Email = u.Firstname + "." + u.Lastame + "@" + "domain.com"
	savedUser, err := ucsi.repository.Save(user)
	if err != nil {
		return nil, err
	}
	if savedUser == nil {
		return nil, errors.New("repository.Save returned nil user without error")
	}

	var response = dtos.ToUserResponse(*savedUser)
	return &response, nil
}

// Send implement interface that send user response via a channel
func (*UserSenderServiceImpl) Send(chUserResp chan dtos.UserResponse, userResponse *dtos.UserResponse) error {
	if userResponse == nil {
		return errors.New("error: user response to send is nil")
	}
	chUserResp <- *userResponse // send userResponse via a chanel chUserResp
	return nil
}
