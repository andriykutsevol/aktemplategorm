package auth

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type (
	UserIDCtx struct{}
)

type (
	CIDCtx struct{}
)

type User struct {
	ID       *string
	UserName string
	Password string
	UUID     *uuid.UUID
}

type Auth struct {
	AccessToken string
	TokenType   string
	ExpiresAt   int64
}