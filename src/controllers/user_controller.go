package controllers

import (
	"github.com/Uallessonivo/golangCrud/src/models/services"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface services.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
}

type userControllerInterface struct {
	service services.UserDomainService
}
