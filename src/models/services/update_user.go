package services

import (
	"github.com/Uallessonivo/golangCrud/src/configuration/logger"
	"github.com/Uallessonivo/golangCrud/src/configuration/rest_errors"
	"github.com/Uallessonivo/golangCrud/src/models"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(userId string, userDomain models.UserDomainInterface) *rest_errors.RestErr {
	logger.Info("Init UpdateUser model", zap.String("journey", "update_user"))
	return nil
}
