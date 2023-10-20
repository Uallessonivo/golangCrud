package models

import (
	"github.com/Uallessonivo/golangCrud/src/configuration/logger"
	"github.com/Uallessonivo/golangCrud/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *UserDomain) UpdateUser(string) *rest_errors.RestErr {
	logger.Info("Init UpdateUser model", zap.String("journey", "update_user"))
	return nil
}
