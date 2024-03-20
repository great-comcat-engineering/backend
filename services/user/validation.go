package user

import (
	"greatcomcatengineering.com/backend/models"
	"regexp"
)

func ValidateUserRegister(user models.RegisterRequest) (string, bool) {
	// Email validation
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(user.Email) {
		return "Invalid email format", false
	}
	// TODO: Password validation
	return "", true
}
