package dtos

import "github.com/trng-tr/golab4/internal2/domain"

func ToUserResponse(model domain.User) UserResponse {
	return UserResponse{
		Id:        model.Id,
		Firstname: model.Firstname,
		Lastame:   model.Lastname,
		Email:     model.Email,
		Poids:     model.Poids,
	}
}

func ToUser(ur UserRequest) *domain.User {
	return &domain.User{
		Firstname: ur.Firstname,
		Lastname:  ur.Lastame,
		Poids:     ur.Poids,
	}
}
