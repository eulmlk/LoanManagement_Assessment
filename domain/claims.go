package domain

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims interface {
	Valid() error
	SetExpiry()
	GetSecretKey() []byte
}

type LoginClaims struct {
	UserID string `json:"id"`
	Type   string `json:"type"`
	jwt.StandardClaims
}

type RegisterClaims struct {
	User `json:"user"`
	jwt.StandardClaims
}

type ResetClaims struct {
	UserID      string `json:"id"`
	NewPassword string `json:"new_password"`
	jwt.StandardClaims
}

func (c *LoginClaims) Valid() error {
	return c.StandardClaims.Valid()
}

func (c *RegisterClaims) Valid() error {
	return c.StandardClaims.Valid()
}

func (c *ResetClaims) Valid() error {
	return c.StandardClaims.Valid()
}

func (c *LoginClaims) SetExpiry() {
	var expiry time.Duration
	if c.Type == "refresh" {
		expiry = time.Hour * 24 * 7
	} else {
		expiry = time.Minute * 5
	}

	c.ExpiresAt = time.Now().Add(expiry).Unix()
}

func (c *RegisterClaims) SetExpiry() {
	c.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
}

func (c *ResetClaims) SetExpiry() {
	c.ExpiresAt = time.Now().Add(time.Hour).Unix()
}

func (c *LoginClaims) GetSecretKey() []byte {
	if c.Type == "refresh" {
		return []byte("my-refresh-secret-key")
	} else {
		return []byte("my-access-secret-key")
	}
}

func (c *RegisterClaims) GetSecretKey() []byte {
	return []byte("my-register-secret-key")
}

func (c *ResetClaims) GetSecretKey() []byte {
	return []byte("my-reset-secret-key")
}
