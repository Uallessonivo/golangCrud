package controllers

import (
	"fmt"

	"github.com/Uallessonivo/golangCrud/src/configuration/rest_errors"
	"github.com/Uallessonivo/golangCrud/src/models/requests"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest requests.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_errors.NewBadRequestErr(fmt.Sprintf("There are some incorrect fields, error: %s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
}
