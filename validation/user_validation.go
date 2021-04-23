package validation

import (
	"blog/exception"
	"blog/model"
	"encoding/json"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateRegisterUser(user model.RegisterUserRequest) {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.Email,
			validation.Required.Error("NOT_BLANK"),
			is.Email.Error("NOT_VALID")),
		validation.Field(&user.Password,
			validation.Required.Error("NOT_BLANK"),
			validation.Match(regexp.MustCompile(`^.{5,}$`)).Error("MIN_5"),
			is.Alphanumeric.Error("NOT_VALID")),
		validation.Field(&user.Phone,
			validation.Required.Error("NOT_BLANK")),
		validation.Field(&user.UserName,
			validation.Required.Error("NOT_BLANK")))

	if err != nil {
		b, _ := json.Marshal(err)
		panic(exception.ValidationError{
			Message: string(b),
		})

	}
}

func ValidateLoginUser(user model.LoginRequest) {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.Email,
			validation.Required.Error("NOT_BLANK"),
			is.Email.Error("NOT_VALID")),
		validation.Field(&user.Password,
			validation.Required.Error("NOT_BLANK"),
			validation.Match(regexp.MustCompile(`^.{5,}$`)).Error("MIN_5"),
			is.Alphanumeric.Error("NOT_VALID")),
	)

	if err != nil {
		b, _ := json.Marshal(err)
		panic(exception.ValidationError{
			Message: string(b),
		})

	}
}

func ValidateUpdate(user model.UpdateProfileRequest) {
	err := validation.ValidateStruct(&user,
		validation.Field(&user.Email,
			is.Email.Error("NOT_VALID")),
	)

	if err != nil {
		b, _ := json.Marshal(err)
		panic(exception.ValidationError{
			Message: string(b),
		})

	}
}
