package view

import (
	"github.com/ogabrielrodrigues/crud-huncoding/internal/handler/model/response"
	"github.com/ogabrielrodrigues/crud-huncoding/internal/model"
)

func ConvertDomainToResponse(user_domain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    user_domain.GetID(),
		Name:  user_domain.GetName(),
		Email: user_domain.GetEmail(),
		Age:   user_domain.GetAge(),
	}
}
