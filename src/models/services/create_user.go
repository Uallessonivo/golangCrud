package services

import (
	"fmt"

	"github.com/Uallessonivo/golangCrud/src/configuration/logger"
	"github.com/Uallessonivo/golangCrud/src/configuration/rest_errors"
	"github.com/Uallessonivo/golangCrud/src/models"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain models.UserDomainInterface) *rest_errors.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "create_user"))

	userDomain.EncryptPassword()
	fmt.Println(ud)

	return nil
}
