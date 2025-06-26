package validator

import (
	"hotel/internal/logger"
	"strings"

	"regexp"
	"github.com/go-playground/validator/v10"
	"github.com/thoas/go-funk"
)

var Validate *validator.Validate

var phoneRegex = regexp.MustCompile(`^\d{3}-\d{3}-\d{4}$`)

func isTypeOfRoom(fl validator.FieldLevel)  bool {
	single := fl.Field().String()
	single = strings.ToLower(single)
	types := [3]string{"suite", "double", "single"}
	return funk.Contains(types, single)
}

func isAmericanPhoneNumber(fl validator.FieldLevel)  bool {
	phoneNumber := fl.Field().String()
	return phoneRegex.MatchString(phoneNumber)
}


func StartValidate() {
	Validate = validator.New()
	Validate.RegisterValidation("typeofroom", isTypeOfRoom)
	Validate.RegisterValidation("phone_number", isAmericanPhoneNumber)
	logger.ZapLogger.Info("validator is ready!")
}