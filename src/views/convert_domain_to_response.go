package views

import (
	"github.com/Uallessonivo/golangCrud/src/controllers/models/responses"
	"github.com/Uallessonivo/golangCrud/src/models"
)

func ConvertDomainToResponse(userDomain models.UserDomainInterface) responses.UserResponse {
	return responses.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
