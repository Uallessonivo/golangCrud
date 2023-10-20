package models

import (
	"github.com/Uallessonivo/golangCrud/src/configuration/logger"
	"github.com/Uallessonivo/golangCrud/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *UserDomain) FindUser(string) (*UserDomain, *rest_errors.RestErr) {
	logger.Info("Init FindUser model", zap.String("journey", "find_user"))
	return nil, nil
}
