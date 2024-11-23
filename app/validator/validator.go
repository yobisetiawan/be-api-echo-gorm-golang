package validator

import (
	"be_api/app/utils"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

type ValidationErrorResponse struct {
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors"`
}

func NewValidationErrorResponse(err error, s any) *ValidationErrorResponse {
	var ve validator.ValidationErrors

	typ := reflect.TypeOf(s)

	fmt.Println(typ)

	if errors.As(err, &ve) {
		errorsMap := make(map[string][]string)

		for _, fe := range ve {

			var sTag = fe.Tag()

			if sTag == "required" {
				sTag = "This field is required"
			}
			if sTag == "min" {
				sTag = "This field is too short"
			}
			if sTag == "max" {
				sTag = "This field is too long"
			}
			if sTag == "url" {
				sTag = "Invalid URL format"
			}
			if sTag == "email" {
				sTag = "Invalid email address"
			}
			if strings.HasPrefix(sTag, "unique") {
				sTag = "Data already exists in our database!"
			}

			field, ok := typ.FieldByName(fe.Field())
			if !ok {
				utils.CustomLog("Field not found")
			}

			errorsMap[field.Tag.Get("json")] = []string{sTag}
		}

		return &ValidationErrorResponse{
			Message: "Validation failed",
			Errors:  errorsMap,
		}
	}

	// Handle other types of errors if necessary
	return &ValidationErrorResponse{
		Message: "Invalid request",
		Errors: map[string][]string{
			"err": {err.Error()},
		},
	}
}

func RegisterValidators(echoValidator echo.Validator) {
	cv, ok := echoValidator.(*CustomValidator)
	if !ok {
		utils.CustomLog("echoValidator not ok!")
		return // Handle this error or log it if needed
	}

	// Register custom validation rules here, for example:
	cv.validator.RegisterValidation("uniqueProductCategoryTitle", UniqueProductCategoryTitle)
}
