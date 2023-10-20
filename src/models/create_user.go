package models

import (
	"fmt"

	"github.com/Uallessonivo/golangCrud/src/configuration/logger"
	"github.com/Uallessonivo/golangCrud/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_errors.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "create_user"))

	ud.EncryptPassword()
	fmt.Println(ud)

	return nil
}
