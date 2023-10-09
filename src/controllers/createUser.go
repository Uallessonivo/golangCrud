package controllers

import (
	"net/http"

	"github.com/Uallessonivo/golangCrud/src/configuration/logger"
	"github.com/Uallessonivo/golangCrud/src/configuration/validations"
	"github.com/Uallessonivo/golangCrud/src/models/requests"
	"github.com/Uallessonivo/golangCrud/src/models/responses"
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

	response := responses.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully", zap.String("journey", "create_user"))

	c.JSON(http.StatusOK, response)
}
