package services

import (
	"github.com/Uallessonivo/golangCrud/src/configuration/logger"
	"github.com/Uallessonivo/golangCrud/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(string) *rest_errors.RestErr {
	logger.Info("Init DeleteUser model", zap.String("journey", "delete_user"))
	return nil
}
