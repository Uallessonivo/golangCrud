package services

import (
	"github.com/Uallessonivo/golangCrud/src/configuration/logger"
	"github.com/Uallessonivo/golangCrud/src/configuration/rest_errors"
	"github.com/Uallessonivo/golangCrud/src/models"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUser(string) (*models.UserDomainInterface, *rest_errors.RestErr) {
	logger.Info("Init FindUser model", zap.String("journey", "find_user"))
	return nil, nil
}
