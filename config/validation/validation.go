package validation

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/ogabrielrodrigues/crud-huncoding/config/rest"
)

var (
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enl := en.New()
		unt := ut.New(enl, enl)
		transl, _ = unt.GetTranslator("en")

		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

// func ValidateErr(err error) *rest.RestErr {
// 	var json_err *json.UnmarshalTypeError
// 	var validation_err validator.ValidationErrors

// 	if errors.As(err, &json_err) {
// 		return rest.NewBadRequestErr("invalid field type", nil)
// 	} else if errors.As(err, &validation_err) {
// 		causes := []rest.Cause{}

// 		for _, e := range err.(validator.ValidationErrors) {
// 			cause := rest.Cause{
// 				Message: e.Translate(transl),
// 				Field:   e.Field(),
// 			}

// 			fmt.Println(cause)

// 			causes = append(causes, cause)
// 		}

// 		return rest.NewBadRequestErr("some field are invalid", causes)
// 	} else {
// 		return rest.NewBadRequestErr("error trying to convert fields", nil)
// 	}
// }

func ValidateUserError(
	validation_err error,
) *rest.RestErr {

	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest.NewBadRequestErr("Invalid field type", nil)
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest.Cause{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest.Cause{
				Message: strings.ToLower(e.Translate(transl)),
				Field:   strings.ToLower(e.Field()),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest.NewBadRequestErr("Some fields are invalid", errorsCauses)
	} else {
		return rest.NewBadRequestErr("Error trying to convert fields", nil)
	}
}
