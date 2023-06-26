package internal

import "fmt"

var (
	ErrEmailExist = fmt.Errorf("EmailExist")

	ErrWrongCaptcha = fmt.Errorf("WrongCaptcha")

	ErrEmailNotFound = fmt.Errorf("EmailNotFound")

	ErrPasswordIncorrect = fmt.Errorf("PasswordIncorrect")
)
