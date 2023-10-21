package controllers

import (
	"net/http"

	"github.com/Uallessonivo/golangCrud/src/configuration/logger"
	"github.com/Uallessonivo/golangCrud/src/configuration/validations"
	"github.com/Uallessonivo/golangCrud/src/controllers/models/requests"
	"github.com/Uallessonivo/golangCrud/src/models"
	"github.com/Uallessonivo/golangCrud/src/models/services"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "create_user"))

	var userRequest requests.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error while binding JSON", err, zap.String("journey", "create_user"))

		restErr := validations.ValidateUserErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := models.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	service := services.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "create_user"))

	c.String(http.StatusOK, "")
}
