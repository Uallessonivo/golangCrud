package models

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Uallessonivo/golangCrud/src/configuration/rest_errors"
)

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type UserDomainInterface interface {
	CreateUser() *rest_errors.RestErr
	UpdateUser(string) *rest_errors.RestErr
	FindUser(string) (*UserDomain, *rest_errors.RestErr)
	DeleteUser(string) *rest_errors.RestErr
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
