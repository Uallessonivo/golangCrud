package validations

import (
	"encoding/json"
	"errors"

	"github.com/Uallessonivo/golangCrud/src/configuration/rest_errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserErr(validationErr error) *rest_errors.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_errors.NewBadRequestErr("invalid json body")
	} else if errors.As(validationErr, &jsonValidationErr) {
		errorCauses := []rest_errors.Causes{}

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_errors.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return rest_errors.NewBadRequestValidationErr("Some fields are invalid", errorCauses)
	} else {
		return rest_errors.NewBadRequestErr("Error trying to convert fields")
	}
}
