package services

import (
	"github.com/Uallessonivo/golangCrud/src/configuration/rest_errors"
	"github.com/Uallessonivo/golangCrud/src/models"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{
		
	}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(models.UserDomainInterface) *rest_errors.RestErr
	UpdateUser(string, models.UserDomainInterface) *rest_errors.RestErr
	FindUser(string) (*models.UserDomainInterface, *rest_errors.RestErr)
	DeleteUser(string) *rest_errors.RestErr
}
