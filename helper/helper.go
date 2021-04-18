package helper

import (
	"blog/exception"

	"golang.org/x/crypto/bcrypt"
)

func HashPwd(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	exception.PanicIfNeeded(err)

	return string(hash)
}
