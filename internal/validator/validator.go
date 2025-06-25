package validator

import (
	"hotel/internal/logger"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/thoas/go-funk"
)

var Validate *validator.Validate


func isTypeOfRoom(fl validator.FieldLevel)  bool {
	single := fl.Field().String()
	single = strings.ToLower(single)
	types := [3]string{"suite", "double", "single"}
	return funk.Contains(types, single)
}



func StartValidate() {
	Validate = validator.New()
	Validate.RegisterValidation("typeofroom", isTypeOfRoom)
	logger.ZapLogger.Info("validator is ready!")
}