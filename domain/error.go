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
	ErrAlreadyHasLoan               = errors.New("user already has a loan")
	ErrInvalidUserID                = errors.New("invalid user ID")
	ErrLoanNotFoundByUserID         = errors.New("this user does not have a loan")
	ErrOnlyRootCanPromote           = errors.New("only root can promote")
	ErrAlreadyPromoted              = errors.New("this user is already an admin")
	ErrAlreadyDemoted               = errors.New("this user is already a regular user")
	ErrInvalidIDPromote             = errors.New("invalid ID to promote")
	ErrInvalidUserIDDelete          = errors.New("invalid user ID to delete")
	ErrPageNotFound                 = errors.New("page not found")
	ErrOnlyAdminCanViewAllUsers     = errors.New("only an admin can view all users")
	ErrOnlyAdminCanDelete           = errors.New("only an admin can delete other users")
	ErrOnlyRootCanDelete            = errors.New("only root can delete admin users")
	ErrCantDeleteRoot               = errors.New("cannot delete root user")
)

func GetStatus(Err error) int {
	switch Err {
	case ErrUsernameAlreadyExists, ErrEmailAlreadyExists, ErrAlreadyHasLoan, ErrAlreadyPromoted, ErrAlreadyDemoted:
		return http.StatusConflict
	case ErrInvalidUsernameLength, ErrInvalidUsernameChars, ErrInvalidEmailLength, ErrInvalidEmailFormat, ErrWeakPasswordLength, ErrWeakPasswordUpper, ErrWeakPasswordLower, ErrWeakPasswordNumber, ErrWeakPasswordSpecial, ErrInvalidUsernameEmailPassword, ErrInvalidID, ErrInvalidUserID, ErrInvalidIDPromote, ErrInvalidUserIDDelete:
		return http.StatusBadRequest
	case ErrInvalidToken:
		return http.StatusUnauthorized
	case ErrUserNotFoundByID, ErrUserNotFoundByEmail, ErrLoanNotFoundByUserID, ErrPageNotFound:
		return http.StatusNotFound
	case ErrOnlyRootCanPromote, ErrOnlyAdminCanViewAllUsers, ErrOnlyAdminCanDelete, ErrOnlyRootCanDelete, ErrCantDeleteRoot:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
