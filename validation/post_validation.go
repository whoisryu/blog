package validation

import (
	"blog/exception"
	"blog/model"
	"encoding/json"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidatePost(post model.CreatePostRequest) {
	err := validation.ValidateStruct(&post,
		validation.Field(&post.Content,
			validation.Required.Error("NOT_BLANK"),
			validation.Match(regexp.MustCompile(`^[0-9]{20,}$`)).Error("MIN_20")),
		validation.Field(&post.Title,
			validation.Required.Error("NOT_BLANK"),
			is.Alphanumeric.Error("NOT_VALID")),
		validation.Field(&post.IsPublished,
			validation.Required.Error("NOT_BLANK")),
		validation.Field(&post.Categories,
			validation.Required.Error("NOT_BLANK")),
	)

	if err != nil {
		b, _ := json.Marshal(err)
		panic(exception.ValidationError{
			Message: string(b),
		})

	}
}

func ValidateUpdatePost(post model.UpdatePostRequest) {
	err := validation.ValidateStruct(&post,
		validation.Field(&post.Content,
			validation.Required.Error("NOT_BLANK"),
			validation.Match(regexp.MustCompile(`^[0-9]{20,}$`)).Error("MIN_20")),
		validation.Field(&post.Title,
			validation.Required.Error("NOT_BLANK"),
			is.Alphanumeric.Error("NOT_VALID")),
		validation.Field(&post.IsPublished,
			validation.Required.Error("NOT_BLANK")),
		validation.Field(&post.Categories,
			validation.Required.Error("NOT_BLANK")),
	)

	if err != nil {
		b, _ := json.Marshal(err)
		panic(exception.ValidationError{
			Message: string(b),
		})

	}
}
