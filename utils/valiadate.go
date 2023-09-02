package utils

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("patternAlphaUnderscore", PatternAlphaUnderscore)
}

func PatternAlphaUnderscore(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	pattern := `^[a-zA-Z0-9_,: \[\]{}%\-"]*$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(value)
}

func ValidasiStruct(data interface{}) bool {
	if err := validate.Struct(data); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return false
		}

		var errorMessages []string
		for _, e := range validationErrors {
			fieldName := e.Field()
			errorMessage := fmt.Sprintf("Field '%s' validation failed", fieldName)
			errorMessages = append(errorMessages, errorMessage)
		}

		fmt.Println("errors :", errorMessages)
		return false
	} else {
		return true
	}

}

func ValidateBackDate(start, end string) bool {
	startDate, _ := time.Parse("2006-01-02", start)

	endDate, _ := time.Parse("2006-01-02", end)

	if startDate.After(endDate) {
		return false
	} else {
		return true
	}

}

func ValidateChar(input string) bool {
	regexPattern := "^[A-Za-z0-9 %,]+$"

	match, err := regexp.MatchString(regexPattern, input)
	if err != nil {
		return false
	}
	return match
}
