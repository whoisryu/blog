package validation

import (
	"blog/exception"
	"blog/model"
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateComment(comment model.CommentRequest) {
	err := validation.ValidateStruct(&comment,
		validation.Field(&comment.Content,
			validation.Required.Error("NOT_BLANK")),
		validation.Field(&comment.ParentID,
			validation.Required.Error("NOT_BLANK"),
			is.Alphanumeric.Error("NOT_VALID")),
		validation.Field(&comment.CommentBy,
			validation.Required.Error("NOT_BLANK")),
		validation.Field(&comment.PostID,
			validation.Required.Error("NOT_BLANK")),
	)

	if err != nil {
		b, _ := json.Marshal(err)
		panic(exception.ValidationError{
			Message: string(b),
		})

	}
}
