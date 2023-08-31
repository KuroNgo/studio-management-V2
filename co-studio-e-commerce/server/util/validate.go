package util

import (
	"net/mail"
	"regexp"
)

type IValidate interface {
	EmailValid(email string) bool
	PhoneValid(phone string) bool
}

func EmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func PhoneValid(phone string) bool {
	// Biểu thức chính quy để kiểm tra số điện thoại hợp lệ với đúng 10 chữ số
	pattern := `^(84|0[35789])\d{8}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(phone)
}
