package util

import (
	"net/mail"
	"regexp"
)

// Done
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

func PasswordStrong(password string) bool {
	switch {
	case regexp.MustCompile(`^.{8,}$`).MatchString(password):
		return true
	case regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`).MatchString(password):
		return true
	case regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[^\da-zA-Z]).{8,}$`).MatchString(password):
		return true
	default:
		return false
	}
}
