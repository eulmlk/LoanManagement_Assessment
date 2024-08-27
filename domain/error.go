package domain

import (
	"errors"
	"net/http"
)

var (
	ErrUsernameAlreadyExists        = errors.New("username already exists")
	ErrEmailAlreadyExists           = errors.New("email already exists")
	ErrInvalidUsernameLength        = errors.New("username must be between 3 and 30 characters")
	ErrInvalidUsernameChars         = errors.New("username can only contain alphanumeric characters, underscores, or hyphens")
	ErrInvalidEmailLength           = errors.New("email must be between 3 and 320 characters")
	ErrInvalidEmailFormat           = errors.New("invalid email address format")
	ErrWeakPasswordLength           = errors.New("password must be at least 8 characters long")
	ErrWeakPasswordUpper            = errors.New("password must contain at least one uppercase letter")
	ErrWeakPasswordLower            = errors.New("password must contain at least one lowercase letter")
	ErrWeakPasswordNumber           = errors.New("password must contain at least one number")
	ErrWeakPasswordSpecial          = errors.New("password must contain at least one special character")
	ErrInvalidToken                 = errors.New("the token is invalid")
	ErrInvalidUsernameEmailPassword = errors.New("incorrect username, email or password")
	ErrUserNotFoundByID             = errors.New("user with the given ID not found")
	ErrUserNotFoundByEmail          = errors.New("user with the given email not found")
	ErrInvalidID                    = errors.New("invalid ID")
)

func GetStatus(Err error) int {
	switch Err {
	case ErrUsernameAlreadyExists, ErrEmailAlreadyExists:
		return http.StatusConflict
	case ErrInvalidUsernameLength, ErrInvalidUsernameChars, ErrInvalidEmailLength, ErrInvalidEmailFormat, ErrWeakPasswordLength, ErrWeakPasswordUpper, ErrWeakPasswordLower, ErrWeakPasswordNumber, ErrWeakPasswordSpecial, ErrInvalidUsernameEmailPassword, ErrInvalidID:
		return http.StatusBadRequest
	case ErrInvalidToken:
		return http.StatusUnauthorized
	case ErrUserNotFoundByID, ErrUserNotFoundByEmail:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
