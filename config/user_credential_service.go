package config

import (
	"loans/domain"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func IsValidUsername(username string) error {
	// Check if the username is between 3 and 30 characters
	if len(username) < 3 || len(username) > 30 {
		return domain.ErrInvalidUsernameLength
	}

	// Check if the username contains only alphanumeric characters and underscores or hyphens
	for _, char := range username {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') || char == '_' || char == '-') {
			return domain.ErrInvalidUsernameChars
		}
	}

	return nil
}

func IsValidEmail(email string) error {
	// Check if the email is between 3 and 320 characters
	if len(email) < 3 || len(email) > 320 {
		return domain.ErrInvalidEmailLength
	}

	// Regular expression for validating an email address
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	result := re.MatchString(email)

	if !result {
		return domain.ErrInvalidEmailFormat
	}

	return nil
}

func IsStrongPassword(password string) error {
	// Check if the password is at least 8 characters long
	if len(password) < 8 {
		return domain.ErrWeakPasswordLength
	}

	// The password must contain at least one uppercase letter, one lowercase letter,
	// one digit, and one special character to be considered strong
	hasUppercase, hasLowercase, hasDigit, hasSpecial := false, false, false, false

	for _, char := range password {
		// Check if the password contains at least one uppercase letter
		if char >= 'A' && char <= 'Z' {
			hasUppercase = true
		}

		// Check if the password contains at least one lowercase letter
		if char >= 'a' && char <= 'z' {
			hasLowercase = true
		}

		// Check if the password contains at least one digit
		if char >= '0' && char <= '9' {
			hasDigit = true
		}

		// Check if the password contains at least one special character
		if (char >= 33 && char <= 47) || (char >= 58 && char <= 64) ||
			(char >= 91 && char <= 96) || (char >= 123 && char <= 126) {
			hasSpecial = true
		}
	}

	if !hasUppercase {
		return domain.ErrWeakPasswordUpper
	}

	if !hasLowercase {
		return domain.ErrWeakPasswordLower
	}

	if !hasDigit {
		return domain.ErrWeakPasswordNumber
	}

	if !hasSpecial {
		return domain.ErrWeakPasswordSpecial
	}

	return nil
}
